import Image, { ImageProps } from "next/image"
import { cn } from "@/lib/utils" // optional if using Tailwind Merge or clsx

interface MetaMaskImageProps extends Omit<ImageProps, "src" | "alt"> {
  className?: string
}

export default function MetaMaskImage({ className, ...props }: MetaMaskImageProps) {
  return (
    <Image
      src="/metamask.svg"
      alt="MetaMask"
      width={40}
      height={40}
      className={cn("object-contain", className)}
      {...props}
    />
  )
}
