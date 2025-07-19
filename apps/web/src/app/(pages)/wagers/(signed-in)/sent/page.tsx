import { useWagers } from '../get-wagers'
import { WagerStatus } from '@/types/wager'
import SentList from './sent-list'
import { getUser } from '@/lib/server-only-utils'


export default async function Page() {
  const user = await getUser()
  if (!user) {
    return null
  }
  return (
    <div>
      <div className="tabs tabs-border">
        <SentList user={user} />
      </div>
    </div>

  )
}

