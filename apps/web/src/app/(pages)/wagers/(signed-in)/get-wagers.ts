import { queryOptions } from '@tanstack/react-query'

export const getWagers = queryOptions({
  queryKey: ['wagers'],
  queryFn: async () => {
    const url = process.env.NEXT_PUBLIC_API_URL + "/wager/list"
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    })

    return response.json()
  },
})

type PersistedWager =  newwa{
}

const fetchWagers = async (limit = 10): Promise<Array<Post>> => {
  const response = await fetch('https://jsonplaceholder.typicode.com/posts')
  const data = await response.json()
  return data.filter((x: Post) => x.id <= limit)
}

const usePosts = (limit: number) => {
  return useQuery({
    queryKey: ['posts', limit],
    queryFn: () => fetchPosts(limit),
  })
}

export { usePosts, fetchPosts }
