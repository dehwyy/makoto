<script lang="ts">
	import Input from '$lib/components/input.svelte'
	import { FilterStore } from '../../../routes/store'

	let query = ''
	$: FilterStore.set(query)
</script>

<Input bind:value={query} placeholder="Enter word to find..." />
