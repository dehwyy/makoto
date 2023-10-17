import { AuthClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'
import { MakotoCookiesAutorization } from '@makoto/lib/cookies'
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
