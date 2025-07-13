
export type User = {
  name: string;
  walletAddress: string;
  iss: string,
  exp: number
}

export enum Authorisation {
  Authorised = "authorised",
  Unauthorised = "unauthorised",
  Expired = "expired",
}

export type Result<T, E> =
  | { ok: true; value: T }
  | { ok: false; error: E }

export const ok = <T>(value: T): Result<T, never> => ({ ok: true, value })
export const err = <E>(error: E): Result<never, E> => ({ ok: false, error })

export class WagerError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "WagerError";

    // Set the prototype explicitly to maintain instanceof checks
    Object.setPrototypeOf(this, WagerError.prototype);
  }
}
