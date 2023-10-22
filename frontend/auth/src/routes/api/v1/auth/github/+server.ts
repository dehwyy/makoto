import { redirect, type RequestHandler } from '@sveltejs/kit'
import { config } from '@makoto/config'

export const GET: RequestHandler = async () => {
	const createAuthorizationUrl = (
		client_id: string,
		redirect_url: string,
		state: string,
		scoped: string[] = []
	) => {
		// prettier-ignore
		return `https://github.com/login/oauth/authorize?client_id=${client_id}&redirect_uri=${redirect_url}&scope=${scoped.join(',')}&state=${state}`
	}
	const authorizationUrl = createAuthorizationUrl(
		config.GITHUB_CLIENT_ID,
		config.GITHUB_REDIRECT_URL,
		config.GITHUB_STATE
	)

	throw redirect(301, authorizationUrl)
}
