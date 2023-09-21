import type { Handle } from '@sveltejs/kit'

export const handle: Handle = async ({ event, resolve }) => {
	const authHeader = event.request.headers.get('Authorization')
	const cookieKey = 'auth-token'

	console.log(authHeader, cookieKey)

	// if (authHeader && cookieKey) {
	// 	event.cookies.set(cookieKey, authHeader)
	// }

	const response = await resolve(event)
	// response.headers.set('x-custom-header', 'potato')

	return response
}
