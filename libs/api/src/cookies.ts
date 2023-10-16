import type { RpcMetadata } from '@protobuf-ts/runtime-rpc/build/types/rpc-metadata'
import type { Cookies as CookiesType } from '@sveltejs/kit'

export class MakotoCookies {
	static setGlobal(cookies: CookiesType, key: string, value: string) {
		cookies.set(key, value, {
			path: '/',
			httpOnly: true
		})
	}
}

export class MakotoCookiesAutorization {
	static setToken(headers: RpcMetadata, cookies: CookiesType) {
		const split_token = (headers['authorization'] as string).split(' ')
		if (split_token.length < 2) return

		MakotoCookies.setGlobal(cookies, 'token', split_token[1])
	}
}
