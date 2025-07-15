import { PaginatedResponse } from '@/types/common'
import { Wager, WagerStatus } from '@/types/wager'
import { queryOptions, useQuery } from '@tanstack/react-query'

const fetchWagers = async (status: number, page: number, limit: number): Promise<PaginatedResponse<Wager[]>> => {
  const url = process.env.NEXT_PUBLIC_API_URL + `/wager/list?status=${status}&page=${page}&limit=${limit}`
  const response = await fetch(url, {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
  })
  return await response.json()
}

const useWagers = (status: WagerStatus, page = 1, limit = 20) => {
  return useQuery({
    queryKey: ['posts', limit],
    queryFn: () => fetchWagers(status, page, limit),
  })
}

export { useWagers, fetchWagers }
