import type { LayoutServerLoad } from './$types'
import { SafeTwirpClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const load: LayoutServerLoad = async ({ cookies, params }) => {
	/**
	 * recently had a bug which was caused by the `DataRace`
	 * it was appearing when `AuthToken` (stored in cookies) was expired and have to be renewed
	 * As 2 request with expired token were send, the `DataRace` appeared in Auth microservice
	 * To prevent such kind of behavior, we use sequential request instead of multiple in the same time
	 * To control request sequence, they should be in the same location
	 * (f.e only in +layout.server.ts, not in +page.server.ts + +layout.server.ts as data from +layout.server.ts is accessible in +page.svelte)
	 *
	 * If request doesn't use `AuthToken` そのルールを守る必要ない (don't have to control request sequence)
	 */

	const token = cookies.get('token')

	const { response: tags_response } = await SafeTwirpClient(cookies).Hashmap.getTags(
		{ userId: '' },
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(token)]
		}
	)

	const { response: items_response } = await SafeTwirpClient(cookies).Hashmap.getItems(
		{
			userId: '',
			part: 0,
			partSize: 50,
			query: '',
			tags: []
		},
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(token)]
		}
	)

	const { response: signin_response } = await SafeTwirpClient(cookies).Authorization.signIn(
		{
			authMethod: {
				oneofKind: 'token',
				token: ''
			}
		},
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(token)]
		}
	)
	return {
		userId: signin_response.userId ?? '',
		username: signin_response.username ?? '',
		tags: structuredClone(tags_response.tags) ?? [],
		items: structuredClone(items_response.items) ?? []
	}
}
