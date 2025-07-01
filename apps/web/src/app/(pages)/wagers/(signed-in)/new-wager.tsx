'use client'

import { useForm, useFieldArray } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { z } from 'zod'
import Button from '@/components/common/button'
import { categories } from '@/lib/wager-categories'  // Your categories array here
import Image from 'next/image'
import SignInOptions from '@/components/sign-in-options'
import { useAccount, useConnect, useSignMessage } from 'wagmi'
import { User } from '@/types'

// Create Zod enum dynamically from categories array
const categoryEnum = z.enum(categories)

// Ethereum address validation schema
const ethAddress = z
  .string()
  .min(42, 'Address must be 42 characters')
  .max(42, 'Address must be 42 characters')
  .refine((val) => val.startsWith('0x'), 'Address must start with 0x')

const newChallengeSchema = z.object({
  name: z.string().min(1, 'We need to call this challenge something???').max(50, `That name is long af, chill.`),
  category: categoryEnum,
  description: z.string().max(500, `Chill, thats enough info`),
  location: z.string().min(1, 'Required').max(50, `Theres no ways thats a real place bro`),
  stake: z.number().min(1, 'Required').max(1_000_000, `You're not that rich. Be real`),
  currency: z.literal('USDC'),
  chain: z.literal('Base'),
  participantsAddresses: z
    .array(ethAddress)
    .min(1, 'At least one participant')
    .max(3, 'Maximum 3 participants allowed'),
})

type NewChallengeForm = z.infer<typeof newChallengeSchema>

export default function NewWager({ user }: { user: User }) {
  const { address, isConnected } = useAccount()

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    control,
  } = useForm<NewChallengeForm>({
    resolver: zodResolver(newChallengeSchema),
    defaultValues: {
      name: '',
      category: categories[0], // default to first category
      description: '',
      location: '',
      stake: 0,
      currency: 'USDC',
      chain: "Base",
      participantsAddresses: [''], // One participant by default
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
    try {
      // this screen is protected by sign in but we might not have wallet connected
      if (!isConnected || address !== user.walletAddress) {
        const dialog = document.getElementById('sign_in_modal_2') as HTMLDialogElement
        dialog.showModal()
        return
      }
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/challenges`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify(data),
      })
      if (!res.ok) throw new Error('Failed to create challenge')
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <div>
      <Button
        onClick={() => {
          // document.getElementById('new_wager_modal')?.showModal()
          const dialog = document.getElementById('new_wager_modal') as HTMLDialogElement
          dialog.showModal()
        }}
        variant="primary"
        tabIndex={0}
        role="button"
      >
        +
      </Button>
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
                id="stake"
                {...register('stake')}
                placeholder="Stake"
                className="input input-bordered w-full"
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
              <Button type="submit" disabled={isSubmitting}>
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
          <button>close</button>
        </form>
      </dialog>

      {/* sign in modal 2 */}
      <dialog id="sign_in_modal_2" className="modal">
        <div className={"modal-box w-[400px] absolute top-4 right-4 border bg-background border-muted shadow-lg"} >
          <h3 className="font-bold text-lg">Sign In</h3>
          <ul className="py-2 mt-1 rounded-2xl border shadow-xl hover:bg-base-100 w-72 animate-fade-in">
            <SignInOptions />
          </ul>
        </div>

        <form method="dialog" className="modal-backdrop">
          <button className="cursor-default">close</button>
        </form>
      </dialog>
    </div>
  )
}
