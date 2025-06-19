'use client'

import { WAGERS_ROUTES } from '@/lib/wagers/constants'
import Link from 'next/link'

/**
 * @TODO add tabs for wagers - Completed, Invitations, Ongoing, Claimable?
 * @TODO add new wager button & form 
 */
export default function Page() {


  return (
    <div>

      <div className="tabs tabs-border">
        {WAGERS_ROUTES.map((route) => (
          <>
          <Link key={route.href} href={route.href}>
            <input type="radio" name={route.href} className="tab"  aria-label={route.label} />
            <div className="tab-content border-base-300 bg-base-100 p-10">{route.label}</div>
          </Link>
</>
        ))}
      </div>

    </div>

  )
}

