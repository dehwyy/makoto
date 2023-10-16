import { redirect, type RequestHandler } from '@sveltejs/kit'
import { AuthClient } from '@makoto/grpc/$lib/clients'
import { RpcInterceptors } from '@makoto/grpc/$lib/interceptors'
import { MakotoCookiesAutorization } from '@makoto/lib/cookies'

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
