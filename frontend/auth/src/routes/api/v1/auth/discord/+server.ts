import { redirect, type RequestHandler } from '@sveltejs/kit'
import { config } from '@makoto/config'

export const GET: RequestHandler = async () => {
	const createAuthorizationUrl = (client_id: string, redirect_url: string, state: string) => {
		return `https://discord.com/oauth2/authorize?response_type=code&client_id=${client_id}&redirect_uri=${redirect_url}&scope=identify%20email`
	}
	const authorizationUrl = createAuthorizationUrl(
		config.DISCORD_CLIENT_ID,
		config.DISCORD_REDIRECT_URL,
		''
	)

	throw redirect(307, authorizationUrl)
}
