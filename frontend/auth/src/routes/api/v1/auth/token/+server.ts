import { SafeAuthClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'
import type { RequestHandler } from '@sveltejs/kit'

export const POST: RequestHandler = async ({ cookies }) => {
	const { response, headers } = await SafeAuthClient(cookies).signIn(
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

	return new Response(null, {
		status: 200
	})
}
