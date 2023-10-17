import type { PageServerLoad } from './$types'
import { SafeHashmapClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const load: PageServerLoad = async event => {
	const { response, status } = await SafeHashmapClient.getItems(
		{
			userId: ''
		},
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(event.cookies.get('token'))]
		}
	)

	if (status.code.startsWith('4')) {
		console.log(status)
	}

	return {
		Items: structuredClone(response.items) || []
	}
}
