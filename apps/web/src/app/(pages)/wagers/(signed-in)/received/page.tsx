import { getUser } from '@/lib/server-only-utils'
import ReceivedList from './received-list'

export default async function Page() {
    const user = await getUser()
    if (!user) {
        return null
    }

    return (
        <div>
            <div className="tabs tabs-border">
                <ReceivedList user={user} />
            </div>
        </div>
    )
}

