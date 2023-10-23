'use server'

import { cookies } from 'next/headers'

export async function ClearCookies() {
  cookies().delete('token')
}
