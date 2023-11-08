import type { RequestHandler } from '@sveltejs/kit'
import { SafeTwirpClient } from '@makoto/grpc/clients'
import { SignInFetch } from '$lib/api/fetches'
import { TypedFetch as tp } from '@makoto/lib/typed-fetch'
import { HandleAuth } from '$lib/api/handle_auth'

export const POST: RequestHandler = async ({ cookies, request }) => {
	const req = await tp.Get(request, SignInFetch)

	const { status } = await SafeTwirpClient(cookies).Authorization.signIn({
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

	return HandleAuth.Handle({ status, cookies })
}
