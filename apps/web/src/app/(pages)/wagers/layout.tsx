'use client'

import { useSelectedLayoutSegment } from 'next/navigation'
import Link from 'next/link'
import { cn } from '@/lib/utils' // optional utility for conditional classes

const WAGERS_ROUTES = [
  { href: '/wagers/pending', label: 'In Progress' },
  { href: '/wagers/claimable', label: 'Claimable' },
  { href: '/wagers/received', label: 'Received' },
  { href: '/wagers/sent', label: 'Sent' },
  { href: '/wagers/history', label: 'History' },
]

export default function WagersLayout({
  children,
}: {
  children: React.ReactNode
}) {
  const segment = useSelectedLayoutSegment()

  return (
    <div className='w-full flex flex-col max-w-[500px] items-center'>
      <div className="tabs justify-between w-full tabs-border mb-4">
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
      <div className="p-4 border border-base-300 w-full bg-base-100 rounded-box">
        {children}
      </div>
    </div>
  )
}
