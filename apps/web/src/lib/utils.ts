import { User, Authorisation } from "@/types"
import { clsx, type ClassValue } from "clsx"
import { NextRequest } from "next/server"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

function parseJwt(token: string | null): User | null {
  if (!token) return null
  try {
    const payload = token.split('.')[1];
    const decoded = atob(payload); // base64 decode
    const json = JSON.parse(decoded);
    return {
      name: json.name,
      walletAddress: json.walletAddress,
      iss: json.iss,
      exp: json.exp
    }
  } catch (err) {
    return null;
  }
}

export function getAuthStatus(token: string | null): [User | null, Authorisation] {
  const user = parseJwt(token);
  if (user && user.exp < Date.now() / 1000) {
    return [null, Authorisation.Expired]
  } else if (user) {
    return [user, Authorisation.Authorised]
  } else {
    return [null, Authorisation.Unauthorised]
  }
}
