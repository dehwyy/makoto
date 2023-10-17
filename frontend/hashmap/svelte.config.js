import adapter from '@sveltejs/adapter-auto'
import { vitePreprocess } from '@sveltejs/kit/vite'
import preprocess from 'svelte-preprocess'

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: [
		vitePreprocess({
			style: {
				css: {
					postcss: true
				}
			}
		}),
		preprocess({
			scss: true,
			postcss: true
		})
	],
	onwarn: (warning, handler) => {
		console.log('HELLO>')
		if (warning.code.startsWith('a11y-')) return
		handler(warning)
	},
	kit: {
		adapter: adapter()
	}
}

export default config
