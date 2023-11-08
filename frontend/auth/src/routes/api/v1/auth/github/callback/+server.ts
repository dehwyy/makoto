import { SafeTwirpClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'
import type { RequestHandler } from '@sveltejs/kit'
import { HandleAuth } from '$lib/api/handle_auth'

export const GET: RequestHandler = async ({ url, cookies }) => {
	const code = url.searchParams.get('code')
	const state = url.searchParams.get('state')

	console.log('code & state', code, state)
	if (!code) return new Response(null, { status: 403 })

	const { response, headers, status } = await SafeTwirpClient(cookies).Authorization.signIn(
		{
			authMethod: {
				oneofKind: 'oauth2',
				oauth2: {
					provider: 'github',
					code
				}
			}
		},
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(cookies.get('token'))]
		}
	)

	return HandleAuth.Handle({ status })
}
