import type { RequestHandler } from '@sveltejs/kit'
import { SafeTwirpClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'
import { HandleAuth } from '$lib/api/handle_auth'

export const GET: RequestHandler = async ({ url, cookies }) => {
	const code = url.searchParams.get('code')
	if (!code) return new Response(null, { status: 403 })

	const token = cookies.get('token')

	const { response, headers, status } = await SafeTwirpClient(cookies).Authorization.signIn(
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

	return HandleAuth.Handle({ status })
}
