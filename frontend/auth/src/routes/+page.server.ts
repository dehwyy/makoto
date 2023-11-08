import { redirect } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ cookies, url }) => {
	const redirect_url = url.searchParams.get('redirect')

	if (redirect_url && redirect_url != '') {
		// TODO:
		throw redirect(307, `http://localhost:3000${redirect_url}`)
	}
}
