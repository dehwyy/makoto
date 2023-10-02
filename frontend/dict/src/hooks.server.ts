import { SetAccessTokenSession } from '$lib/api/set-tokens'
import { Tokens } from '$lib/const'
import type { Handle } from '@sveltejs/kit'

export const handle: Handle = async ({ event, resolve }) => {
	const authToken = event.cookies.get(Tokens.access)
	const refreshToken = event.cookies.get(Tokens.refresh)

	SetAccessTokenSession(authToken, refreshToken, event)

	return await resolve(event)
}
