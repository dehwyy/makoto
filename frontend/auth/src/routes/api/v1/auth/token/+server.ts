import { SafeTwirpClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'
import type { RequestHandler } from '@sveltejs/kit'
import { HandleAuth } from '$lib/api/handle_auth'

export const POST: RequestHandler = async ({ cookies }) => {
	const { status } = await SafeTwirpClient(cookies).Authorization.signIn(
		{
			authMethod: {
				oneofKind: 'token',
				token: ''
			}
		},
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(cookies.get('token'))]
		}
	)

	return HandleAuth.Handle({ status })
}
