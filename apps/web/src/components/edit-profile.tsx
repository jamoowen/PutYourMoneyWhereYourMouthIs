'use client'

import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import Button from '@/components/common/button'
import { useState } from 'react'
import { useRouter } from 'next/navigation'

const schema = z.object({
  name: z.string().min(1, 'Name is required').max(50, 'That name is way too long'),
})

type FormData = z.infer<typeof schema>

export default function EditProfile() {
  const router = useRouter()

  const [isSubmitting, setIsSubmitting] = useState(false)
  const [editNameErr, setEditNameErr] = useState<string | null>(null)

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(schema),
  })

  async function onSubmit(data: FormData) {
    try {
      setIsSubmitting(true)
      console.log(`submitting form: `, data)
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth`, {
        method: 'PATCH',
        body: JSON.stringify(data),
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
      })
      if (!res.ok) throw new Error('Update failed')
      setEditNameErr(null)
      const dialog = document.getElementById('edit_profile_modal') as HTMLDialogElement
      dialog.close()
      router.refresh()

      // Optionally close modal or show success
    } catch (err) {
      console.error(err)
      setEditNameErr(`Failed to update name`)
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <>
      <dialog id="edit_profile_modal" className="modal">
        <div className="modal-box ">
          <h3 className="text-lg font-bold mb-4">Update Your Name</h3>
          {editNameErr && (
            <p className="text-sm text-red-500 mt-3 mr-auto">{editNameErr}</p>
          )}

          <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
            <div>
              <label htmlFor="name" className="block mb-1 font-medium">
                Name
              </label>
              <input
                id="name"
                {...register('name')}
                placeholder="Your new name"
                className="input input-bordered w-full"
              />
              {errors.name && (
                <p className="text-red-500 text-sm mt-1">{errors.name.message}</p>
              )}
            </div>

            <div className="modal-action">
              <Button type="submit" disabled={isSubmitting} isLoading={isSubmitting}>
                Save
              </Button>
              <Button
                type="button"
                variant="ghost"
                onClick={() => {
                  const dialog = document.getElementById('edit_profile_modal') as HTMLDialogElement
                  dialog.close()
                }}
              >
                Cancel
              </Button>
            </div>
          </form>
        </div>

        <form method="dialog" className="modal-backdrop">
          <button className="cursor-pointer">close</button>
        </form>
      </dialog>
    </>
  )
}
