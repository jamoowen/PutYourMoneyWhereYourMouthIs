'use client'

import { useAccount, useReadContract, useSimulateContract, useWriteContract } from 'wagmi'
import { useEffect } from 'react'
import { erc20Abi } from 'viem'

interface UseApproveIfNeededProps {
  token: `0x${string}`
  spender: `0x${string}`
  amount: bigint
  onApproved?: () => void
}

export function useApproveIfNeeded({ token, spender, amount, onApproved }: UseApproveIfNeededProps) {
  const { address } = useAccount()

  const {
    data: allowance,
    refetch: refetchAllowance,
    isLoading: isAllowanceLoading,
  } = useReadContract({
    address: token,
    abi: erc20Abi,
    functionName: 'allowance',
    args: [address!, spender],
    query: {
      enabled: !!address,
    },
  })

  const {
    data: simulation,
    isSuccess: canApprove,
  } = useSimulateContract({
    address: token,
    abi: erc20Abi,
    functionName: 'approve',
    args: [spender, amount],
    query: {
      enabled: !!address && !!allowance && allowance < amount,
    },
  })

  const {
    writeContract,
    isPending,
    isSuccess,
    data: txHash,
  } = useWriteContract()

  useEffect(() => {
    if (!address || allowance === undefined) return

    if (allowance < amount && simulation?.request) {
      console.log('🔐 Sending approval transaction...')
      writeContract(simulation.request)
    } else if (allowance >= amount) {
      console.log('✅ Already approved')
      onApproved?.()
    }
  }, [allowance, simulation?.request, address])

  useEffect(() => {
    if (isSuccess) {
      console.log('✅ Approval transaction succeeded:', txHash)
      refetchAllowance()
      onApproved?.()
    }
  }, [isSuccess])

  return {
    isLoading: isAllowanceLoading || isPending,
    isApproved: allowance !== undefined && allowance >= amount,
    isApproving: isPending,
  }
}
