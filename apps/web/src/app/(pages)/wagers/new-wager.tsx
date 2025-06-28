'use client'

import { useForm, useFieldArray } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { z } from 'zod'
import Button from '@/components/common/button'
import { categories } from '@/lib/categories'  // Your categories array here

// Create Zod enum dynamically from categories array
const categoryEnum = z.enum(categories)

// Ethereum address validation schema
const ethAddress = z
  .string()
  .min(42, 'Address must be 42 characters')
  .max(42, 'Address must be 42 characters')
  .refine((val) => val.startsWith('0x'), 'Address must start with 0x')

export const newChallengeSchema = z.object({
  name: z.string().min(1, 'Required'),
  category: categoryEnum,
  description: z.string().min(1, 'Required'),
  location: z.string().min(1, 'Required'),
  stake: z.string().min(1, 'Required'),
  currency: z.string().min(1, 'Required'),
  participantsAddresses: z
    .array(ethAddress)
    .min(1, 'At least one participant')
    .max(3, 'Maximum 3 participants allowed'),
})

export type NewChallengeForm = z.infer<typeof newChallengeSchema>

export default function NewWager() {
  const [isOpen, setIsOpen] = useState(false)

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
      stake: '',
      currency: '',
      participantsAddresses: [''], // One participant by default
    },
  })

  const { fields, append, remove } = useFieldArray({
    control,
    name: 'participantsAddresses',
  })

  async function onSubmit(data: NewChallengeForm) {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/challenges`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify(data),
      })
      if (!res.ok) throw new Error('Failed to create challenge')
      setIsOpen(false)
    } catch (err) {
      console.error(err)
    }
  }

  return (
    <div>
      <Button
        onClick={() => {
          setIsOpen(true)
          // @ts-ignore
          document.getElementById('new_wager_modal')?.showModal()
        }}
        variant="primary"
        tabIndex={0}
        role="button"
      >
        +
      </Button>

      {isOpen && (
        <dialog id="new_wager_modal" className="modal modal-open">
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
                  placeholder="Wager name"
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
                  Provide a detailed description of the wager.
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
                  Specify the location where this wager applies.
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
                  Define the stake amount for the wager.
                </p>
                {errors.stake && (
                  <p className="text-red-500 text-sm mt-1">{errors.stake.message}</p>
                )}
              </div>

              {/* Currency */}
              <div>
                <label htmlFor="currency" className="block  mb-1">
                  Currency
                </label>
                <input
                  id="currency"
                  {...register('currency')}
                  placeholder="Currency"
                  className="input input-bordered w-full"
                />
                <p className="text-sm text-gray-500 mt-1 mb-3">
                  Specify the currency for the stake (e.g., ETH, USD).
                </p>
                {errors.currency && (
                  <p className="text-red-500 text-sm mt-1">{errors.currency.message}</p>
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
                  onClick={() => setIsOpen(false)}
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
      )}
    </div>
  )
}
