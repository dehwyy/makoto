import type { RequestHandler } from '@sveltejs/kit'
import { SignUpFetch, GetFromRequest } from '$lib/api/typed-fetch'
import { AuthClient } from '$lib/api/twirp-client'
import { MakotoCookiesAutorization } from '$lib/api/cookies'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const req = await GetFromRequest(request, SignUpFetch)

	const { response, headers } = await AuthClient.signUp(req)

	MakotoCookiesAutorization.setToken(headers, cookies)

	console.log(response)

	return new Response(null, {
		status: 200
	})
}
