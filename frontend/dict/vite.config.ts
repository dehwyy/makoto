import { sveltekit } from '@sveltejs/kit/vite'
import houdini from 'houdini/vite'
import { defineConfig } from 'vitest/config'

export default defineConfig({
	plugins: [sveltekit(), houdini()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
})
