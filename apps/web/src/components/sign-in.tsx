
'use client'

import { useState } from 'react'
import Button from './common/button'
import { useChainId, useAccount, useConnect, useSignMessage } from 'wagmi'
import { CreateConnectorFn, signMessage } from '@wagmi/core'
import { getConfig } from '@/wagmi'

import { metaMask } from '@wagmi/connectors'
import { supportedWallets as SUPPORTED_WALLETS } from '@/lib/wallets'

const SIGN_IN_STRING = "PYMWYMI sign in"

export default function SignIn() {
  const { connect } = useConnect()
  const { address, connector, isConnected } = useAccount()
  const { signMessageAsync } = useSignMessage()

  async function handleSignIn() {
    try {
      if (!isConnected) {
        await connect({ connector: metaMask() })
      }

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

      // optionally handle response
    } catch (error) {
      console.error('Sign-in error:', error)
    }
  }

  return (
    <div className="relative z-50">
      <div className="relative z-50">
        <div className="dropdown  dropdown-start">
          <Button
            variant='primary'
            tabIndex={0}
            role="button"
          >
            <span className="font-medium text-sm">Sign In</span>
          </Button>

          <ul
            tabIndex={0}
            className="dropdown-content menu p-3 mt-4 bg-background rounded-2xl border  shadow-xl w-60  animate-fade-in"
          >
            {
              SUPPORTED_WALLETS.map((wallet) => (
                <li key={wallet.name}>
                  <button
                    onClick={() => handleSignIn(wallet.connector)}
                    className="flex items-center justify-start gap-3 hover:bg-muted-foreground p-3 rounded-xl transition-all"
                  >
                    <div className="relative w-8 h-8  rounded-md">
                      {wallet.image}
                    </div>
                    <span className="font-medium text-sm ">{wallet.name}</span>
                  </button>
                </li>
              ))
            }
          </ul>
        </div>
      </div>
    </div>
  )
}
