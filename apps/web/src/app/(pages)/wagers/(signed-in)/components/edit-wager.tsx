'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { Vote, VoteIntent, Wager, WagerStatus } from '@/types/wager'
import { cn } from '@/lib/utils'
import { useAccount, useConfig } from 'wagmi'
import { writeContract } from '@wagmi/core'
import contractJson from '@/contracts/WagerEscrow.json' assert { type: 'json' }
import { getChainId, getPYMWYMIContractAddress, getTokenAddress } from '@/lib/blockchain'
import { Abi } from 'viem'

enum EditOptions {
    ACCEPT = 'Accept',
    DECLINE = 'Decline',
    CANCEL = 'Cancel',
    VOTE = 'Vote',
    CLAIM = 'Claim',
}


export default function EditWager({
    wager,
    walletAddress,
}: {
    wager: Wager | null
    walletAddress: string
}) {
    if (!wager) {
        return
    }

    const router = useRouter()
    const { address, isConnected } = useAccount()
    const config = useConfig()

    const [isSubmitting, setIsSubmitting] = useState(false)
    const [error, setError] = useState<string | null>(null)
    const [selectedAction, setSelectedAction] = useState<EditOptions | null>(null)
    const [voteIntent, setVoteIntent] = useState<VoteIntent | null>(null)

    const participant = wager?.participants.find((p) => p.walletAddress !== walletAddress)
    const user = wager?.participants.find((p) => p.walletAddress === walletAddress)

    const handleAcceptWager = async () => {
        // first we need user to sign transaction
        // then once we recievieve the transaction hash we must post to backend
        // we want to invalidate the received wagers
        if (!wager || !participant) return
        try {
            if (!isConnected || address !== walletAddress) {
                const dialog = document.getElementById('sign_in_modal') as HTMLDialogElement
                dialog.showModal()
                return
            }
            setIsSubmitting(true)

            const selectedChainId = getChainId(wager.chain)
            const contractAddress = getPYMWYMIContractAddress()
            const tokenAddress = getTokenAddress(wager.currency, selectedChainId)

            // first we must await the user sign the transaction...
            const transactionHash = await writeContract(config, {
                address: contractAddress,
                abi: contractJson.abi as Abi,
                functionName: 'acceptWager',
                args: [
                    wager.smartContractId,
                    BigInt(wager.stake),
                    tokenAddress,
                ],
            })
            console.log(`transaction hash: `, transactionHash)

            const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/wager/${wager.id}/accept`, {
                method: 'PATCH',
                body: JSON.stringify({
                    wagerId: wager.id,
                    stakeSignature: transactionHash
                }),
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                },
            })
            const data = await res.text()
            console.log(`api response: `, data)
            if (!res.ok) throw new Error(`Failed to submit vote: ${data}`)
        } catch (error) {
            console.error(error)
        }

    }

    const handleCancelWager = async () => {
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

    const handleClaimWinnings = async () => {
        console.log(wager)
        //
    }

    const editActions: Record<EditOptions, { label: string, handler: () => Promise<void>, confirmation: string, affectedQueryKeys: string[] }> = {
        [EditOptions.ACCEPT]: {
            label: "Accept",
            handler: handleAcceptWager,
            confirmation: "Are you sure you want to accept this wager?",
            affectedQueryKeys: ["receivedWagers", "pendingWagers"]
        },
        [EditOptions.DECLINE]: {
            label: "Decline",
            handler: handleCancelWager,
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
            label: "Vote",
            handler: handleVote,
            confirmation: "Are you sure you want to vote for this wager?",
            affectedQueryKeys: ["pendingWagers"]
        },
        [EditOptions.CLAIM]: {
            label: "Claim",
            handler: handleClaimWinnings,
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

    async function handleVote() {
        if (!wager || !participant || !voteIntent) return
        try {
            let vote: Vote
            if (voteIntent === VoteIntent.Winner) {
                vote = { hasVoted: true, intent: voteIntent, winner: walletAddress }
            } else {
                vote = { hasVoted: true, intent: voteIntent, winner: "" }
            }
            const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/wager/${wager.id}/vote`, {
                method: 'PATCH',
                body: JSON.stringify({
                    wagerId: wager.id,
                    vote
                }),
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                },
            })
            const data = await res.text()
            if (!res.ok) throw new Error(`Failed to submit vote: ${data}`)
        } catch (error) {
            console.error(error)
        }
    }

    async function submitAction(action: EditOptions) {
        try {
            setIsSubmitting(true)


            const editAction = editActions[action]
            if (!editAction) return

            console.log(`submitting action: ${editAction.label}`)

            await editAction.handler()

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
                                    editActions[action]
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
                    {
                        selectedAction === EditOptions.VOTE ? (
                            <>
                                <input
                                    type="radio"
                                    name="radio-12"
                                    defaultChecked
                                    onChange={() => setVoteIntent(VoteIntent.Cancel)}
                                    className="radio bg-red-100 border-red-300 checked:bg-red-200 checked:text-red-600 checked:border-red-600" />
                                <input
                                    type="radio"
                                    name="radio-12"
                                    defaultChecked
                                    onChange={() => setVoteIntent(VoteIntent.Winner)}
                                    className="radio bg-blue-100 border-blue-300 checked:bg-blue-200 checked:text-blue-600 checked:border-blue-600" />
                            </>
                        ) : (
                            <>

                            </>
                        )

                    }
                    <p className="py-4">
                        Are you sure you want to <strong>{selectedAction}</strong> this wager?
                    </p>
                    <div className="flex justify-end gap-3">
                        <button
                            className="btn btn-primary"
                            disabled={isSubmitting}
                            onClick={() => selectedAction && submitAction(selectedAction)}
                        >
                            {isSubmitting ? 'Processing...' : 'Confirm'}
                        </button>
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
                    </div>
                </div>
            </dialog>
        </>
    )
}
