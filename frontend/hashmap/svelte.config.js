import adapter_node from '@sveltejs/adapter-node'
import { vitePreprocess } from '@sveltejs/kit/vite'
import preprocess from 'svelte-preprocess'

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: [vitePreprocess()],
	onwarn: (warning, handler) => {
		if (warning.code.startsWith('a11y-')) return
		handler(warning)
	},
	kit: {
		csrf: {
			checkOrigin: false
		},
		adapter: adapter_node({
			precompress: true
		})
	}
}

export default config
