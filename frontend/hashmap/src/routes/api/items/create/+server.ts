import type { RequestHandler } from '@sveltejs/kit'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { CreateItem } from '$lib/api/fetches'
import { SafeHashmapClient } from '@makoto/grpc/clients'
import { RpcInterceptors } from '@makoto/grpc'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const { key, value, extra, tags } = await tp.Get(request, CreateItem)

	const { response, status } = await SafeHashmapClient.createItem(
		{
			key,
			value,
			extra,
			tags
		},
		{
			interceptors: [RpcInterceptors.WithToken(cookies)]
		}
	)
	console.log('create', response, status)

	return new Response(null, { status: 200 })
}
