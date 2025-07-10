import MetaMaskImage from '@/images/metamask'
import { err, ok, Result } from '@/types'
import { metaMask } from '@wagmi/connectors'
import { Address, formatUnits, parseUnits } from 'viem'

export const supportedWallets = [
  { name: 'MetaMask', image: <MetaMaskImage />, connector: metaMask() },
]

export type SupportedChain = 'Base';

const chainIdMap: Record<SupportedChain, { mainnet: number; testnet: number }> = {
  Base: {
    mainnet: 8453,
    testnet: 84532,
  },
};

// https://developers.circle.com/stablecoins/usdc-contract-addresses
const tokenAddressMap: Record<string, Record<number, Address | null>> = {
  USDC: {
    8453: '0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913',
    84532: '0x036CbD53842c5426634e7929541eC2318f3dCF7e',
  }
}

/**
 * @description for parsing a human readable usd amount and returning a string of the smallest unit of usdc
 * @returns string => the smallest unit of usdc
 */
export function fromUSDCLarge(value: string): string {
  return parseUnits(value, 6).toString()
}

/**
 * @description for parsing a string of the smallest unit of usdc and returning a human readable usd amount
 * @returns string => the largest unit of usdc
 */
export function fromUSDCSmall(value: string): string {
  const bigIntVal = BigInt(value)
  return formatUnits(bigIntVal, 6)
}

type ValidChainId = 8453 | 84532;

export function getChainId(chain: SupportedChain): ValidChainId {
  const chainIds = chainIdMap[chain];
  const env = process.env.NEXT_PUBLIC_ENVIRONMENT;

  return env === 'dev'
    ? chainIds?.testnet as ValidChainId
    : chainIds?.mainnet as ValidChainId;
}

export function getTokenAddress(token: string, chainId: number): Address {
  const address = tokenAddressMap[token]?.[chainId]
  return address ?? "" as Address
}

export function getPYMWYMIContractAddress(): Address {
  const env = process.env.NEXT_PUBLIC_ENVIRONMENT;
  return env === "dev"
    ? "0x7c471fcf09959b8522760ca69bddf3c91900d834"
    : "0x7c471fcf09959b8522760ca69bddf3c91900d834"
}
