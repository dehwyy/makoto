import type { RequestHandler } from '@sveltejs/kit'
import { SignUpFetch } from '$lib/api/fetches'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { SafeAuthClient } from '@makoto/grpc/clients'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const req = await tp.Get(request, SignUpFetch)

	const { response, headers } = await SafeAuthClient(cookies).signUp(req)

	return new Response(null, {
		status: 200
	})
}
