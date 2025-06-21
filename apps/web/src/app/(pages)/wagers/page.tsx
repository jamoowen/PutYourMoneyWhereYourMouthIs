'use client'

import Button from '@/components/common/button'
import SignIn from '@/components/sign-in'
import Link from 'next/link'

export default function Page() {
  return (
    <div>
      You have not used our app before.
      sign in with your wallet address
      <SignIn />
    </div>

  )
}

