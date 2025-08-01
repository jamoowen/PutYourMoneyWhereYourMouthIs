'use client'

import { cn } from "@/lib/utils"
import Link from "next/link"
import { usePathname } from "next/navigation"

const WAGERS_ROUTES = [
    { href: '/wagers/pending', label: 'Pending' },
    { href: '/wagers/claimable', label: 'Claimable' },
    { href: '/wagers/received', label: 'Received' },
    { href: '/wagers/sent', label: 'Sent' },
    { href: '/wagers/history', label: 'History' },
]

export default function WagersTabs() {
    const pathname = usePathname().split('/').pop()
    return (
        <div className="flex  justify-between">
            {WAGERS_ROUTES.map((route) => {
                const path = route.href.split('/').pop()
                const isActive = pathname === path
                console.log(`isActive: ${isActive}`)

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
    )
}
