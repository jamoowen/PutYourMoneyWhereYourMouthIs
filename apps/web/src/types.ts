
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

