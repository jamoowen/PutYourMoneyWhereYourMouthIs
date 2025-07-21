import Button from "@/components/common/button";
import { Player, Wager } from "@/types/wager";
import { useState } from "react";
import EditWager from "./edit-wager";

const getStatusLabel = (status: number) => {
    const labels = ["Created", "Pending", "Cancelled", "Completed", "Claimed"];
    return labels[status] ?? "Unknown";
};

const extractOtherParticipants = (participants: Player[], walletAddress: string): string[] => {
    return participants
        .filter((participant) => participant.walletAddress !== walletAddress)
        .map((participant) => participant.walletAddress)
}

const shortenParticipants = (participants: string[]): string => {
    return participants[0].substring(0, 10) + "..."
}

export default function WagersList({
    title,
    wagers,
    walletAddress,
}: {
    title: string;
    wagers: Wager[];
    walletAddress: string;
}) {
    const [selectedWager, setSelectedWager] = useState<Wager | null>(null);
    return (
        <>
            <div className="w-full ">
                <h2 className="text-lg font-semibold mb-2">{title}</h2>
                <ul className="rounded-box w-full shadow-md divide-y ">
                    {wagers.map((wager) => (
                        <li
                            key={wager.id}
                            className="flex items-center justify-between p-4 hover:bg-base-200 transition cursor-pointer"
                            onClick={() => {
                                setSelectedWager(wager)
                                const dialog = document.getElementById('edit_wager_modal') as HTMLDialogElement
                                dialog.showModal()
                            }}
                        >
                            <div className="flex flex-col space-y-1">
                                <span className="text-sm font-semibold">{wager.name}</span>
                                <span className="text-xs opacity-70">
                                    {/* {wager.participants.length} participant */}
                                    {/* {wager.participants.length !== 1 ? "s" : ""} */}

                                    Participant:
                                    {
                                        " " + shortenParticipants(extractOtherParticipants(wager.participants, walletAddress))
                                    }
                                </span>
                                <span className="text-xs opacity-70">
                                    Stake: {wager.stake} {wager.currency}
                                </span>
                                <span className="text-xs font-medium">
                                    Status: {getStatusLabel(wager.status)}
                                </span>
                            </div>
                            <div className="flex gap-2">
                            </div>
                        </li>
                    ))}
                </ul>
            </div>
            <EditWager wager={selectedWager} walletAddress={walletAddress} />
        </>
    );
}
