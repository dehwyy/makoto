<script lang="ts">
	import { fade } from 'svelte/transition'

	import Item from '$lib/components/item.svelte'
	import SelectTags from '../lib/components/tags.svelte'
	import SearchPanel from '$lib/components/index/search-panel.svelte'

	import { FilteredItems, Items } from '$lib/stores/items-store'
	import { RemoveItem } from '$lib/api/fetches'
	import { onMount } from 'svelte'
	import type { PageData } from './$types'
	import AddItem from '$lib/components/index/add-item.svelte'

	let isMounted = false
	onMount(() => {
		isMounted = true
	})

	const removeItem = async (itemId: number) => {
		Items.RemoveById(itemId)

		RemoveItem({ itemId })
	}

	export let data: PageData
	Items.Set(data.items || [])
</script>

{#if isMounted}
	<main class="py-20 w-[70%] mx-auto flex flex-col gap-y-24 items-center">
		<!-- search bar -->
		<div
			transition:fade={{ duration: 300, delay: 0 }}
			class="w-full lg:w-1/2 mx-auto font-Content flex flex-col gap-y-3">
			<SearchPanel />
			<section class="flex gap-x-5">
				<SelectTags />
			</section>
		</div>

		<!-- add new word button and modal-->
		<section transition:fade={{ duration: 300, delay: 100 }} class="w-[80%] mx-auto mb-[-3rem]">
			<AddItem />
		</section>

		<!-- items -->
		<section
			transition:fade={{ duration: 300, delay: 200 }}
			class="flex flex-col gap-y-10 items-center lg:w-[80%] w-full">
			<div class="w-full flex flex-col gap-y-5">
				{#each $FilteredItems as item}
					<Item
						removeItem={() => removeItem(item.id)}
						tags={item.tags.map(tag => tag.text)}
						item_id={item.id}
						item={item.key}
						value={item.value}
						extra={item.extra} />
				{/each}
			</div>
		</section>
	</main>
{/if}
