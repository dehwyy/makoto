import type { LayoutServerLoad } from './$types'
import { SafeHashmapClient, SafeAuthClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const load: LayoutServerLoad = async ({ cookies }) => {
	const request_get_tags = SafeHashmapClient(cookies).getTags({})
	const request_signin = SafeAuthClient(cookies).signIn(
		{
			authMethod: {
				oneofKind: 'empty',
				empty: {}
			}
		},
		{
			interceptors: [RpcInterceptors.WithToken(cookies)]
		}
	)

	const [{ response: tags_response }, { response: signin_response }] = await Promise.all([
		request_get_tags,
		request_signin
	])

	return {
		userId: signin_response.userId || '',
		username: signin_response.username || '',
		tags: structuredClone(tags_response.tags) || []
	}
}
