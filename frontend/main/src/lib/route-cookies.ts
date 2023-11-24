import { cookies as nextCookies } from 'next/headers'
import { MakotoCookiesInterface, CookieSetOptions } from '@makoto/lib/cookies'

class Cookies implements MakotoCookiesInterface {
  get(name: string): string | undefined {
    return nextCookies().get(name)?.value
  }

  set(name: string, value: string, opts?: CookieSetOptions): void {
    nextCookies().set(name, value, opts)
  }

  delete(name: string): void {
    nextCookies().delete(name)
  }
}

export const cookies = new Cookies()
