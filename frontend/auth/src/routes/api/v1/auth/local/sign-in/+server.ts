import type { RequestHandler } from '@sveltejs/kit'
import { AuthClient } from '@makoto/grpc/$lib/clients'
import { RpcInterceptors } from '@makoto/grpc/$lib/interceptors'
import { SignInFetch } from '$lib/api/fetches'
import { MakotoCookiesAutorization } from '@makoto/lib/cookies'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const req = await tp.Get(request, SignInFetch)

	const { headers, response } = await AuthClient.signIn({
		authMethod: {
			oneofKind: 'credentials',
			credentials: {
				password: req.password,
				uniqueIdentifier: {
					oneofKind: 'username',
					username: req.username
				}
			}
		}
	})

	MakotoCookiesAutorization.setToken(headers, cookies)

	console.log(headers, response)

	return new Response(null, {
		status: 200
	})
}
