import { MakotoCookiesAutorization } from '$lib/api/cookies'
import { RpcInterceptors } from '$lib/api/prc-interceptors'
import { AuthClient } from '$lib/api/twirp-client'
import type { RequestHandler } from '@sveltejs/kit'

export const POST: RequestHandler = async ({ cookies }) => {
	const { response, headers } = await AuthClient.signIn(
		{
			authMethod: {
				oneofKind: 'empty',
				empty: {}
			}
		},
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(cookies.get('token'))]
		}
	)

	console.log(response, headers)
	MakotoCookiesAutorization.setToken(headers, cookies)

	return new Response(null, {
		status: 200
	})
}
