import { graphql } from '$houdini'
import { setAuthCookies } from '$lib/api/set-cookies'
import type { RequestHandler } from './$types'

export const POST: RequestHandler = async event => {
	const { request, cookies } = event
	const { username, password } = await request.json()

	const mutation = graphql(`
		mutation SignIn($username: String!, $password: String!) {
			signIn(input: { username: $username, password: $password }) {
				userId
				tokens {
					access_token
					refresh_token
				}
			}
		}
	`)

	const res = await mutation.mutate({ username, password }, { event })
	if (!res.data) return new Response(JSON.stringify(res.errors), { status: 400 })

	const { tokens, userId } = res.data.signIn

	setAuthCookies({
		access_token: tokens.access_token,
		refresh_token: tokens.refresh_token,
		user_id: userId,
		cookies
	})

	return new Response(JSON.stringify(userId), { status: 200 })
}
