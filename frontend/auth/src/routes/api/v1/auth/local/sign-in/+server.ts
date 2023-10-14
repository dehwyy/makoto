import type { RequestHandler } from '@sveltejs/kit'
import { AuthClient } from '$lib/api/twirp-client'
import { RpcInterceptors } from '$lib/api/prc-interceptors'
import { GetFromRequest, SignInFetch } from '$lib/api/typed-fetch'
import { MakotoCookiesAutorization } from '$lib/api/cookies'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const req = await GetFromRequest(request, SignInFetch)

	const { headers, response } = await AuthClient.signIn({
		authMethod: {
			oneofKind: 'credentials',
			credentials: {
				password: req.password,
				uniqueIdentifier: {
					oneofKind: 'username',
					username: req.username
				}
			}
		}
	})

	MakotoCookiesAutorization.setToken(headers, cookies)

	console.log(headers, response)

	return new Response(null, {
		status: 200
	})
}
