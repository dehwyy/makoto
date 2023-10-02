import { graphql } from '$houdini'
import { SetTokensCookies } from '$lib/api/set-tokens'
import type { RequestHandler } from './$types'

export const POST: RequestHandler = async event => {
	const { request, cookies } = event
	const { word, value, extra, tags } = await request.json()

	const mutation = graphql(`
		mutation CreateWord($word: String!, $value: String!, $extra: String!, $tags: [String!]!) {
			createWord(word: { word: $word, value: $value, extra: $extra, tags: $tags }) {
				access_token
				refresh_token
			}
		}
	`)

	const res = await mutation.mutate({ word, value, extra, tags }, { event })

	console.log('ERRORS', res.errors)
	if (res.errors || !res.data?.createWord)
		return new Response(JSON.stringify(res.errors), { status: 400 })

	const { access_token, refresh_token } = res.data.createWord!
	SetTokensCookies({ access: access_token, refresh: refresh_token, cookies })

	return new Response(null, { status: 200 })
}
