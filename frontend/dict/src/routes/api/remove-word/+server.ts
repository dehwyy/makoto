import { graphql } from '$houdini'
import { SetTokensCookies } from '$lib/api/set-tokens'
import type { RequestHandler } from './$types'

export const POST: RequestHandler = async event => {
	const { request, cookies } = event
	const { wordId } = await request.json()

	const mutation = graphql(`
		mutation DeleteWord($wordId: ID!) {
			RemoveWord(wordId: $wordId) {
				access_token
				refresh_token
			}
		}
	`)

	const res = await mutation.mutate({ wordId: wordId }, { event })

	console.log('ERRORS', res.errors)
	if (res.errors || !res.data?.RemoveWord)
		return new Response(JSON.stringify(res.errors), { status: 400 })

	const { access_token, refresh_token } = res.data.RemoveWord!
	SetTokensCookies({ access: access_token, refresh: refresh_token, cookies })

	return new Response(null, { status: 200 })
}
