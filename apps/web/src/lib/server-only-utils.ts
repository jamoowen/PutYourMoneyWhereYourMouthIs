
import { cookies } from "next/headers"
import { NextRequest } from "next/server"
import { getAuthStatus } from "./utils";
import { User } from "@/types/common";

export async function getUser(): Promise<User | null> {
    const allCookies = await cookies()
    const token = allCookies.get('pymwymi_auth_token')?.value ?? null;
    const [user] = getAuthStatus(token)
    return user
}
