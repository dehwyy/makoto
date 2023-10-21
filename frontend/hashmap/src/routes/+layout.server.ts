import type { LayoutServerLoad } from './$types'
import { SafeHashmapClient, SafeAuthClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const load: LayoutServerLoad = async ({ cookies }) => {
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

	// send this request in parallel as it doesn't use `AuthToken`
	const request_get_tags = SafeHashmapClient(cookies).getTags({})

	// but not this one as it use `AuthToken` so we have to wait until response would be get
	const { response: items_response } = await SafeHashmapClient(cookies).getItems(
		{
			userId: ''
		},
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(cookies.get('token'))]
		}
	)

	// only after this sending request. As this is the latest request we don't have to wait it => instead we'll wait for all parallel at the end
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
		userId: signin_response.userId ?? '',
		username: signin_response.username ?? '',
		tags: structuredClone(tags_response.tags) ?? [],
		items: structuredClone(items_response.items) ?? []
	}
}
