import type { RequestHandler } from '@sveltejs/kit'
import { SignUpFetch } from '$lib/api/fetches'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { SafeTwirpClient } from '@makoto/grpc/clients'
import { HandleAuth } from '$lib/api/handle_auth'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const req = await tp.Get(request, SignUpFetch)
	const { status } = await SafeTwirpClient(cookies).Authorization.signUp(req)

	return HandleAuth.Handle({ status, cookies })
}
