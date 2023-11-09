import type { RequestHandler } from '@sveltejs/kit'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { GetItems } from '$lib/api/fetches'
import { SafeTwirpClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'
import { HandleResponse } from '$lib/api/handle_response'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const { userId, part, partSize, query, tags } = await tp.Get(request, GetItems)

	// TODO: fix
	const { response, status } = await SafeTwirpClient(cookies).Hashmap.getItems(
		{
			userId: '',
			part,
			partSize,
			query,
			tags
		},
		{
			interceptors: [RpcInterceptors.WithToken(cookies)]
		}
	)

	const success = HandleResponse.Handle(status.code)
	if (!success) {
		return new Response(null, { status: 500 })
	}

	return new Response(JSON.stringify({ items: response.items }), { statusText: status.code })
}
