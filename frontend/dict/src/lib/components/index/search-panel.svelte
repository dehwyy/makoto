<script lang="ts">
	import { Input } from 'makoto-ui-svelte'
	import { FilterStore } from '$lib/stores/items-store'
	import Placeholder from '$lib/components/placeholder.svelte'

	const placeholder = 'Enter word to find...'
	let query = ''
	$: FilterStore.set(query)
</script>

<Input bind:value={query} {placeholder}>
	<Placeholder {placeholder} />
</Input>
