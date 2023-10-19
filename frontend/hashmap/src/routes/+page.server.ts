import type { PageServerLoad } from './$types'
import { SafeHashmapClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const load: PageServerLoad = async ({ cookies }) => {
	const { response, status } = await SafeHashmapClient(cookies).getItems(
		{
			userId: ''
		},
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(cookies.get('token'))]
		}
	)

	if (status.code.startsWith('4')) {
		console.log(status)
	}

	return {
		Items: structuredClone(response.items) || []
	}
}
