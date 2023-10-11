import { redirect, type RequestHandler } from '@sveltejs/kit'
import { AuthClient } from '$lib/api/twirp-client'
import { RpcInterceptors } from '$lib/api/prc-interceptors'
import { RpcPayloads } from '$lib/api/rpc-payload'
import { MakotoCookies } from '$lib/api/cookies'

export const GET: RequestHandler = async event => {
	const code = event.url.searchParams.get('code')
	console.log(event.url.searchParams)
	if (!code) return new Response(null, { status: 403 })

	const token = event.cookies.get('token')

	const { response, headers, status } = await AuthClient.signIn(
		RpcPayloads.SignIn({ code, token, provider: 'google' }),
		{
			interceptors: [RpcInterceptors.AddAuthorizationHeader(token)]
		}
	)

	console.log(code, status, response, headers)

	MakotoCookies.setGlobal(event.cookies, 'token', headers['authorization'] as string)

	throw redirect(301, '/')
}
