import { http, cookieStorage, createConfig, createStorage } from 'wagmi'
import { base, baseSepolia } from 'wagmi/chains'
import { metaMask } from 'wagmi/connectors'

export function getConfig() {
  return createConfig({
    chains: [base, baseSepolia],
    connectors: [
      metaMask()
    ],
    storage: createStorage({
      storage: cookieStorage,
    }),
    ssr: true,
    transports: {
      [base.id]: http(),
      [baseSepolia.id]: http(),
    },
  })
}

declare module 'wagmi' {
  interface Register {
    config: ReturnType<typeof getConfig>
  }
}
