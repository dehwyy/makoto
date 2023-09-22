import { graphql } from '$houdini'
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
	cookies.set('auth-token', tokens.access_token, { httpOnly: true, path: '/' })
	cookies.set('refresh-token', tokens.refresh_token, { httpOnly: true, path: '/' })

	return new Response(JSON.stringify(userId), { status: 200 })
}
