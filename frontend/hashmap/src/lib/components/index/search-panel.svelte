<script lang="ts">
	import { Input } from 'makoto-ui-svelte'
	import { ItemsFilterQueryStore } from '$lib/stores/items-store'
	import Placeholder from '$lib/components/placeholder.svelte'

	const placeholder = 'Enter word to find...'
	let query = ''
	$: ItemsFilterQueryStore.set(query)
</script>

<Input bind:value={query} {placeholder}>
	<Placeholder {placeholder} />
</Input>
