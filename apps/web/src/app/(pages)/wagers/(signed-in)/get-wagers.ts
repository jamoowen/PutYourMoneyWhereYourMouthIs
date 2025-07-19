import { PaginatedResponse } from '@/types/common'
import { Wager, WagerStatus } from '@/types/wager'
import { queryOptions, useQuery } from '@tanstack/react-query'

export type ExtraQueryOptions = {
    creator?: boolean
    winner?: boolean
}

const fetchWagers = async (status: number, page: number, limit: number, options?: ExtraQueryOptions): Promise<PaginatedResponse<Wager[]>> => {
    let url = process.env.NEXT_PUBLIC_API_URL + `/wager/list?status=${status}&page=${page}&limit=${limit}`
    if (options?.creator) {
        url += `&creator=${options.creator}`
    }
    if (options?.winner) {
        url += `&winner=${options.winner}`
    }

    const response = await fetch(url, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
    })
    return await response.json()
}

const useWagers = (status: WagerStatus, page = 1, limit = 20) => {
    return useQuery({
        queryKey: ['wagers', limit],
        queryFn: () => fetchWagers(status, page, limit),
    })
}

const useCreatedWagers = (page = 1, creator: boolean, limit = 20) => {
    return useQuery({
        queryKey: ['createdWagers', limit],
        queryFn: () => fetchWagers(WagerStatus.Created, page, limit, { creator }),
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

export { fetchWagers, useWagers, useCreatedWagers, useClaimableWagers, usePastWagers }
