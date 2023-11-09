<script lang="ts">
	import { Input } from 'makoto-ui-svelte'
	import { ItemsFilterQueryStore } from '$lib/stores/items-store'
	import Placeholder from '$lib/components/placeholder.svelte'

	const debounce_t = 250
	const placeholder = 'Enter word to find...'
	let query = ''
	let timeout = 0

	$: {
		clearTimeout(timeout)
		timeout = window.setTimeout(() => {
			ItemsFilterQueryStore.set(query)
		}, debounce_t)
	}
</script>

<Input bind:value={query} {placeholder}>
	<Placeholder {placeholder} />
</Input>
