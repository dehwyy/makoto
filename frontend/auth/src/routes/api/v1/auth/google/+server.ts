import { redirect, type RequestHandler } from '@sveltejs/kit'
import { google as GoogleLib } from 'googleapis'

// TODO: fix config

// export const GET: RequestHandler = async () => {
// 	const google = CONFIG.oauth2.google
// 	const oauth2Client = new GoogleLib.auth.OAuth2(
// 		google.client_id,
// 		google.secret,
// 		google.redirect_url
// 	)

// 	const authorizationUrl = oauth2Client.generateAuthUrl({
// 		access_type: 'offline',
// 		scope: [
// 			'https://www.googleapis.com/auth/userinfo.email',
// 			'https://www.googleapis.com/auth/userinfo.profile'
// 		],
// 		include_granted_scopes: true
// 	})

// 	throw redirect(301, authorizationUrl)
// }
