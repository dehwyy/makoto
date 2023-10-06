import { HoudiniClient } from '$houdini'

const url = process.env['PROD'] ? `http://distributor:4000/query` : `http://localhost:4000/query`

export default new HoudiniClient({
	url
})
