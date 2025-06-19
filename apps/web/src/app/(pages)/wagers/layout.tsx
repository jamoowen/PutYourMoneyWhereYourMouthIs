'use client'

import { useSelectedLayoutSegment } from 'next/navigation'
import Link from 'next/link'
import { cn } from '@/lib/utils' // optional utility for conditional classes
import { WAGERS_ROUTES } from '@/lib/wagers/constants'



export default function WagersLayout({
  children,
}: {
  children: React.ReactNode
}) {
  const segment = useSelectedLayoutSegment()

  return (
    <div>
      <div className="tabs tabs-border mb-4">
        {WAGERS_ROUTES.map((route) => {
          const path = route.href.split('/').pop()
          const isActive = segment === path

          return (
            <Link
              key={route.href}
              href={route.href}
              className={cn('tab', isActive && 'tab-active')}
            >
              {route.label}
            </Link>
          )
        })}
      </div>
      <div className="p-4 border border-base-300 bg-base-100 rounded-box">
        {children}
      </div>
    </div>
  )
}
