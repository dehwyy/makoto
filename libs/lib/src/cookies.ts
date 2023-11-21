export interface CookieSetOptions {
	httpOnly?: boolean
	path?: string
}

export interface MakotoCookiesInterface {
	get(name: string): string | undefined;
	set(name: string, value: string, opts?: CookieSetOptions): void;
	delete(name: string): void;
}

export class MakotoCookies {
	static setGlobal(cookies: MakotoCookiesInterface, key: string, value: string) {
		cookies.set(key, value, {
			path: '/',
			httpOnly: true
		})
	}

	static delete(cookies: MakotoCookiesInterface, key: string) {
		cookies.delete(key)
	}
}
