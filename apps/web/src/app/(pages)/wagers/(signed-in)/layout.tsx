import Link from 'next/link'
import { cn, getAuthStatus } from '@/lib/utils' // optional utility for conditional classes
import NewWager from './new-wager'
import { cookies, headers } from 'next/headers'
import WagersTabs from './wagers-tabs'
import { getUser } from '@/lib/server-only-utils'

export default async function WagersLayout({
    children,
}: {
    children: React.ReactNode
}) {
    const user = await getUser()

    if (!user) {
        return null
    }

    return (
        <div className='w-full flex flex-col max-w-[500px] items-center'>
            <div className='w-full flex flex-col space-y-4'>
                <NewWager user={user} />
                <WagersTabs />
            </div>
            <div className="tabs justify-between w-full tabs-border mb-4">
            </div>
            <div className="p-4 border border-base-300 w-full bg-base-100 rounded-box">
                {children}
            </div>
        </div>
    )
}
