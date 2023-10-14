import { redirect, type RequestHandler } from '@sveltejs/kit'
import { AuthClient } from '$lib/api/twirp-client'
import { RpcInterceptors } from '$lib/api/prc-interceptors'
import { MakotoCookiesAutorization } from '$lib/api/cookies'

export const GET: RequestHandler = async event => {
	const code = event.url.searchParams.get('code')
	if (!code) return new Response(null, { status: 403 })

	const token = event.cookies.get('token')

	const { response, headers, status } = await AuthClient.signIn(
		// RpcPayloads.SignIn({ code, provider: 'google' }),
		{
			authMethod: {
				oneofKind: 'oauth2',
				oauth2: {
					provider: 'google',
					code
				}
			}
		},
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(token)]
		}
	)

	console.log(code, status, response, headers)

	MakotoCookiesAutorization.setToken(headers, event.cookies)

	throw redirect(301, '/')
}
