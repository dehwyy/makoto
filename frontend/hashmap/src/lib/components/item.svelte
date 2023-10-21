<script lang="ts">
	import { fade, scale } from 'svelte/transition'
	import ItemEditor from './item-editor.svelte'
	import { Items } from '$lib/stores/items-store'
	import { EditItem } from '$lib/api/fetches'

	export let item_id = 0
	export let item: string
	export let value: string
	export let extra: string
	export let tags: string[]

	export let removeItem = () => {}

	let isOpenOptions = false
	let isEdit = false

	const Delete = () => {
		isOpenOptions = false
		removeItem()
	}

	const Edit = () => {
		isOpenOptions = false
		isEdit = true
	}

	const Apply = (payload: { key: string; value: string; extra: string; tags: string[] }) => {
		const { key, value, extra, tags } = payload

		// Update $Store
		Items.Edit({
			id: item_id,
			key,
			value,
			extra,
			tags: tags.map(tag => ({ tagId: 0 << 19, text: tag }))
		})

		// Request 2 backend
		EditItem({ itemId: item_id, key, value, extra, tags })
	}

	const Copy2Buffer = (text: string) => {
		navigator.clipboard.writeText(`item?id=${text}`)
	}
</script>

<svelte:body on:click={() => (isOpenOptions = false)} />

<ItemEditor
	defaultValues={{ key: item, value, extra, tags }}
	finalButtonText="Apply"
	bind:isEdit
	onFinalButtonClick={Apply}>
	<div
		transition:fade={{ duration: 250 }}
		aria-hidden="true"
		class="flex flex-col select-none items-center bg-base-200 w-full px-5 py-2 rounded-xl border-2 border-base-100 relative">
		<div class="word_wrapper_content grid font-Jua text-xl place-items-center w-full">
			<div>{item}</div>
			<div class="text-end w-full">{value}</div>
			<div class="">{extra}</div>

			<!-- options button -->
			<div>
				<div
					aria-hidden="true"
					on:click={() => {
						/* seems to be svelte-bug: when using `e.stopPropagation()`,
					it applies to all intances of this component
					(as Item used multiple times, it applies to every instance)

					solution is to create Macrotask,
					but isOpenOptions would be always false as event is not prevented (onClick on svelte:body sets isOpenOptions to `false`)
					to remember current value, `isOpen` exists
					*/
						const isOpen = isOpenOptions

						setTimeout(() => {
							isOpenOptions = !isOpen
						}, 0)
					}}
					class="cursor-pointer hover:text-red-500 transition-all font-ContentT select-none flex flex-col gap-y-1 absolute right-2 top-1/2 -translate-y-1/2 p-2">
					{#each { length: 3 } as _, i}
						<div class="w-[4px] h-[4px] bg-white rounded-full" />
					{/each}
				</div>
				{#if isOpenOptions}
					<div
						transition:scale={{ duration: 300 }}
						class="options_wrapper absolute -right-28 top-2/3 -translate-y-1/2 bg-base-200 rounded-xl py-2 px-1 z-10">
						<button
							on:click={e => {
								e.stopPropagation()
								Copy2Buffer(item_id.toString())
								isOpenOptions = false
							}}
							class="text-info">copy <span class="underline">id</span></button>
						<div />
						<button on:click={Edit} class="text-warning">edit</button>
						<div />
						<button on:click={Delete} class="text-error">delete</button>
					</div>
				{/if}
			</div>
		</div>
	</div>
</ItemEditor>

<style lang="scss">
	.word_wrapper_content {
		grid-template-columns: 20fr 30fr 50fr;
	}

	.options_wrapper {
		& button {
			padding: 0.25rem 1rem 0;
			text-align: center;
			width: 100%;
			@apply rounded-xl;
			&:hover {
				@apply bg-base-100;
			}
		}
		& div {
			@apply bg-base-100 w-full h-[1px] px-1 my-2;
		}
	}
</style>
