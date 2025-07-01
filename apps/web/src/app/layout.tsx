import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

import { headers } from 'next/headers'
import { type ReactNode } from 'react'
import { cookieToInitialState } from 'wagmi'
import { getConfig } from '../wagmi'
import { Providers } from './providers'

import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'
import Nav from "./nav";
import EditProfile from "@/components/edit-profile";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Put Your Money Where Your Mouth Is",
  description: "Trustless wagers and betting amongst friends",
};
const queryClient = new QueryClient()

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  const initialState = cookieToInitialState(
    getConfig(),
    (await headers()).get('cookie'),
  )

  return (
    <html lang="en" className="dark">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <Providers initialState={initialState}>
          <main className="flex p-2 flex-col items-center h-[100vh] justify-items-center bg-background text-white">
            <Nav />
            <EditProfile />
            {children}
          </main>
        </Providers>
      </body>
    </html>
  );
}
