import { redirect, type RequestHandler } from '@sveltejs/kit'
import { SafeAuthClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const GET: RequestHandler = async ({ url, cookies }) => {
	const code = url.searchParams.get('code')
	const state = url.searchParams.get('state')

	console.log('code & state', code, state)
	if (!code) return new Response(null, { status: 403 })

	const { response, headers, status } = await SafeAuthClient(cookies).signIn(
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

	throw redirect(301, '/')
}
