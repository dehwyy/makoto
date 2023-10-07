<script lang="ts">
	// fonts and tailwind
	import '../app.css'
	import '@fontsource-variable/comfortaa'
	import '@fontsource/klee-one'
	import '@fontsource/jua'
	import '@fontsource-variable/shantell-sans'
	// components
	import Header from '$lib/components/header.svelte'
	import { StarsBackground } from 'makoto-ui-svelte'
	// stores
	import { AuthUsername } from '$lib/stores/user-store'

	// fetch current user via auth token if exists
	import type { LayoutData } from './$types'
	import { SetTags } from '$lib/stores/items-store'
	export let data: LayoutData

	$: ({ getUsetById, getTags } = data)
	$: AuthUsername.set($getUsetById.data?.getUserById.username ?? '')
	$: SetTags($getTags.data?.getTags.tags! ?? [])
</script>

<StarsBackground>
	<Header />
	<slot />
</StarsBackground>
