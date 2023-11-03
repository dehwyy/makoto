import type { RequestHandler } from '@sveltejs/kit'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { RemoveItem } from '$lib/api/fetches'
import { SafeTwirpClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const { itemId } = await tp.Get(request, RemoveItem)

	// TODO: Fix
	const { response, status } = await SafeTwirpClient(cookies).Hashmap.removeItem(
		{
			itemId,
			userId: ''
		},
		{
			interceptors: [RpcInterceptors.WithToken(cookies)]
		}
	)

	return new Response(null, { status: 200 })
}
