import { graphql, load_GetUserById } from '$houdini'
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

export const load: LayoutServerLoad = async event => {
	const user_id = event?.cookies?.get(USER_ID)

	const { GetUserById } = await load_GetUserById({
		event,
		variables: {
			userId: user_id || ''
		},
		then: res => {
			if (!res) return
			const { access_token, refresh_token } = res.getUserById.auth.tokens!

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
		getUsetById: GetUserById
	}
}
