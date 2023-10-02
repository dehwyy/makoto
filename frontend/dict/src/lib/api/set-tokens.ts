import { setSession } from '$houdini'
import { Tokens } from '$lib/const'
import type { Cookies, RequestEvent } from '@sveltejs/kit'

type EventT = RequestEvent<Partial<Record<string, string>>, string | null>

interface SetTokensCookies {
	refresh: string
	access: string
	cookies: Cookies
}

interface SetTokens {
	refresh: string
	access: string
	event: EventT
}

// should be run only on server
export function SetTokens({ access, refresh, event }: SetTokens) {
	SetTokensCookies({ access, refresh, cookies: event.cookies })
	SetAccessTokenSession(access, refresh, event)
}

export function SetAccessTokenSession(
	access_token: string | undefined,
	refresh_token: string | undefined,
	event: EventT
) {
	setSession(event, { token: access_token, refresh_token })
}

export function SetTokensCookies({ access, refresh, cookies }: SetTokensCookies) {
	cookies.set(Tokens.access, access, { httpOnly: true, path: '/' })
	cookies.set(Tokens.refresh, refresh, { httpOnly: true, path: '/' })
}
