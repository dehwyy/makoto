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
	import { Username, UserId } from '$lib/stores/user-store'
	import { Tags } from '$lib/stores/tags-store'

	// server types
	import type { LayoutData } from './$types'
	export let data: LayoutData

	Tags.Set(data.tags.map(t => ({ ...t, usages: 1 })))
	Username.set(data.username)
	UserId.set(data.userId)
</script>

<StarsBackground>
	<Header />
	<slot />
</StarsBackground>
