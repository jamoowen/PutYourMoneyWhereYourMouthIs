'use client'
import Button from './common/button'
import { User } from '@/types/common'
import SignInButton from './sign-in-button'
import { disconnect as coreDisconnect } from '@wagmi/core'

import { useRouter } from "next/navigation"
import { useState } from "react"
import { useAccount, useConfig } from 'wagmi'

export default function UserProfile({ user }: { user: User | null }) {

  const config = useConfig()
  const [isSignOutLoading, setIsSignOutLoading] = useState(false)
  const router = useRouter()
  const { isConnected } = useAccount()

  async function handleSignOut() {
    setIsSignOutLoading(true)
    if (isConnected) {
      await coreDisconnect(config)
    }
    await fetch('/api/auth', {
      method: 'DELETE',
      credentials: 'include',
    })
    router.refresh()
    setIsSignOutLoading(false)
  }


  return (
    <div className="relative z-50">
      {
        user === null ? (
          <SignInButton />
        )
          : (
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
                  <Button isLoading={isSignOutLoading} variant="ghost" type="submit" onClick={handleSignOut} className="">
                    Logout
                  </Button>
                </li>
                <li>
                  <Button variant="ghost" onClick={() => {
                    const dialog = document.getElementById('edit_profile_modal') as HTMLDialogElement
                    dialog.showModal()
                  }}>
                    Edit Profile
                  </Button>

                </li>
              </ul>
            </div>
          )
      }
    </div>
  )
}
