import MetaMaskImage from '@/images/metamask'
import { err, ok, Result } from '@/types'
import { metaMask } from '@wagmi/connectors'
import { Address, formatUnits, parseUnits } from 'viem'

export const supportedWallets = [
  { name: 'MetaMask', image: <MetaMaskImage />, connector: metaMask() },
]

type SupportedChain = 'Base';

const chainIdMap: Record<SupportedChain, { mainnet: number; testnet: number }> = {
  Base: {
    mainnet: 8453,
    testnet: 84532,
  },
};

const tokenAddressMap: Record<string, Record<number, string | null>> = {
  USDC: {
    8453: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
    84532: '0x...YourUSDCOnBaseSepolia',
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

export function getChainId(chain: SupportedChain): Result<ValidChainId, Error> {
  const chainIds = chainIdMap[chain];
  if (!chainIds) return err(new Error(`Chain ${chain} not found`));

  const env = process.env.NEXT_PUBLIC_ENVIRONMENT;
  if (env === 'dev') return ok(chainIds.testnet as ValidChainId);
  if (env === 'prod') return ok(chainIds.mainnet as ValidChainId);

  return err(new Error(`Invalid environment: ${env}`));
}

export function getTokenAddress(token: string, chainId: number): Result<string, Error> {
  const addresses = tokenAddressMap[token]
  if (addresses == null) return err(new Error(`Token ${token} not found`))

  const address = addresses[chainId]
  if (address == null) return err(new Error(`Token ${token} not found on chain ${chainId}`))

  return ok(address)
}

export function getPYMWYMIContractAddress(): Result<Address, Error> {
  if (process.env.NEXT_PUBLIC_ENVIROMENT == 'dev') return ok("0x7c471fcf09959b8522760ca69bddf3c91900d834")
  else if (process.env.NEXT_PUBLIC_ENVIROMENT == 'prod') return ok("0x7c471fcf09959b8522760ca69bddf3c91900d834")

  return err(new Error(`Environment ${process.env.NEXT_PUBLIC_ENVIROMENT} not found`))
}
