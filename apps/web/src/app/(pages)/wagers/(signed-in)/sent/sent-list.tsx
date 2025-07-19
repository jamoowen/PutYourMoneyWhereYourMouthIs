
'use client'

import { useCreatedWagers } from '../get-wagers'
import { WagerStatus } from '@/types/wager'
import { useState } from 'react'
import WagersList from '../wagers-list'
import { User } from '@/types/common'


export default function SentList({ user }: { user: User }) {
    const [page, setPage] = useState(1)
    const { data, isLoading, error } = useCreatedWagers(page, true)
    return (
        <div className='w-full'>
            <WagersList walletAddress={user.walletAddress} title='Wagers you sent' wagers={data?.data ?? []} />
        </div>

    )
}

