import Link from 'next/link'
import Notifications from './components/notifications'
import { WAGERS_ROUTES } from '@/lib/wagers/constants'
import ConnectWallet from './components/connect-wallet'

export default function Nav() {
  return (
    <nav className="sticky top-0 z-50 w-full px-2 mb-4 border-b shadow-sm">
      <div className="navbar shadow-sm">
        <div className="navbar-start">
          <div className="text-lg font-bold">PYMWYMI</div>
        </div>
        <div className="navbar-center">
          <div className="dropdown dropdown-hover">
            <Link href="/wagers"><div tabIndex={0} role="button" className=" text-lg font-bold btn btn-ghost ">Wagers</div></Link>
            <ul tabIndex={0} className="dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm">
              {WAGERS_ROUTES.map((route) => (
                <li key={route.href}><Link href={route.href}>{route.label}</Link></li>
              ))}
            </ul>
          </div>
        </div>
        <div className="navbar-end">
          <ConnectWallet />
          <Notifications />
        </div>
      </div>
    </nav >
  )
}
