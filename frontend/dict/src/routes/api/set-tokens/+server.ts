import { SetTokens, SetTokensCookies } from '$lib/api/set-tokens'
import type { RequestHandler } from './$types'

export const POST: RequestHandler = async event => {
	const { request } = event
	const { access_token, refresh_token } = await request.json()

	SetTokensCookies({ access: access_token, refresh: refresh_token, cookies: event.cookies })

	return new Response(null, { status: 200 })
}
