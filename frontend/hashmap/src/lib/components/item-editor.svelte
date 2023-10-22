<script lang="ts">
	import { fade, scale } from 'svelte/transition'
	import { Button, Modal } from 'makoto-ui-svelte'
	import Input from './input.svelte'
	import { TagsStore } from '$lib/stores/tags-store'
	// props
	export let finalButtonText = ''
	export let isEdit = false
	export let defaultValues = {
		key: '',
		value: '',
		extra: '',
		tags: [] as string[]
	}
	export let isCreateItemMode = false

	export let onFinalButtonClick = (payload: {
		key: string
		value: string
		extra: string
		tags: string[]
	}) => {}

	// form_fields
	let key_saved = defaultValues.key
	let key = defaultValues.key

	let value_saved = defaultValues.value
	let value = defaultValues.value

	let extra_saved = defaultValues.extra
	let extra = defaultValues.extra

	// tags
	let time_tag = ''

	let tags_saved = defaultValues.tags
	let tags = defaultValues.tags // temporary stor

	//
	let isInAddingTagsMode = false
	let isTransitionTags = false
	let isInputFocus = false
	let tagsBlock: HTMLElement
	let tagsBlockHeight = 0

	const makeTranstionTags = () => {
		isTransitionTags = true
		setTimeout(() => {
			isTransitionTags = false
		}, 200)
	}

	const SaveTag = () => {
		makeTranstionTags()
		isInAddingTagsMode = false
		setTimeout(() => {
			tags = Array.from(new Set<string>(tags).add(time_tag))
			time_tag = ''
		}, 150)

		tagsBlockHeight = tagsBlock?.clientHeight
	}

	const RemoveTag = (tag: string) => {
		makeTranstionTags()
		setTimeout(() => {
			tags = tags.filter(t => t !== tag)
		}, 150)

		tagsBlockHeight = tagsBlock?.clientHeight
	}

	const SelectAutosuggestedTag = (tag: string) => {
		time_tag = tag
		SaveTag()
	}

	const CloseModal = () => {
		isEdit = false
		isInAddingTagsMode = false
		time_tag = ''

		if (isCreateItemMode) {
			key = ''
			value = ''
			extra = ''
			tags = []

			return
		}
		// restore values from saved
		key = key_saved
		value = value_saved
		extra = extra_saved
		tags = tags_saved
	}

	const FinalButtonClick = () => {
		onFinalButtonClick({ key, value, extra, tags })
		// reset
		isInAddingTagsMode = false
		isEdit = false
		time_tag = ''

		if (isCreateItemMode) {
			key = ''
			value = ''
			extra = ''
			tags = []

			return
		}

		// set *_saved values
		key_saved = key
		value_saved = value
		extra_saved = extra
		tags_saved = tags
	}

	let autocompleteTags: string[]

	$: {
		/**
		 * 1st priopity are tags which start with written query (time_tag)
		 * 2nd priority are tags which not start BUT contain query from any position
		 * 3rd priority are tags which contain query from any position
		 */
		const exac_match: string[] = []
		const contains: string[] = []
		const someway_match: string[] = []

		const query = new RegExp(
			time_tag
				.split('')
				.map(w => w + '[a-zа-я\\s]*')
				.join(''),
			'ig'
		)

		$TagsStore.forEach(tag => {
			if (tags.includes(tag.text)) {
				return
			}

			if (tag.text.startsWith(time_tag)) {
				exac_match.push(tag.text)
			} else if (tag.text.includes(time_tag)) {
				contains.push(tag.text)
			} else if (tag.text.match(query)) {
				someway_match.push(tag.text)
			}
		})

		const result = [...exac_match, ...contains, ...someway_match]
		autocompleteTags = result.length > 3 ? result.slice(0, 3) : result
	}
</script>

<svelte:body on:click={() => (isInputFocus = false)} />

