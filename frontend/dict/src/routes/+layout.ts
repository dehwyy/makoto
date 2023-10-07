import { graphql, load_GetTags, load_GetUserById } from '$houdini'
import { SetTokensFromResponse } from '$lib/api/set-tokens-request'
import { USER_ID } from '$lib/const'
import type { LayoutServerLoad } from './$types'

graphql(`
	query GetUserById($userId: ID!) {
		getUserById(input: { userId: $userId }) {
			auth {
				tokens {
					access_token
					refresh_token
				}
			}
			username
		}
	}
`)

graphql(`
	query GetTags {
		getTags {
			tags {
				tagId
				text
			}
			tokens {
				refresh_token
				access_token
			}
		}
	}
`)

export const load: LayoutServerLoad = async event => {
	const user_id = event?.cookies?.get(USER_ID)

	const tags_response = load_GetTags({
		event,
		then: res => SetTokensFromResponse(res?.getTags.tokens!)
	})

	const user_response = load_GetUserById({
		event,
		variables: {
			userId: user_id || ''
		},
		then: res => SetTokensFromResponse(res?.getUserById.auth.tokens)
	})

	const [{ GetTags }, { GetUserById }] = await Promise.all([tags_response, user_response])

	return {
		getUsetById: GetUserById,
		getTags: GetTags
	}
}
