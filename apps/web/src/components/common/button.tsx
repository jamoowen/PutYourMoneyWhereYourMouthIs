// src/components/ui/Button.tsx
import { ButtonHTMLAttributes } from 'react'
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
  const base = 'relative isolate inline-flex items-center gap-2 px-5 py-2 rounded-lg font-medium cursor-pointer'

  const variants: Record<Variant, string> = {
    none: '',
    primary: 'bg-gradient-to-r from-indigo-500 to-purple-600 text-foreground hover:outline',
    secondary: 'bg-gray-800 text-white hover:bg-gray-700  bg-gradient-to-br from-indigo-500 to-purple-600 bg-clip-text text-transparent hover:outline',
    outline: cn(
      'text-indigo-500 bg-transparent',
      'before:absolute before:inset-0 before:rounded-lg before:p-[1px] before:bg-gradient-to-r before:from-indigo-500 before:to-purple-600 before:z-[-1] before:content-[""]',
      'after:absolute after:inset-[1px] after:rounded-lg after:bg-background after:z-[-1]',
      'hover:outline'
    ),
    ghost: 'bg-transparent text-gray-400 hover:bg-gray-100 hover:text-indigo-500',
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
