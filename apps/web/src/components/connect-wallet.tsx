'use client'

import Image from 'next/image'
import Button from './common/button'
import MetaMaskImage from '@/images/metamask'
import { useConnect, useAccount, useDisconnect, useEnsName, useEnsAvatar, CreateConnectorFn } from 'wagmi'
import { metaMask } from '@wagmi/connectors'
import { supportedWallets } from '@/lib/wallets'


function ConnectWalletOrShowConnected() {
  const { isConnected } = useAccount()
  if (isConnected) {
    return <ConnectedWallet />
  }
  return <ConnectWallet />
}

function ConnectedWallet() {
  const { address } = useAccount()
  const { disconnect } = useDisconnect()
  const { data: ensName } = useEnsName({ address })
  const { data: ensAvatar } = useEnsAvatar({ name: ensName! })

  return (
    <div className="relative z-50">
      <div className="dropdown  dropdown-end">
        <Button
          variant='outline'
          tabIndex={0}
          role="button"
        >
          {ensAvatar ? (
            <Image
              src={ensAvatar}
              alt="ENS Avatar"
              width={32}
              height={32}
              className="rounded-full border"
            />
          ) : (
            null
          )}
          <div className="flex flex-col text-sm max-w-[100px] truncate">
            <span className="text-muted-foreground truncate">
              {address}
            </span>
          </div>
        </Button>

        <ul
          tabIndex={0}
          className="dropdown-content menu  shadow-xl w-60  animate-fade-in"
        >
          <li >
            <Button
              variant="none"
              className=" text-red-500 bg-transparent hover:bg-transparent "
              onClick={() => disconnect()}
            >
              Disconnect
            </Button>
          </li>
        </ul>
      </div>
    </div>

  )
}

function ConnectWallet({ }: {}) {
  const { connect } = useConnect();

  function handleConnect(connectFn: CreateConnectorFn) {
    connect({ connector: connectFn })
  }

  return (
    <div className="relative z-50">
      <div className="dropdown  dropdown-end">
        <Button
          variant='outline'
          tabIndex={0}
          role="button"
        >
          <span className="font-medium text-sm">Connect wallet</span>
        </Button>

        <ul
          tabIndex={0}
          className="dropdown-content menu p-3 mt-4 bg-background rounded-2xl border  shadow-xl w-60  animate-fade-in"
        >
          {
            supportedWallets.map((wallet) => (
              <li key={wallet.name}>
                <button
                  onClick={() => handleConnect(wallet.connector)}
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
  )
}

export { ConnectWalletOrShowConnected, ConnectWallet, ConnectedWallet }
