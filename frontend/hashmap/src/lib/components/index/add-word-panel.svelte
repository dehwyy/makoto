<script lang="ts">
	import { scale, fade } from 'svelte/transition'
	import { AddItem } from '$lib/stores/items-store'
	import { CreateItem } from '$lib/api/fetches'
	import { Input, Button, Modal } from 'makoto-ui-svelte'
	import Placeholder from '$lib/components/placeholder.svelte'

	let key = ''
	let value = ''
	let extra = ''

	let time_tag = ''
	let tags: string[] = []

	$: isEdit = false
	$: isInAddingTagsMode = false

	const addWord = async () => {
		// adding new item and request to update db
		AddItem({
			id: 0,
			key,
			value,
			extra,
			tags: tags.map(tag => ({ tagId: 0, text: tag }))
		})

		CreateItem({ key, value, extra, tags })

		// reset
		isEdit = false
		key = ''
		value = ''
		extra = ''

		time_tag = ''
		tags = []
	}

	const saveTag = () => {
		tags = Array.from(new Set<string>(tags).add(time_tag))
		time_tag = ''

		isInAddingTagsMode = false
	}
</script>

{#if isEdit}
	<article
		transition:fade={{ duration: 150, delay: 0 }}
		class="fixed top-0 left-0 right-0 bottom-0 z-50">
		<Modal base_width={400} isOpen={isEdit} close={() => (isEdit = !isEdit)}>
			<div class="modal_wrapper">
				<p class="modal_heading">
					New <span class="underline">Word</span>
				</p>
				<Input bind:value={key} placeholder="key">
					<Placeholder placeholder="key" />
				</Input>
				<Input bind:value placeholder="value">
					<Placeholder placeholder="value" />
				</Input>
				<Input bind:value={extra} placeholder="extra">
					<Placeholder placeholder="extra" />
				</Input>
				<div class="modal_tags_wrapper">
					{#each tags as tag}
						<div transition:scale={{ duration: 500, delay: 0 }} class="modal_tags_content">
							{tag}
							<button
								on:click={() => (tags = tags.filter(t => t !== tag))}
								class="modal_tags_content_close"
								><span class="pt-[2px]">X</span>
							</button>
						</div>
					{/each}
					<button
						on:click={() => (isInAddingTagsMode = !isInAddingTagsMode)}
						class="modal_tags_add-tag">
						+
					</button>
					{#if isInAddingTagsMode}
						<div
							transition:fade={{ duration: 300, delay: 0 }}
							class="absolute z-30 right-10 left-10 bottom-28 flex gap-x-5 text-md font-[600]">
							<Input autofocus={true} bind:value={time_tag} placeholder="tag">
								<Placeholder placeholder="tag" />
							</Input>
							<div
								on:click={saveTag}
								aria-hidden="true"
								class="cursor-pointer flex items-end rounded-full">
								<span class="hover:text-green-400 transition-all duration-300 font-Jua text-lg"
									>Save</span>
							</div>
						</div>
					{/if}
				</div>
				<div class="w-full mt-16 font-ContentT text-xl">
					<Button onClick={addWord}>Add</Button>
				</div>
			</div>
		</Modal>
	</article>
{/if}
<div class="w-full md:w-1/2">
	<Button onClick={() => (isEdit = !isEdit)}>
		<span class="text-lg font-ContentT font-[600]">Add new word</span></Button>
</div>

<style lang="scss">
	.modal {
		&_wrapper {
			@apply p-10 flex flex-col gap-y-7 w-full font-Content;
		}
		&_heading {
			@apply text-2xl text-center font-Content text-white;
		}
		&_tags {
			&_wrapper {
				@apply flex flex-wrap gap-x-2 gap-y-2;
			}
			&_content {
				@apply flex-auto py-2 px-7 rounded-full bg-base-200 text-center font-Content relative pr-9;

				&_close {
					@apply hover:text-red-500 transition-all duration-300 cursor-pointer bg-base-100 bg-opacity-80 text-gray-400 rounded-full absolute right-[5px] top-[7.5px] h-[25px] w-[25px] text-sm font-[700] font-Jua px-1 grid place-items-center;
				}
			}
			&_add-tag {
				@apply select-none cursor-pointer hover:text-green-400 transition-all duration-300 font-[700] text-lg grid place-items-center w-[40px] h-[40px] bg-base-200 rounded-full;
			}
		}
	}
</style>
