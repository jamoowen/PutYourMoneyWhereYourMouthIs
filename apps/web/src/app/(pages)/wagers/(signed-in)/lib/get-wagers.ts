import { PaginatedResponse } from '@/types/common'
import { Wager, WagerStatus } from '@/types/wager'
import { queryOptions, useQuery } from '@tanstack/react-query'

export type ExtraQueryOptions = {
    creator?: boolean
    winner?: boolean
}

const fetchWagers = async (status: number, page: number, limit: number, options?: ExtraQueryOptions): Promise<PaginatedResponse<Wager[]>> => {
    let url = process.env.NEXT_PUBLIC_API_URL + `/wager/list?status=${status}&page=${page}&limit=${limit}`
    if (options?.creator != null) {
        url += `&creator=${options.creator}`
    }

    if (options?.winner != null) {
        url += `&winner=${options.winner}`
    }

    const response = await fetch(url, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
    })
    return await response.json()
}

// not accepted by opponent yet
const useSentWagers = (page = 1, limit = 20) => {
    return useQuery({
        queryKey: ['sentWagers', limit],
        queryFn: () => fetchWagers(WagerStatus.Created, page, limit, { creator: true }),
    })
}

// not accepted by user yet
const useReceivedWagers = (page = 1, limit = 20) => {
    return useQuery({
        queryKey: ['receivedWagers', limit],
        queryFn: () => fetchWagers(WagerStatus.Created, page, limit, { creator: false }),
    })
}

const usePendingWagers = (page = 1, limit = 20) => {
    return useQuery({
        queryKey: ['pendingWagers', limit],
        queryFn: () => fetchWagers(WagerStatus.Pending, page, limit),
    })
}


const useClaimableWagers = (page = 1, limit = 20) => {
    return useQuery({
        queryKey: ['claimableWagers', limit],
        queryFn: () => fetchWagers(WagerStatus.Completed, page, limit, { winner: true }),
    })
}

const usePastWagers = (page = 1, limit = 20) => {
    return useQuery({
        queryKey: ['pastWagers', limit],
        queryFn: () => fetchWagers(WagerStatus.Claimed, page, limit),
    })
}

export { fetchWagers, useSentWagers, useReceivedWagers, useClaimableWagers, usePastWagers }
