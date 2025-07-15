import { getAuthStatus } from '@/lib/utils';
import { Authorisation } from '@/types/common';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export default async function Page() {

  const allCookies = await cookies()
  const token = allCookies.get('pymwymi_auth_token')?.value ?? null;
  const [user, auth] = getAuthStatus(token);
  if (auth === Authorisation.Authorised) {
    redirect("/wagers/received")
  }

  return (
    <div className='flex flex-col items-center'>
      Sign in to view your wagers or create new ones.
    </div>

  )
}

