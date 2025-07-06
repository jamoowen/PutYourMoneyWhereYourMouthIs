
// src/components/ui/Text.tsx
import { HTMLAttributes } from 'react'
import clsx from 'clsx'
import { cn } from '@/lib/utils'

type Variant = 'primary' | 'secondary' | 'none'

interface TextProps extends HTMLAttributes<HTMLElement> {
  variant?: Variant
  className?: string
  isLoading?: boolean
  isDisabled?: boolean
}

export default function Text({
  variant = 'primary',
  className,
  isLoading,
  isDisabled,
  children,
  ...props
}: TextProps) {
  const base = ''

  const variants: Record<Variant, string> = {
    none: '',
    primary: 'text-lg font-bold  bg-gradient-to-br from-indigo-500 to-purple-600 bg-clip-text text-transparent',
    secondary: 'text-lg font-bold  bg-clip-text text-transparent hover:bg-gradient-to-br hover:from-indigo-500 hover:to-purple-600 hover:bg-clip-text hover:text-transparent',
  }

  return (
    <span
      className={cn(base, variants[variant], className)}
      {...props}
    >
      {children}
    </span>
  )
}
