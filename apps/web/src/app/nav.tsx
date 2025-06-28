import Link from 'next/link'
import Notifications from '@/components/notifications'
import SignIn from '@/components/sign-in'
import UserAccount from '@/components/user-profile'

export default async function Nav() {

  return (
    <nav className="sticky top-0 z-50 w-full  mb-4 border-b shadow-sm">
      <div className="navbar shadow-sm">
        <div className="navbar-start">
          <Link href="/" className="text-lg font-bold">PYMWYMI</Link>
        </div>
        <div className="navbar-center">

          <Link href="/wagers" className=" text-lg font-bold cursor-pointer">Wagers</Link>
          {/* <div className="dropdown dropdown-hover"> */}
          {/*   <Link href="/wagers"><div tabIndex={0} role="button" className=" text-lg font-bold btn btn-ghost ">Wagers</div></Link> */}
          {/*   <ul tabIndex={0} className="dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm"> */}
          {/*     {WAGERS_ROUTES.map((route) => ( */}
          {/*       <li key={route.href}><Link href={route.href}>{route.label}</Link></li> */}
          {/*     ))} */}
          {/*   </ul> */}
          {/* </div> */}
        </div>
        <div className="navbar-end">
          <UserAccount />
        </div>
      </div>
    </nav >
  )
}
