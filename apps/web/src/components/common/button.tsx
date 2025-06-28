// src/components/ui/Button.tsx
import { ButtonHTMLAttributes } from 'react'
import clsx from 'clsx'
import { cn } from '@/lib/utils'

type Variant = 'primary' | 'secondary' | 'outline' | 'ghost' | 'none'

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: Variant
  className?: string
  isLoading?: boolean
  isDisabled?: boolean
}

export default function Button({
  variant = 'primary',
  className,
  isLoading,
  isDisabled,
  children,
  ...props
}: ButtonProps) {
  const base = 'flex items-center gap-2 px-5 py-2 rounded-lg font-medium transition-all duration-200 cursor-pointer'

  const variants: Record<Variant, string> = {
    none: '',
    primary: 'bg-gradient-to-r from-indigo-500 to-purple-600 text-foreground hover:shadow-lg',
    secondary: 'bg-gray-800 text-white hover:bg-gray-700',
    outline: 'border border-indigo-500 text-indigo-500 hover:text-foreground',
    ghost: 'bg-transparent text-gray-700 hover:bg-gray-100',
  }

  return (
    <button
      disabled={isLoading || isDisabled}
      className={cn(base, variants[variant], className)}
      {...props}
    >
      {isLoading ? <>
        <span className="loading loading-spinner"></span>
        loading
      </>
        : children}
    </button>
  )
}
