<script lang="ts">
	import { fade, scale } from 'svelte/transition'
	import { Button, Modal } from 'makoto-ui-svelte'
	import Input from './input.svelte'

	// props
	export let finalButtonText = ''
	export let isEdit = false
	export let defaultValues = {
		key: '',
		value: '',
		extra: '',
		tags: [] as string[]
	}
	export let onFinalButtonClick = (payload: {
		key: string
		value: string
		extra: string
		tags: string[]
	}) => {}

	// form_fields
	let key = defaultValues.key
	let value = defaultValues.value
	let extra = defaultValues.extra

	// tags
	let time_tag = ''
	let tags: string[] = defaultValues.tags // temporary stor

	$: isInAddingTagsMode = false

	const SaveTag = () => {
		tags = Array.from(new Set<string>(tags).add(time_tag))
		time_tag = ''

		isInAddingTagsMode = false
	}

	const FinalButtonClick = () => {
		onFinalButtonClick({ key, value, extra, tags })
		// reset
		isEdit = false
		key = ''
		value = ''
		extra = ''
		time_tag = ''
		tags = []
	}
</script>

{#if isEdit}
	<article
		transition:fade={{ duration: 150, delay: 0 }}
		class="fixed top-0 left-0 right-0 bottom-0 z-50">
		<!-- Modal itself -->
		<Modal base_width={400} isOpen={isEdit} close={() => (isEdit = !isEdit)}>
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
				<div class="flex flex-wrap gap-x-2 gap-y-2">
					<!-- Already created tags -->
					{#each tags as tag}
						<div
							transition:scale={{ duration: 500, delay: 0 }}
							class="flex-auto py-2 px-7 rounded-full bg-base-200 text-center font-Content relative pr-9">
							{tag}
							<button
								on:click={() => (tags = tags.filter(t => t !== tag))}
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
							transition:fade={{ duration: 300, delay: 0 }}
							class="absolute z-30 right-10 left-10 bottom-28 flex gap-x-5 text-md font-[600]">
							<Input bind:value={time_tag} placeholder="tag" />
							<div
								on:click={SaveTag}
								aria-hidden="true"
								class="cursor-pointer flex items-end rounded-full">
								<span class="hover:text-green-400 transition-all duration-300 font-Jua text-lg"
									>Save</span>
							</div>
						</div>
					{/if}
				</div>

				<!-- Button to create word (locally and server) -->
				<div class="w-full mt-16 font-ContentT text-xl">
					<Button onClick={FinalButtonClick}>{finalButtonText}</Button>
				</div>
			</div>
		</Modal>
	</article>
{/if}
<slot />
