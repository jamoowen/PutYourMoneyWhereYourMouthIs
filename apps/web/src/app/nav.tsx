import { cookies } from 'next/headers'
import { getAuthStatus } from '@/lib/utils'
import Link from 'next/link'
import Notifications from '@/components/notifications'
import UserProfile from '@/components/user-profile'

export default async function Nav() {
  const allCookies = await cookies()
  const token = allCookies.get('pymwymi_auth_token')?.value ?? null;
  const [user] = getAuthStatus(token)

  return (
    <nav className="sticky top-0 z-50 w-full  mb-4 border-b shadow-sm">
      <div className="navbar shadow-sm">
        <div className="navbar-start">
          <Link href="/" className="text-lg font-bold">PYMWYMI</Link>
        </div>
        <div className="navbar-center">

          <Link href="/wagers" className=" text-lg font-bold cursor-pointer">Wagers</Link>
        </div>
        <div className="navbar-end">
          <UserProfile user={user} />
        </div>
      </div>
    </nav >
  )
}
