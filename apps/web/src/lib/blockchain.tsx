import MetaMaskImage from '@/images/metamask'
import { err, ok, Result } from '@/types'
import { metaMask } from '@wagmi/connectors'
import { Address, Chain, formatUnits, parseUnits } from 'viem'
import { WagerEscrowAddressBaseMainnet, WagerEscrowAddressBaseTestnet, ChainEnvironment, ChainName, USDCAddressBaseTestnet, USDCAdressBaseMainnet, ChainId } from './constants'

export const supportedWallets = [
  { name: 'MetaMask', image: <MetaMaskImage />, connector: metaMask() },
]

/**
 * @description for parsing a human readable usd amount and returning a string of the smallest unit of usdc
 * @returns string => the smallest unit of usdc
 */
export function toWeiUSDC(value: string): string {
  return parseUnits(value, 6).toString()
}

/**
 * @description for parsing a string of the smallest unit of usdc and returning a human readable usd amount
 * @returns string => the largest unit of usdc
 */
export function fromWeiUSDC(value: string): string {
  const bigIntVal = BigInt(value)
  return formatUnits(bigIntVal, 6)
}

// type ValidChainIds = 8453 | 84532
type ValidChainIds = typeof ChainId[keyof typeof ChainId]
export function getChainId(chain: string): ValidChainIds {
  const env = process.env.NEXT_PUBLIC_CHAIN_ENVIRONMENT;
  switch (chain) {
    case ChainName.Base:
      return env === ChainEnvironment.Mainnet
        ? ChainId.BaseMainnet
        : ChainId.BaseTestnet;
    default:
      return ChainId.BaseTestnet
  }
}

export function getTokenAddress(token: string, chainId: number): Address {
  switch (token) {
    case "USDC":
      return chainId === ChainId.BaseMainnet
        ? USDCAdressBaseMainnet
        : USDCAddressBaseTestnet
    default:
      return USDCAddressBaseTestnet
  }
}

export function getPYMWYMIContractAddress(): Address {
  const env = process.env.NEXT_PUBLIC_CHAIN_ENVIRONMENT;
  return env === ChainEnvironment.Mainnet
    ? WagerEscrowAddressBaseMainnet
    : WagerEscrowAddressBaseTestnet
}
