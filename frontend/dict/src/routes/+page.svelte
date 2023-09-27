<script lang="ts">
	import { fade } from 'svelte/transition'

	import AddWordPanel from '$lib/components/index/add-word-panel.svelte'
	import Word from '$lib/components/word.svelte'
	import SearchPanel from '$lib/components/index/search-panel.svelte'

	import { FilteredItems, RemoveItemById, ItemsStore } from './store'
	import { onMount } from 'svelte'

	let isMounted = false

	onMount(() => {
		isMounted = true
		ItemsStore.set([{ id: 1, word: '星々', translate: 'stars', extra: 'some extra' }])
	})
</script>

{#if isMounted}
	<main class="py-20 w-[70%] mx-auto flex flex-col gap-y-24 items-center">
		<!-- search bar -->
		<section transition:fade={{ duration: 300, delay: 0 }} class="w-1/2 mx-auto">
			<SearchPanel />
		</section>

		<!-- add new word button and modal-->
		<section
			transition:fade={{ duration: 300, delay: 100 }}
			class="self-start w-[80%] mx-auto mb-[-3rem]">
			<AddWordPanel />
		</section>

		<!-- words -->
		<section
			transition:fade={{ duration: 300, delay: 200 }}
			class="w-[80%] flex flex-col gap-y-10 items-center">
			<div class="w-full flex flex-col gap-y-5">
				{#each $FilteredItems as item}
					<Word
						onCloseButtonClick={() => RemoveItemById(item.id)}
						word={item.word}
						translate={item.translate}
						extra={item.extra} />
				{/each}
			</div>
		</section>
	</main>
{/if}
