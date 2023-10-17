import type { RequestHandler } from '@sveltejs/kit'
import { SignUpFetch } from '$lib/api/fetches'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { AuthClient } from '@makoto/grpc/clients'
import { MakotoCookiesAutorization } from '@makoto/lib/cookies'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const req = await tp.Get(request, SignUpFetch)

	const { response, headers } = await AuthClient.signUp(req)

	MakotoCookiesAutorization.setToken(headers, cookies)

	console.log(response)

	return new Response(null, {
		status: 200
	})
}
