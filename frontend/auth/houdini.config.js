const url = process.env['PROD'] ? `http://distributor:4000/query` : `http://localhost:4000/query`

/** @type {import('houdini').ConfigFile} */
const config = {
	watchSchema: {
		url
	},

	plugins: {
		'houdini-svelte': {}
	}
}

export default config
