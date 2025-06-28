import Button from './common/button'
import { cookies } from 'next/headers'
import { getAuthStatus } from '@/lib/utils'
import { Authorisation } from '@/types'
import SignIn from './sign-in'
import Link from 'next/link'
import SignOutButton from './sign-out-button'

export default async function UserAccount() {
  const allCookies = await cookies()
  const token = allCookies.get('pymwymi_auth_token')?.value ?? null;
  const [user, auth] = getAuthStatus(token)

  if (auth !== Authorisation.Authorised) {
    return <SignIn />
  }

  return (
    <div className="relative z-50">
      <div className="dropdown dropdown-end">
        <Button variant="outline" tabIndex={0} role="button">
          <div className="flex flex-col text-sm max-w-[140px] truncate text-left">
            <span className="font-medium truncate">
              {user?.name || 'Unnamed User'}
            </span>
            <span className="text-muted-foreground text-xs truncate">
              {user?.walletAddress}
            </span>
          </div>
        </Button>

        <ul
          tabIndex={0}
          className="dropdown-content menu p-2 mt-2 bg-base-100 rounded-box w-52 shadow-xl"
        >
          <li>
            <SignOutButton />
          </li>
        </ul>
      </div>
    </div>
  )
}
