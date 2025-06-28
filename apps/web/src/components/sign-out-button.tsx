'use client'

import { useRouter } from "next/navigation"
import { useState } from "react"
import Button from "./common/button"

export default function SignOutButton() {
  const [isSignOutLoading, setIsSignOutLoading] = useState(false)
  const router = useRouter()

  async function handleSignOut() {
    setIsSignOutLoading(true)
    await fetch('/api/auth', {
      method: 'DELETE',
      credentials: 'include',
    })
    router.refresh()
  }

  return (
    <Button isLoading={isSignOutLoading} variant="none" type="submit" onClick={handleSignOut} className="hover:bg-base-200 rounded-md w-full text-left">
      Logout
    </Button>

  )

}
