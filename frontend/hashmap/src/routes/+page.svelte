<script lang="ts">
	import { fade } from 'svelte/transition'

	import AddWordPanel from '$lib/components/index/add-word-panel.svelte'
	import Word from '$lib/components/word.svelte'
	import SelectTags from '../lib/components/tags.svelte'
	import SearchPanel from '$lib/components/index/search-panel.svelte'

	import { FilteredItems, Items } from '$lib/stores/items-store'
	import { RemoveItem } from '$lib/api/fetches'
	import { onMount } from 'svelte'
	import type { PageData } from './$types'

	let isMounted = false
	onMount(() => {
		isMounted = true
	})

	const removeWord = async (itemId: number) => {
		Items.RemoveById(itemId)

		RemoveItem({ itemId })
	}

	export let data: PageData
	Items.Set(data.Items || [])
</script>

{#if isMounted}
	<main class="page_wrapper">
		<!-- search bar -->
		<div transition:fade={{ duration: 300, delay: 0 }} class="page_search">
			<SearchPanel />
			<section class="flex gap-x-5">
				<SelectTags />
			</section>
		</div>

		<!-- add new word button and modal-->
		<section transition:fade={{ duration: 300, delay: 100 }} class="page_add-word">
			<AddWordPanel />
		</section>

		<!-- words -->
		<section transition:fade={{ duration: 300, delay: 200 }} class="page_word_wrapper">
			<div class="page_word_item">
				{#each $FilteredItems as item}
					<Word
						onCloseButtonClick={() => removeWord(item.id)}
						word={item.key}
						translate={item.value}
						extra={item.extra} />
				{/each}
			</div>
		</section>
	</main>
{/if}

<style lang="scss">
	.page {
		&_wrapper {
			@apply py-20 w-[70%] mx-auto flex flex-col gap-y-24 items-center;
		}
		&_search {
			@apply w-full lg:w-1/2 mx-auto font-Content flex flex-col gap-y-3;
		}
		&_add-word {
			@apply w-[80%] mx-auto mb-[-3rem];
		}
		&_word {
			&_wrapper {
				@apply flex flex-col gap-y-10 items-center;
			}
			&_item {
				@apply w-full flex flex-col gap-y-5;
			}
		}
	}
</style>
