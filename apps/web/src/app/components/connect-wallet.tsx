'use client'

import { useState } from 'react'
import Image from 'next/image'

export default function ConnectWallet() {
  const [address, setAddress] = useState<string | null>(null)

  const connectMetaMask = async () => {
    if (typeof window.ethereum === 'undefined') {
      alert('MetaMask is not installed. Please install it to continue.')
      return
    }

    try {
      const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' })
      setAddress(accounts[0])
    } catch (err) {
      console.error('MetaMask connection error:', err)
    }
  }

  return (
    <div className="relative z-50">
      <div className="dropdown dropdown-end">
        <div
          tabIndex={0}
          role="button"
          className="flex items-center gap-2 bg-gradient-to-r from-indigo-500 to-purple-600 text-white px-5 py-2 rounded-2xl shadow-md hover:shadow-lg transition-all duration-200"
        >
          {address ? (
            <span className="truncate max-w-[160px] text-sm font-medium">
              {address}
            </span>
          ) : (
            <>
              <span className="font-medium text-sm">Connect Wallet</span>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth="1.5"
                stroke="currentColor"
                className="w-5 h-5"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M15.75 6a3.75 3.75 0 00-7.5 0v3.75h-1.5A2.25 2.25 0 004.5 12v6A2.25 2.25 0 006.75 20.25h10.5A2.25 2.25 0 0019.5 18v-6a2.25 2.25 0 00-2.25-2.25h-1.5V6z"
                />
              </svg>
            </>
          )}
        </div>

        {!address && (
          <ul
            tabIndex={0}
            className="dropdown-content menu p-3 mt-2 bg-white rounded-2xl shadow-xl w-60 border border-gray-100 animate-fade-in"
          >
            <li>
              <button
                onClick={connectMetaMask}
                className="flex items-center justify-start gap-3 hover:bg-gray-100 p-3 rounded-xl transition-all"
              >
                <div className="relative w-8 h-8 bg-white rounded-md">
                  <Image
                    src="/metamask.svg"
                    alt="MetaMask"
                    fill
                    className="object-contain"
                  />
                </div>
                <span className="font-medium text-sm text-gray-800">MetaMask</span>
              </button>
            </li>
          </ul>
        )}
      </div>
    </div>
  )
}
