
'use client'


import { useState } from 'react'
import Button from './common/button'
import { useAccount, useConnect, useSignMessage } from 'wagmi'
import { CreateConnectorFn } from '@wagmi/core'

import { supportedWallets as SUPPORTED_WALLETS } from '@/lib/wallets'
import { hashMessage } from 'viem'
import { useRouter } from 'next/navigation'

const SIGN_IN_STRING = "PYMWYMI_sign_in"
const SIGNED_UP = "SIGNED_UP"
const SIGNED_IN = "SIGNED_IN"

/**
 * @TODO if auth response is SIGNED_UP -> we need to prompt user with another form to add name
 */
export default function SignIn() {
  const router = useRouter()
  const { connect } = useConnect({ mutation: { onSuccess: signTransactionAndPost } })
  const { address, isConnected } = useAccount()
  const { signMessageAsync } = useSignMessage()
  const [signInErr, setSignInErr] = useState<string | null>(null)
  const [signInLoading, setSignInLoading] = useState(false)

  // this runds onSuccess via wallet connect props
  async function signTransactionAndPost(data: any) {
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
        console.log(`new user signed up`)
        //prompt user to enter their name 
      }
      router.refresh()
    } catch (error) {
      console.log(error)
      setSignInErr(`Failed to sign in`)
      setSignInLoading(false)
    } finally {
    }
  }

  async function handleSignIn(connectFn: CreateConnectorFn) {
    setSignInLoading(true)
    try {
      if (!isConnected) {
        const res = await connect({
          connector: connectFn,
        })
      } else {
        signTransactionAndPost("yes")
      }
    } catch (error) {
      console.error('Sign-in error:', error)
      setSignInErr(`Failed to sign in`)
      setSignInLoading(false)
    }
  }

  return (
    <>
      <Button
        // @ts-ignore
        onClick={() => document.getElementById('my_modal_2').showModal()}
        variant='outline'
        tabIndex={0}
        role="button"
      >
        <span className="font-medium text-sm">Sign In</span>
      </Button>

      <dialog id="my_modal_2" className="modal modal-middle bg-background">
        <div className="modal-box border bg-background border-muted">
          <p className="py-4">Sign in with the wallet you want to use to send and accept wagers.</p>
          <ul
            className="py-2 mt-1 rounded-2xl border  shadow-xl w-72  animate-fade-in"
          >
            {
              SUPPORTED_WALLETS.map((wallet) => (
                <li key={wallet.name}>
                  <Button
                    isLoading={signInLoading}
                    variant='none'
                    onClick={() => handleSignIn(wallet.connector)}
                    className=" flex items-center justify-start w-full hover:cursor-pointer p-3 transition-all"
                  >
                    <div className="relative w-8 h-8  rounded-md">
                      {wallet.image}
                    </div>
                    <span className="font-medium text-sm ">{wallet.name}</span>
                  </Button>
                </li>
              ))
            }
          </ul>

        </div>
        <form method="dialog" className="modal-backdrop">
          <button className='cursor-default'>close</button>
        </form>
      </dialog>
    </>

  )
}
