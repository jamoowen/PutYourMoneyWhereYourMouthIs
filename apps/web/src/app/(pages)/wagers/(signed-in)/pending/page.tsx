'use client'

import { useAccount, useConnect, useDisconnect } from 'wagmi'

export default function Page() {
  const account = useAccount()
  const { connectors, connect, status, error } = useConnect()
  const { disconnect } = useDisconnect()

  return (
    <div>
      ppending page
      alskdjfladjkspoqwip
      lkjfsdfjalsdkfj;lk
    </div>
  )
}
