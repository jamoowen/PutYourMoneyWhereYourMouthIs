import { getUser } from '@/lib/server-only-utils'
import NewWager from './components/new-wager'
import WagersTabs from './components/wagers-tabs'

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
