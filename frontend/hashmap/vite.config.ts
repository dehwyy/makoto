import { PORTS } from '@makoto/config'
import { sveltekit } from '@sveltejs/kit/vite'
import { defineConfig } from 'vite'

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		port: PORTS.HASHMAP
	}
})
