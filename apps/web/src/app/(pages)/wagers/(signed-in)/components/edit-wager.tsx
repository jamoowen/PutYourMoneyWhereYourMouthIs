'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { Wager, WagerStatus } from '@/types/wager'
import { cn } from '@/lib/utils'

enum EditOptions {
    ACCEPT = 'Accept',
    DECLINE = 'Decline',
    CANCEL = 'Cancel',
    VOTE = 'Vote',
    CLAIM = 'Claim',
}

const handleAcceptWager = async (wager: Wager) => {
    // first we need user to sign transaction

}

const handleVoteWager = async (wager: Wager, vote: { hasVoted: boolean, intent: string }) => {
    const payload = {
        wagerId: wager.id,
        vote
    }
    try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/wager/vote`, {
            method: 'PATCH',
            body: JSON.stringify(payload),
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
        })
        if (!res.ok) throw new Error('Request failed')
    } catch (err) {
        console.error(err)
    }
}

const handleCancelWager = async (wager: Wager) => {
    const payload = {
        wagerId: wager.id,
        vote: {
            hasVoted: true,
            intent: "cancel"
        }
    }
    try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/wager/vote`, {
            method: 'PATCH',
            body: JSON.stringify(payload),
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
        })
        if (!res.ok) throw new Error('Request failed')
    } catch (err) {
        console.error(err)
    }
}

export default function EditWager({
    wager,
    walletAddress,
}: {
    wager: Wager | null
    walletAddress: string
}) {
    const router = useRouter()
    const [isSubmitting, setIsSubmitting] = useState(false)
    const [error, setError] = useState<string | null>(null)
    const [selectedAction, setSelectedAction] = useState<EditOptions | null>(null)

    const participant = wager?.participants.find((p) => p.walletAddress !== walletAddress)
    const user = wager?.participants.find((p) => p.walletAddress === walletAddress)

    const EditActions = {
        [EditOptions.ACCEPT]: {
            label: "Accept",
            confirmation: "Are you sure you want to accept this wager?",
            affectedQueryKeys: ["receivedWagers", "pendingWagers"]
        },
        [EditOptions.DECLINE]: {
            label: "Decline",
            confirmation: "Are you sure you want to decline this wager?",
            affectedQueryKeys: ["receivedWagers"]
        },
        [EditOptions.CANCEL]: {
            label: "Cancel",
            handler: handleCancelWager,
            confirmation: "Are you sure you want to cancel this wager?",
            affectedQueryKeys: ["receivedWagers", "pendingWagers"]
        },
        [EditOptions.VOTE]: {
            route: "/vote",
            confirmation: "Are you sure you want to vote for this wager?",
            affectedQueryKeys: ["pendingWagers"]
        },
        [EditOptions.CLAIM]: {
            route: "/claim",
            confirmation: "Are you sure you want to claim this wager?",
            affectedQueryKeys: ["pendingWagers"]
        },
    }


    if (!wager) return null

    const getAvailableActions = () => {
        switch (wager.status) {
            case WagerStatus.Created:
                if (user?.hasStaked === false) {
                    return [EditOptions.ACCEPT, EditOptions.DECLINE]
                }
                return [EditOptions.CANCEL]
            case WagerStatus.Pending:
                return [EditOptions.VOTE, EditOptions.CANCEL]
            case WagerStatus.Completed:
                if (wager.winner === walletAddress) {
                    return [EditOptions.CLAIM]
                }
                return []
            default:
                return []
        }
    }

    const availableActions = getAvailableActions()

    async function submitAction(action: EditOptions) {
        try {
            const payload =
                setIsSubmitting(true)
            const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/wager/${wager?.id}/action`, {
                method: 'POST',
                body: JSON.stringify({ action }),
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                },
            })

            if (!res.ok) throw new Error('Request failed')

            setError(null)
            const editModal = (document.getElementById('edit_wager_modal') as HTMLDialogElement).close()
            const confirmModal = (document.getElementById('confirm_action_modal') as HTMLDialogElement).close()
            router.refresh()
        } catch (err) {
            console.error(err)
            setError(`Failed to perform action: ${action}`)
        } finally {
            setIsSubmitting(false)
            setSelectedAction(null)
        }
    }

    return (
        <>
            {/* Main Wager Modal */}
            <dialog id="edit_wager_modal" className="modal">
                <div className="modal-box">
                    <h3 className="text-lg font-bold mb-4">Wager Details</h3>

                    <div className="text-sm space-y-2 mb-4">
                        <p><strong>Name:</strong> {wager.name}</p>
                        <p><strong>Category:</strong> {wager.category}</p>
                        <p><strong>Description:</strong> {wager.description}</p>
                        <p><strong>Stake:</strong> {wager.stake} {wager.currency}</p>
                        <p><strong>Location:</strong> {wager.location}</p>
                        <p><strong>Status:</strong> {WagerStatus[wager.status]}</p>
                        <p><strong>Opponent:</strong> {participant?.walletAddress ?? "N/A"}</p>
                    </div>

                    {error && <p className="text-sm text-red-500 mb-4">{error}</p>}

                    <div className="flex flex-wrap gap-2">
                        {availableActions.map((action) => (
                            <button
                                key={action}
                                disabled={isSubmitting}
                                onClick={() => {
                                    setSelectedAction(action)
                                    const confirmModal = document.getElementById('confirm_action_modal') as HTMLDialogElement
                                    confirmModal?.showModal()
                                }}
                                className={cn(
                                    'btn',
                                    isSubmitting && 'loading',
                                    action === EditOptions.ACCEPT && 'btn-success',
                                    (action === EditOptions.DECLINE || action === EditOptions.CANCEL) && 'btn-error',
                                    action === EditOptions.VOTE && 'btn-primary',
                                    action === EditOptions.CLAIM && 'btn-success'
                                )}
                            >
                                {action}
                            </button>
                        ))}
                    </div>
                </div>

                <form method="dialog" className="modal-backdrop">
                    <button className="cursor-pointer">close</button>
                </form>
            </dialog>

            {/* Confirmation Modal */}
            <dialog id="confirm_action_modal" className="modal">
                <div className="modal-box">
                    <h3 className="font-bold text-lg">Confirm Action</h3>
                    <p className="py-4">
                        Are you sure you want to <strong>{selectedAction}</strong> this wager?
                    </p>
                    <div className="flex justify-end gap-3">
                        <button
                            className="btn btn-outline"
                            onClick={() => {
                                const dialog = document.getElementById('confirm_action_modal') as HTMLDialogElement
                                dialog?.close()
                                setSelectedAction(null)
                            }}
                        >
                            Cancel
                        </button>
                        <button
                            className="btn btn-primary"
                            disabled={isSubmitting}
                            onClick={() => selectedAction && submitAction(selectedAction)}
                        >
                            {isSubmitting ? 'Processing...' : 'Confirm'}
                        </button>
                    </div>
                </div>
            </dialog>
        </>
    )
}
