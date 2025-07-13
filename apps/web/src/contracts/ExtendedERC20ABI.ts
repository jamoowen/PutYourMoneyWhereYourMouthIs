import { erc20Abi as baseErc20Abi } from 'viem'

export const extendedErc20Abi = [
  ...baseErc20Abi,
  {
    name: 'increaseAllowance',
    type: 'function',
    stateMutability: 'nonpayable',
    inputs: [
      { name: 'spender', type: 'address' },
      { name: 'addedValue', type: 'uint256' }
    ],
    outputs: [{ name: '', type: 'bool' }]
  }
]
