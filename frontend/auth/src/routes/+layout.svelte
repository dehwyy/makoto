<script lang="ts">
	import '../app.css'
	import '@fontsource-variable/comfortaa'
	import '@fontsource/klee-one'
	import '@fontsource/jua'
	import '@fontsource-variable/shantell-sans'

	import { StarsBackground } from 'makoto-ui-svelte'

	import { fade } from 'svelte/transition'
	import { onMount } from 'svelte'
	import { SignInByToken } from '$lib/api/typed-fetch'

	let isMounted = false
	onMount(() => {
		isMounted = true
		SignInByToken()
	})
</script>

<StarsBackground>
	{#if isMounted}
		<main class="h-full w-full grid place-items-center overflow-hidden">
			<div
				data-testid="form"
				class="w-5/6 md:w-2/3 lg:w-1/3 max-w-[400px] mx-auto"
				transition:fade={{ duration: 1500, delay: 100 }}>
				<slot />
			</div>
		</main>
	{/if}
</StarsBackground>
