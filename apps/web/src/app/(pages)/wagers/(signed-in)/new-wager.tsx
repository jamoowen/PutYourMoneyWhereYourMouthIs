'use client'

import { useForm, useFieldArray } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { useEffect, useState } from 'react'
import { z } from 'zod'
import Button from '@/components/common/button'
import { categories } from '@/lib/wager-categories'  // Your categories array here
import Image from 'next/image'
import SignInOptions from '@/components/sign-in-options'
import contractJson from '@/contracts/WagerEscrow.json' assert { type: 'json' }
import { User } from '@/types'
import { Abi, Address, erc20Abi } from 'viem'
import { fromUSDCLarge, getTokenAddress, getChainId, getPYMWYMIContractAddress, fromUSDCSmall } from '@/lib/blockchain'
import { useAccount, useSwitchChain, useWriteContract, useReadContract, useBalance } from 'wagmi'
import { getConfig } from '@/wagmi'

type NewWagerPayload = {
  transactionHash: string
  name: string
  category: string
  description: string
  location: string
  stake: string         // e.g. amount in base units
  currency: string      // e.g. "USDC"
  chain: string
  participantsAddresses: string[]
}

// Create Zod enum dynamically from categories array
const categoryEnum = z.enum(categories)

// Ethereum address validation schema
const ethAddress = z
  .string()
  .min(42, 'That address looks too short')
  .max(42, 'That address looks too long')
  .refine((val) => val.startsWith('0x'), 'Address must start with 0x')

const newChallengeSchema = z.object({
  name: z.string().min(1, 'We need to call this challenge something???').max(50, `That name is long af, chill.`),
  category: categoryEnum,
  description: z.string().max(500, `Chill, thats enough info.`).optional(),
  location: z.string().max(100, `There's no way that is a real place.`).optional(),
  stake: z.number().min(1, 'Required').max(1_000_000, `You're not that rich. Be real`),
  currency: z.literal('USDC'),
  chain: z.literal('Base'),
  participantsAddresses: z
    .array(ethAddress)
    .min(1, 'At least one participant')
    .max(3, 'Maximum 3 participants allowed'),
})

type NewChallengeForm = z.infer<typeof newChallengeSchema>

function logError(error: any) {
  console.log(`ERROR INTERACTING WITH PYMWYMI: ${error}`)
}

