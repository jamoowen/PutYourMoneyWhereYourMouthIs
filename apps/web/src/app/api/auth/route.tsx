import { NextResponse } from 'next/server'

export async function DELETE() {
  const res = NextResponse.json({ success: true })
  console.log(`delete req`)

  // This removes the cookie
  res.cookies.set('pymwymi_auth_token', '', {
    httpOnly: true,
    path: '/',             // must match the original cookie path
    secure: true,          // match original if it was secure
    sameSite: 'lax',       // match original policy
    expires: new Date(0),  // remove it
  })

  return res
}
