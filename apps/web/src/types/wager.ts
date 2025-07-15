// submitting a vote changes the opinion of the participant on the outcome of the wager
export enum VoteIntent {
  Unknown = "unknown",
  Cancel = "cancel",
  Winner = "winner"
}

// pending means we have not yet confirmed this on chain
export enum InteractionStatus {
  Dormant = "dormant",
  Pending = "pending",
  Confirmed = "confirmed"
}

export enum WagerStatus {
  Created = 0,
  Pending = 1,
  Cancelled = 2,
  Completed = 3,
  Claimed = 4
}

export type Vote = {
  hasVoted: boolean
  Intent: VoteIntent
  Winner: string
}

export type Player = {
  walletAddress: string
  vote: string
  hasStaked: boolean
  stakeStatus: string
  hasWithdrawn: boolean
  withdrawalStatus: string
}

export type Wager = {
  id: number
  transactionHash: string
  creationStatus: number
  name: string
  category: string
  description: string
  location: string
  stake: number
  currency: string
  chain: string
  participants: Player[]
  status: number
  winner: string
  createdAt: string
  updatedAt: string
}


