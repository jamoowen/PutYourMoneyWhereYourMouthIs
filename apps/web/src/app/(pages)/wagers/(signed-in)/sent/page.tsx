'use client'

import Link from 'next/link'
import { useWagers } from '../get-wagers'
import { WagerStatus } from '@/types/wager'
import { useState } from 'react'

/**
 * @TODO add tabs for wagers - Completed, Invitations, Ongoing, Claimable?
 * @TODO add new wager button & form 
 */

// 

export default function Page() {
  const [page, setPage] = useState(1)
  const { data, isLoading, error } = useWagers(WagerStatus.Created, page)

  console.log(`err: ${error}, data: ${JSON.stringify(data)}`)
  return (
    <div>
      Sent wagers
      <div className="tabs tabs-border">
      </div>

    </div>

  )
}

