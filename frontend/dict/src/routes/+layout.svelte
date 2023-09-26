<script lang="ts">
	import HoshiNoSora from './../lib/components/hoshi-no-sora.svelte'

	// fonts and tailwind
	import '../app.css'
	import '@fontsource-variable/comfortaa'
	import '@fontsource/klee-one'
	import '@fontsource/jua'
	import '@fontsource-variable/shantell-sans'
</script>

<svelte:document style="background-color: red;" />

<HoshiNoSora>
	<slot />
</HoshiNoSora>
