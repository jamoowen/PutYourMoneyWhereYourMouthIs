'use client'

import { useSentWagers } from '../lib/get-wagers'
import { useState } from 'react'
import WagersList from '../components/wagers-list'
import { User } from '@/types/common'

const title = "Wagers you have sent, that have not yet been accepted"

export default function SentList({ user }: { user: User }) {
    const [page, setPage] = useState(1)
    const { data, isLoading, error } = useSentWagers(page)

    return (
        <div className='w-full'>
            <WagersList walletAddress={user.walletAddress} title={title} wagers={data?.data ?? []} />
            <div className="join flex justify-center w-full ">
                {page > 1 && <button className="join-item btn" onClick={() => setPage(page - 1)}>«</button>}
                {page === 1 && data?.data && data.data.length > 0 && <button className="join-item btn">{page}</button>}
                {data?.pagination.more && <button className="join-item btn" onClick={() => setPage(page + 1)}>»</button>}
            </div>
        </div>
    )
}

