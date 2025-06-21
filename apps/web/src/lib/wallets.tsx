import MetaMaskImage from '@/images/metamask'
import { metaMask } from '@wagmi/connectors'

export const supportedWallets = [
  { name: 'MetaMask', image: <MetaMaskImage />, connector: metaMask() },
]

