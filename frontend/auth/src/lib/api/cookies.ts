import type { Cookies as CookiesType } from '@sveltejs/kit'

export class MakotoCookies {
	static setGlobal(cookies: CookiesType, key: string, value: string) {
		cookies.set(key, value, {
			path: '/',
			httpOnly: true
		})
	}
}
