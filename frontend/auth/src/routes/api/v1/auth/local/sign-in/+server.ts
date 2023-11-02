import type { RequestHandler } from '@sveltejs/kit'
import { SafeTwirpClient } from '@makoto/grpc/clients'
import { SignInFetch } from '$lib/api/fetches'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const req = await tp.Get(request, SignInFetch)

	const { headers, response } = await SafeTwirpClient(cookies).Authorization.signIn({
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

	console.log(response)

	return new Response(null, {
		status: 200
	})
}
