import { redirect, type RequestHandler } from '@sveltejs/kit'
import { google as GoogleLib } from 'googleapis'
import { config } from '@makoto/config'

export const GET: RequestHandler = async () => {
	const oauth2Client = new GoogleLib.auth.OAuth2(
		config.GOOGLE_CLIENT_ID,
		config.GOOGLE_CLIENT_SECRET,
		config.GOOGLE_REDIRECT_URL
	)

	const authorizationUrl = oauth2Client.generateAuthUrl({
		access_type: 'offline',
		scope: [
			'https://www.googleapis.com/auth/userinfo.email',
			'https://www.googleapis.com/auth/userinfo.profile'
		],
		include_granted_scopes: true
	})

	throw redirect(301, authorizationUrl)
}
