import { Tokens, USER_ID } from '$lib/const'
import type { Cookies } from '@sveltejs/kit'
import type { CookieSerializeOptions } from 'cookie'

interface AuthCookies {
	access_token: string
	refresh_token: string
	user_id: string
	cookies: Cookies
}

export const setAuthCookies = ({ access_token, refresh_token, user_id, cookies }: AuthCookies) => {
	//
	const params: CookieSerializeOptions = { httpOnly: true, path: '/' }

	const set = (key: string, value: string) => {
		cookies.set(key, value, params)
	}

	set(Tokens.access, access_token)
	set(Tokens.refresh, refresh_token)
	set(USER_ID, user_id)
}
