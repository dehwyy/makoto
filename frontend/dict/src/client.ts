import { HoudiniClient } from '$houdini'

export default new HoudiniClient({
	url: 'http://localhost:4000/query',
	fetchParams({ session }) {
		return {
			headers: {
				// @ts-ignore
				Authorization: `Bearer ${session?.token}`,
				// @ts-ignore
				RefreshToken: `Bearer ${session?.refresh_token}`
			}
		}
	}
})
