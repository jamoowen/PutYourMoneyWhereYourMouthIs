import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';
import { Authorisation, User } from '@/types';
import { json } from 'stream/consumers';
import { getAuthStatus } from './lib/utils';

export function middleware(request: NextRequest) {

  const token = request.cookies.get('pymwymi_auth_token')?.value ?? null;
  const [user, auth] = getAuthStatus(token);

  const path = request.nextUrl.pathname

  // Protect a route (example)
  const isProtected = path.startsWith('/wagers/') && path.length > 8;
  if (isProtected && auth != Authorisation.Authorised) {
    return NextResponse.redirect(new URL('/', request.url));
  }

  const response = NextResponse.next();
  // Optionally inject info into request headers (for server-side access)
  if (auth == Authorisation.Authorised) {
    const u = btoa(JSON.stringify({ name: user?.name, walletAddress: user?.walletAddress }))
    response.headers.set('pymwymi-user', u);
  }

  return response;
}
