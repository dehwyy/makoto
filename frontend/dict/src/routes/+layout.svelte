<script lang="ts">
	import StarsBackground from '$lib/components/stars-background.svelte'

	// fonts and tailwind
	import '../app.css'
	import '@fontsource-variable/comfortaa'
	import '@fontsource/klee-one'
	import '@fontsource/jua'
	import '@fontsource-variable/shantell-sans'
	import Header from '$lib/components/header.svelte'
</script>

<svelte:document style="background-color: red;" />

<StarsBackground>
	<Header />
	<slot />
</StarsBackground>
