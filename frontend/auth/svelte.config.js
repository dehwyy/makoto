// import adapter from '@sveltejs/adapter-auto'
import adapter_node from '@sveltejs/adapter-node'
import { vitePreprocess } from '@sveltejs/kit/vite'
import path from 'path'

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),
	kit: {
		csrf: {
			checkOrigin: false
		},
		// adapter-auto only supports some environments, see https://kit.svelte.dev/docs/adapter-auto for a list.
		// If your environment is not supported or you settled on a specific environment, switch out the adapter.
		// See https://kit.svelte.dev/docs/adapters for more information about adapters.
		adapter: adapter_node({
			precompress: true
		}),
		alias: {
			$config: path.resolve('../..', 'config', 'dist', 'config'),
			$rpc: path.resolve('../..', 'libs', 'grpc'),
			$houdini: path.resolve('.', '$houdini')
		}
	}
}

export default config
