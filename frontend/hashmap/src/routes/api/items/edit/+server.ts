import type { RequestHandler } from '@sveltejs/kit'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { EditItem } from '$lib/api/fetches'
import { SafeTwirpClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const { itemId, key, value, extra, tags } = await tp.Get(request, EditItem)

	// TODO: fix
	const { response, status } = await SafeTwirpClient(cookies).Hashmap.editItem(
		{
			itemId,
			key,
			value,
			extra,
			tags,
			userId: ''
		},
		{
			interceptors: [RpcInterceptors.WithToken(cookies)]
		}
	)

	return new Response(null, { status: 200 })
}
