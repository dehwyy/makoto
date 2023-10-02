<script lang="ts">
	import StarsBackground from '$lib/components/stars-background.svelte'

	// fonts and tailwind
	import '../app.css'
	import '@fontsource-variable/comfortaa'
	import '@fontsource/klee-one'
	import '@fontsource/jua'
	import '@fontsource-variable/shantell-sans'
	import Header from '$lib/components/header.svelte'
	import type { LayoutData } from './$types'

	export let data: LayoutData

	$: ({ getUsetById } = data)
</script>

<StarsBackground>
	<Header username={$getUsetById.data?.getUserById.username} />
	<slot />
</StarsBackground>
