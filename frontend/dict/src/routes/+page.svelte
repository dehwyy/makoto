<script lang="ts">
	import { fade } from 'svelte/transition'

	import AddWordPanel from '$lib/components/index/add-word-panel.svelte'
	import Word from '$lib/components/word.svelte'
	import SearchPanel from '$lib/components/index/search-panel.svelte'

	import { FilteredItems, RemoveItemById, SetItems } from './items-store'
	import { getContext, onMount, setContext } from 'svelte'

	let isMounted = false

	onMount(() => {
		isMounted = true
	})

	const removeWord = async (wordId: string) => {
		RemoveItemById(wordId)

		const response = await fetch('/api/remove-word', {
			method: 'POST',
			body: JSON.stringify({
				wordId
			})
		})

		const isError = response.statusText.startsWith('4')
		isError && console.log('Error occured', response.statusText)
	}

	// Loading words
	import type { PageData } from './$houdini'
	import { isAuthed } from './user-store'
	export let data: PageData

	$: ({ GetWords } = data)
	$: SetItems($GetWords.data?.getWords.words || [])

	$: isAuthed.set(!!$GetWords.data?.getWords.tokens?.access_token)
</script>

{#if isMounted}
	<main class="py-20 w-[70%] mx-auto flex flex-col gap-y-24 items-center">
		<!-- search bar -->
		<div transition:fade={{ duration: 300, delay: 0 }} class="w-1/2 mx-auto">
			<SearchPanel />
		</div>

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
						onCloseButtonClick={() => removeWord(item.wordId)}
						word={item.word}
						translate={item.value}
						extra={item.extra} />
				{/each}
			</div>
		</section>
	</main>
{/if}
