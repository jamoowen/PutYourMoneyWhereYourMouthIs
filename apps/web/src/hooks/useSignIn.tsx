'use client'
import { useState } from 'react'
import { useAccount, useConnect, useSignMessage } from 'wagmi'
import { CreateConnectorFn } from '@wagmi/core'

import { supportedWallets as SUPPORTED_WALLETS } from '@/lib/blockchain'
import { hashMessage } from 'viem'
import { useRouter } from 'next/navigation'

const SIGN_IN_STRING = "PYMWYMI_sign_in"
const SIGNED_UP = "SIGNED_UP"
const SIGNED_IN = "SIGNED_IN"

const router = useRouter()
const { connect } = useConnect({ mutation: { onSuccess: signTransactionAndPost } })
const { address, isConnected } = useAccount()
const { signMessageAsync } = useSignMessage()
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

// signTransactionAndPost runs onSuccess so the else block is fine
async function useSignIn(connectFn: CreateConnectorFn) {
  try {
    if (!isConnected) {
      const res = await connect({
        connector: connectFn,
      })
    } else {
      signTransactionAndPost()
    }
  } catch (error) {
    setSignInLoading(false)
  }
}

export { useSignIn as handleSignIn }
