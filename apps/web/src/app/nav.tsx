import Link from 'next/link'

export default function Nav() {
  return (
    <nav className="sticky top-0 z-50 w-full px-6 py-4 bg-white border-b shadow-sm">
      <div className="flex items-center gap-4">
        <Link
          href="/"
          className="text-sm font-medium text-gray-800 hover:text-black transition"
        >
          Home
        </Link>
        <Link
          href="/wagers"
          className="text-sm font-medium text-gray-800 hover:text-black transition"
        >
          Wagers
        </Link>
      </div>
    </nav>
  )
}
