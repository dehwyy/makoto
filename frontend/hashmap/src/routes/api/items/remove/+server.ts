import type { RequestHandler } from '@sveltejs/kit'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { RemoveItem } from '$lib/api/fetches'
import { SafeHashmapClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const { itemId } = await tp.Get(request, RemoveItem)

	const { response, status } = await SafeHashmapClient(cookies).removeItem(
		{
			itemId
		},
		{
			interceptors: [RpcInterceptors.WithToken(cookies)]
		}
	)

	return new Response(null, { status: 200 })
}
