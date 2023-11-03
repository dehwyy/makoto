import type { RequestHandler } from '@sveltejs/kit'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { CreateItem } from '$lib/api/fetches'
import { SafeTwirpClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const { key, value, extra, tags, userId } = await tp.Get(request, CreateItem)

	// TODO:
	const { response } = await SafeTwirpClient(cookies).Hashmap.createItem(
		{
			userId: '',
			key,
			value,
			extra,
			tags
		},
		{
			interceptors: [RpcInterceptors.WithToken(cookies)]
		}
	)

	return new Response(JSON.stringify({ itemId: response.id }), { status: 200 })
}
