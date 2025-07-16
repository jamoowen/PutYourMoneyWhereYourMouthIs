'use client'

import { useState } from 'react'
import Button from './common/button'
import { useAccount, useConfig, useConnect, useDisconnect, useSignMessage } from 'wagmi'
import { CreateConnectorFn, connect as coreConnect, disconnect as coreDisconnect } from '@wagmi/core'

import { supportedWallets as SUPPORTED_WALLETS } from '@/lib/blockchain'
import { useRouter } from 'next/navigation'

const SIGN_IN_STRING = "PYMWYMI_sign_in"
const SIGNED_UP = "SIGNED_UP"
const SIGNED_IN = "SIGNED_IN"

/**
 * @TODO  maybe we shouldnt enforce user to sign transaction agin 
 * if they are already have a token? then just connect wallet
 */
export default function SignInOptions({ onSignIn }: { onSignIn?: () => void }) {
  const router = useRouter()
  const { connect } = useConnect({ mutation: { onSuccess: signTransactionAndPost } })
  const { address, isConnected } = useAccount()
  const { signMessageAsync } = useSignMessage()
  const config = useConfig()

  const [signInErr, setSignInErr] = useState<string | null>(null)
  const [signInLoading, setSignInLoading] = useState(false)

  // this runs onSuccess via wallet connect props
  async function signTransactionAndPost() {
    try {
      const signature = await signMessageAsync({
        message: SIGN_IN_STRING,
      })

      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth`, {
        method: 'POST',
        body: JSON.stringify({ walletAddress: address, signature }),
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
      })

      const data = await res.json()
      if (!res.ok) {
        throw new Error(`ERROR signing in: ${res.status}: ${data}`)
      }
      if (data.authCode === SIGNED_UP) {
        const editProfileDialog = document.getElementById('edit_profile_modal') as HTMLDialogElement
        editProfileDialog.showModal()
      }
      router.refresh()
    } catch (error) {
      console.log(error)
      setSignInErr(`Failed to sign in`)
    } finally {
      setSignInLoading(false)
      const signInDialog = document.getElementById('sign_in_modal') as HTMLDialogElement
      signInDialog.close()

      if (onSignIn) {
        onSignIn()
      }
    }
  }

  async function handleSignIn(connectFn: CreateConnectorFn) {
    setSignInErr(null)
    setSignInLoading(true)
    try {
      if (isConnected) {
        await coreDisconnect(config)
      }
      await connect({ connector: connectFn })
    } catch (error) {
      console.error('Sign-in error:', error)
      setSignInErr(`Failed to sign in`)
      setSignInLoading(false)
    }
  }

  return (
    <>
      <h3 className="font-bold text-lg">Sign In</h3>
      {signInErr && (
        <p className="text-sm text-red-500 mt-3 mr-auto">{signInErr}</p>
      )}
      <ul className="py-2 mt-1 rounded-2xl border shadow-xl hover:bg-base-100 w-72 animate-fade-in">
        {SUPPORTED_WALLETS.map((wallet) => (
          <li key={wallet.name}>
            <Button
              isLoading={signInLoading}
              variant="none"
              onClick={() => handleSignIn(wallet.connector)}
              className="flex items-center justify-start w-full hover:cursor-pointer p-3 transition-all"
            >
              <div className="relative w-8 h-8 rounded-md">{wallet.image}</div>
              <span className="font-medium text-sm">{wallet.name}</span>
            </Button>
          </li>
        ))}
      </ul>
    </>
  )
}