export default function NewWager({ user }: { user: User }) {
  const [newWagerData, setNewWagerData] = useState<NewWagerPayload>({
    transactionHash: '',
    name: '',
    category: categories[0], // default to first category
    description: '',
    location: '',
    stake: "5",
    currency: 'USDC',
    chain: "Base",
    participantsAddresses: [], // One participant by default
  })
  const [stake, setStake] = useState(newWagerData.stake)
  const [isStakeApproved, setIsStakeApproved] = useState(false)
  const [submitError, setSubmitError] = useState<string | null>(null)

  const { address, isConnected, chainId } = useAccount()
  const {
    data: createWagerHash,
    writeContract: writeContractCreateWager,
    isSuccess: isCreateWagerSuccess,
    error: createWagerError,
    failureReason: createWagerFailureReason,
    status: createWagerStatus,
  } = useWriteContract({
    mutation: { onSuccess: postNewWagerResultToBackend }
  })
  const { data: approveHash, writeContract: writeContractApprove, isSuccess: isApproveSuccess, error: approveError } = useWriteContract({
    mutation: { onSuccess: interactWithPYMWYMIContract, onError: logError }
  })
  const { data: balance } = useBalance({
    address: address as `0x${string}`,
    token: getTokenAddress(newWagerData.currency, getChainId(newWagerData.chain)),
  })
  const { switchChain, switchChainAsync, chains, error: switchChainError, isPending: isSwitchChainPending } = useSwitchChain()

  const selectedChainId = getChainId(newWagerData.chain)
  const contractAddress = getPYMWYMIContractAddress()
  const tokenAddress = getTokenAddress(newWagerData.currency, selectedChainId)

  const { data: allowance } = useReadContract({
    address: tokenAddress,
    abi: erc20Abi,
    functionName: 'allowance',
    args: [address as `0x${string}`, contractAddress],
    query: {
      enabled: !!address && !!tokenAddress && !!contractAddress,
    },
  })

  const config = getConfig()

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    control,
    watch,
  } = useForm<NewChallengeForm>({
    resolver: zodResolver(newChallengeSchema),
    defaultValues: {
      name: '',
      category: categories[0], // default to first category
      description: '',
      location: '',
      stake: 5,
      currency: 'USDC',
      chain: "Base",
      participantsAddresses: [], // One participant by default
    },
  })

  const { fields, append, remove } = useFieldArray({
    control,
    //@ts-ignore
    name: 'participantsAddresses',
  })

  // 0 if no wallet connected - sign inn flow
  // 1. open transaction and force user to stake
  // 2. onsuccess, we want to post the result to our backend.
  // 3. backend ABSOLUTELY needs to receive this info
  // 4. 


  async function onSubmit(data: NewChallengeForm) {
    console.log('SUBMITTING', JSON.stringify(data))

    if (!isConnected || address !== user.walletAddress) {
      const dialog = document.getElementById('sign_in_modal_2') as HTMLDialogElement
      dialog.showModal()
      return
    }



    try {

      const stakeAmount = fromUSDCLarge(data.stake.toString())

      console.log(`Balance: `, balance?.value, BigInt(stakeAmount))
      if (balance == null || balance.value < BigInt(stakeAmount)) {
        throw new Error('Not enough balance')
      }

      console.log(`Writing to WITH PYMWYMI`, stakeAmount, tokenAddress, data.participantsAddresses)

      setNewWagerData({
        transactionHash: '',
        name: data.name,
        category: data.category,
        description: data.description ?? "",
        location: data.location ?? "",
        stake: stakeAmount,
        currency: data.currency,
        chain: data.chain,
        participantsAddresses: data.participantsAddresses
      })

      const selectedChainId = getChainId(data.chain)
      if (chainId !== selectedChainId) {
        switchChain({ chainId: selectedChainId })
      }

      const selectedTokenAddress = getTokenAddress(data.currency, selectedChainId)
      // Check allowance
      // if we need to approve allowance first thats fine
      // interactWithPYMWYMIContract is passed to its onSuccess
      if (!allowance || allowance < BigInt(stakeAmount)) {
        // if (true) {
        console.log(`Needs approvale`)
        writeContractApprove({
          address: tokenAddress,
          abi: erc20Abi,
          functionName: 'approve',
          args: [contractAddress, BigInt(stakeAmount)],
        })
        return
      } else {
        console.log(`writing contract: `, contractAddress, tokenAddress, data.participantsAddresses, stakeAmount)
        interactWithPYMWYMIContract()
      }
    } catch (err) {
      console.error('Create Wager Failed:', err)
    }
  }

  async function interactWithPYMWYMIContract() {
    console.log(`INTERACTING WITH PYMWYMI`, newWagerData.stake, tokenAddress, newWagerData.participantsAddresses)
    writeContractCreateWager({
      address: contractAddress,
      abi: contractJson.abi as Abi,
      functionName: 'createWager',
      args: [
        newWagerData.participantsAddresses,
        BigInt(newWagerData.stake),
        tokenAddress,
      ],
    })
  }

  async function postNewWagerResultToBackend() {
    console.log(`POSTING TO BACKEND`, createWagerHash)
    if (!isCreateWagerSuccess || !createWagerHash) {
      console.log(`status: ${createWagerStatus}, reason: ${createWagerFailureReason}, error: ${createWagerError?.message}`)
      return
    }
    console.log(`hash: ${createWagerHash}`)
    try {
      const payload = {
        ...newWagerData,
        transactionHash: createWagerHash
      }

      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/wager/create`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify(payload),
      })
      const responseData = await response.json()
      console.log(`RES: ${response.status}, ${response.statusText},${JSON.stringify(responseData)}`)
      if (!response.ok) throw new Error('Failed to create challenge')
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <div>
      <Button
        onClick={() => {
          if (!isConnected || address !== user.walletAddress) {
            const dialog = document.getElementById('sign_in_modal') as HTMLDialogElement
            dialog.showModal()
          } else {
            const dialog = document.getElementById('new_wager_modal') as HTMLDialogElement
            dialog.showModal()
          }
        }}
        variant="primary"
        tabIndex={0}
        role="button"
      >
        Create
      </Button>
      <dialog id="sign_in_modal_2" className="modal">
        <div className="modal-box max-h-[80vh] overflow-y-auto">
          <SignInOptions />
        </div>
      </dialog>
      <dialog id="new_wager_modal" className="modal">
        <div className="modal-box max-h-[80vh] overflow-y-auto">
          <form onSubmit={handleSubmit(onSubmit)} className="space-y-5">
            {/* Name */}
            <div>
              <label htmlFor="name" className="block  mb-1">
                Wager Name
              </label>
              <input
                id="name"
                {...register('name')}
                placeholder="The Thriller in Manila"
                className="input input-bordered w-full"
              />
              <p className="text-sm text-gray-500 mt-1 mb-3">
                Enter a descriptive name for your wager.
              </p>
              {errors.name && (
                <p className="text-red-500 text-sm mt-1">{errors.name.message}</p>
              )}
            </div>

            {/* Category */}
            <div>
              <label htmlFor="category" className="block  mb-1">
                Category
              </label>
              <select
                id="category"
                {...register('category')}
                className="select select-bordered w-full"
                defaultValue={categories[0]}
              >
                {categories.map((cat) => (
                  <option key={cat} value={cat}>
                    {cat}
                  </option>
                ))}
              </select>
              <p className="text-sm text-gray-500 mt-1 mb-3">
                Select the category that best fits your wager.
              </p>
              {errors.category && (
                <p className="text-red-500 text-sm mt-1">{errors.category.message}</p>
              )}
            </div>

            {/* Description */}
            <div>
              <label htmlFor="description" className="block  mb-1">
                Description
              </label>
              <textarea
                id="description"
                {...register('description')}
                placeholder="Description"
                className="textarea textarea-bordered w-full"
              />
              <p className="text-sm text-gray-500 mt-1 mb-3">
                An optional description for your wager. Maybe set some rules?
              </p>
              {errors.description && (
                <p className="text-red-500 text-sm mt-1">{errors.description.message}</p>
              )}
            </div>
            {/* Location */}
            <div>
              <label htmlFor="location" className="block  mb-1">
                Location
              </label>
              <input
                id="location"
                {...register('location')}
                placeholder="Location"
                className="input input-bordered w-full"
              />
              <p className="text-sm text-gray-500 mt-1 mb-3">
                An optional location. Helps us see where the hell all of you are challenging eachother.
              </p>
              {errors.location && (
                <p className="text-red-500 text-sm mt-1">{errors.location.message}</p>
              )}
            </div>

            {/* Stake */}
            <div>
              <label htmlFor="stake" className="block  mb-1">
                Stake
              </label>
              <input
                max={fromUSDCSmall(balance?.value.toString() ?? "0")}
                type='number'
                id="stake"
                {...register('stake', { valueAsNumber: true })}
                placeholder="Stake"
                className="input input-bordered w-full"
                onChange={(e) => setStake(e.target.value)}
              />
              <p className="text-sm text-gray-500 mt-1 mb-3">
                How much we putting down???
              </p>
              {errors.stake && (
                <p className="text-red-500 text-sm mt-1">{errors.stake.message}</p>
              )}
            </div>

            {/* Currency Picker */}
            <div>
              <label htmlFor="currency" className="block mb-1 font-medium text-sm text-gray-700">
                Currency
              </label>
              <div className="relative">
                <select
                  id="currency"
                  {...register('currency')}
                  className="select select-bordered w-full pl-10 cursor-not-allowed bg-gray-100 text-gray-700"
                  defaultValue="USDC"
                  disabled
                >
                  <option value="USDC">USDC</option>
                </select>
                <span className="absolute inset-y-0 left-0 flex items-center pl-3">
                  <Image src="/usdc.png" alt="USDC" width={20} height={20} />
                </span>
              </div>
              <p className="text-sm text-gray-500 mt-1 mb-3">USDC is the fixed currency for this wager.</p>
              {errors.currency && (
                <p className="text-red-500 text-sm mt-1">{errors.currency.message}</p>
              )}
            </div>

            {/* Chain Picker */}
            <div>
              <label htmlFor="chain" className="block mb-1 font-medium text-sm text-gray-700">
                Chain
              </label>
              <select
                id="chain"
                {...register('chain')}
                className="select select-bordered w-full cursor-not-allowed bg-gray-100 text-gray-700"
                defaultValue="Base"
                disabled
              >
                <option value="Base">Base</option>
              </select>
              <p className="text-sm text-gray-500 mt-1 mb-3">This wager will be deployed on the Base chain.</p>
              {errors.chain && (
                <p className="text-red-500 text-sm mt-1">{errors.chain.message}</p>
              )}
            </div>
            {/* Participants */}
            <div>
              <div className="flex justify-between items-center mb-2">
                <span className="">Participants</span>
                <Button
                  type="button"
                  onClick={() => append('')}
                  variant="outline"
                  className="btn-sm"
                  disabled={fields.length >= 3}
                >
                  + Add
                </Button>
              </div>

              <p className="text-sm text-gray-500 mb-3">
                Add up to 3 Ethereum wallet addresses starting with 0x.
              </p>

              {fields.map((field, index) => (
                <div key={field.id} className="flex gap-2 mb-3 items-center">
                  <label
                    htmlFor={`participantsAddresses.${index}`}
                    className="w-28 "
                  >
                    Participant {index + 1}
                  </label>
                  <input
                    id={`participantsAddresses.${index}`}
                    {...register(`participantsAddresses.${index}` as const)}
                    placeholder="0x..."
                    className="input input-bordered flex-grow"
                  />
                  <Button
                    type="button"
                    onClick={() => remove(index)}
                    variant="ghost"
                    className="btn-sm text-red-500"
                    disabled={fields.length === 1}
                  >
                    ✕
                  </Button>
                </div>
              ))}

              {errors.participantsAddresses && (
                Array.isArray(errors.participantsAddresses) ? (
                  errors.participantsAddresses.map((err, idx) =>
                    err ? (
                      <p key={idx} className="text-red-500 text-sm">
                        Participant {idx + 1}: {err?.message}
                      </p>
                    ) : null
                  )
                ) : (
                  <p className="text-red-500 text-sm">
                    {errors.participantsAddresses.message}
                  </p>
                )
              )}
            </div>
            <div className="modal-action">
              {/* <Button type="button" onClick={interactWithPYMWYMIContract} disabled={isSubmitting}> */}
              {/*   Approve */}
              {/* </Button> */}
              <Button type="button" onClick={handleSubmit(onSubmit)}
                disabled={
                  isSubmitting ||
                  balance == null ||
                  balance.value < BigInt(fromUSDCLarge(stake))
                }
              >
                Submit
              </Button>
              <Button
                type="button"
                onClick={() => {
                  const dialog = document.getElementById('new_wager_modal') as HTMLDialogElement
                  dialog.close()
                }}
                variant="ghost"
              >
                Cancel
              </Button>
            </div>
          </form>
        </div>
        <form method="dialog" className="modal-backdrop">
          <button>close</button>
        </form>
      </dialog>

      {/* sign in modal 2 */}

      <dialog id="sign_in_modal" className="modal">
        <div className="modal-box">
          <h3 className="font-bold text-lg"></h3>
          <p className="py-4">Your wallet is not connected - sign in again before creating a new wager.</p>
          <SignInOptions onSignIn={() => {
            const dialog = document.getElementById('new_wager_modal') as HTMLDialogElement
            dialog.showModal()

          }} />
        </div>
        <form method="dialog" className="modal-backdrop">
          <button>close</button>
        </form>
      </dialog>
    </div>
  )
}
