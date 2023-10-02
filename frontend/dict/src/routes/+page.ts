import { graphql, load_GetWords } from '$houdini'
import { SetTokens } from '$lib/api/set-tokens'
import type { PageLoad } from './$types'

export const _houdini_load = graphql(`
	query GetWords {
		getWords {
			words {
				wordId
				word
				value
				extra
				tags {
					tagId
					text
				}
			}
			tokens {
				access_token
				refresh_token
			}
		}
	}
`)

export const load: PageLoad = async event => {
	const { GetWords } = await load_GetWords({
		event,
		then: res => {
			if (!res) return
			const { access_token, refresh_token } = res.getWords.tokens!

			fetch('/api/set-tokens', {
				method: 'POST',
				body: JSON.stringify({
					access_token: access_token,
					refresh_token: refresh_token
				})
			})
		}
	})
	return {
		GetWords
	}
}