{#if isEdit}
	<article
		transition:fade={{ duration: 150, delay: 0 }}
		class="fixed top-0 left-0 right-0 bottom-0 z-50 font-[600]">
		<!-- Modal itself -->
		<Modal base_width={426} isOpen={isEdit} close={CloseModal}>
			<div class="p-10 flex flex-col gap-y-7 w-full font-Content">
				<!-- Heading -->
				<h2 class="text-2xl text-center font-Content text-white">
					New <span class="underline">Word</span>
				</h2>

				<!-- Form -->
				<Input bind:value={key} placeholder="key" />
				<Input bind:value placeholder="value" />
				<Input bind:value={extra} placeholder="extra" />

				<!-- Tags -->
				<div
					bind:this={tagsBlock}
					style={`${isTransitionTags ? `max-height:${tagsBlockHeight}px` : 'max-height:500px'}`}
					class={`${
						isTransitionTags ? 'opacity-[1%] invisible' : 'opacity-100 visible'
					} flex flex-wrap gap-x-2 gap-y-2 transition-all duration-150`}>
					<!-- Already created tags -->
					{#each tags as tag}
						<div
							class="flex-auto py-2 px-7 rounded-full bg-base-200 text-center font-Content relative pr-9">
							{tag}
							<button
								on:click={() => RemoveTag(tag)}
								class="hover:text-red-500 transition-all duration-300 cursor-pointer bg-base-100 bg-opacity-80 text-gray-400 rounded-full absolute right-[5px] top-[7.5px] h-[25px] w-[25px] text-sm font-[700] font-Jua px-1 grid place-items-center"
								><span class="pt-[2px]">X</span>
							</button>
						</div>
					{/each}

					<!-- Button to enter `createTagMode` -->
					<button
						on:click={() => (isInAddingTagsMode = !isInAddingTagsMode)}
						class={`${
							isInAddingTagsMode ? 'hover:text-red-400' : 'hover:text-green-400'
						} relative select-none cursor-pointer transition-all duration-300 font-[700] text-lg grid place-items-center w-[40px] h-[40px] bg-base-200 rounded-full`}>
						{#if isInAddingTagsMode}
							<span transition:fade={{ duration: 150, delay: 0 }} class="absolute">x</span>
						{:else}
							<span transition:fade={{ duration: 150, delay: 0 }} class="pt-0.5 absolute">+</span>
						{/if}
					</button>

					<!-- `createTagMode` Input and SaveButton -->
					{#if isInAddingTagsMode}
						<div
							transition:fade={{ duration: 200, delay: 0 }}
							class="absolute z-30 right-10 left-10 bottom-32 text-md font-[600]">
							<div class="flex gap-x-5">
								<div
									aria-hidden="true"
									class="w-full"
									on:click={e => {
										e.stopPropagation()
										isInputFocus = true
									}}>
									<Input bind:value={time_tag} placeholder="tag" />
								</div>
								<div
									on:click={SaveTag}
									aria-hidden="true"
									class="cursor-pointer flex items-end rounded-full">
									<span class="hover:text-green-400 transition-all duration-300 font-Jua text-lg"
										>Save</span>
								</div>
							</div>
							<div
								class={`${
									autocompleteTags.length && isInputFocus
										? 'max-h-[150px]'
										: 'max-h-[0px] opacity-0 invisible'
								} overflow-hidden transition-max-height absolute top-[4.7rem] left-0 right-0 bg-base-300 py-1 border-solid border border-gray-300`}>
								{#each autocompleteTags as tag, i}
									<button
										on:click={() => SelectAutosuggestedTag(tag)}
										class="px-4 py-2 font-[600] font-ContentT cursor-pointer hover:bg-[#191919] w-full text-start"
										>{tag}</button>
									{#if i != 2}
										<div class="h-[1px] w-full bg-gray-500 rounded-full" />
									{/if}
								{/each}
							</div>
						</div>
					{/if}
				</div>

				<!-- Button to create word (locally and server) -->
				<div class="w-full mt-28 font-ContentT text-xl">
					<Button onClick={FinalButtonClick}>{finalButtonText}</Button>
				</div>
			</div>
		</Modal>
	</article>
{/if}
<slot />

<style>
	.transition-max-height {
		transition: max-height 0.5s;
	}
</style>
