<script lang="ts">
	import { fly } from 'svelte/transition'
	import { circOut, circIn } from 'svelte/easing'
	import {
		TagsStore as Tags,
		ToggleTag,
		GetTagValue,
		FilterTagQueryStore,
		FilteredTags
	} from '$lib/stores/items-store'
	import { Input } from 'makoto-ui-svelte'
	import Placeholder from './placeholder.svelte'

	let isOpen = false
	const toggleOpenMode = (e: MouseEvent) => {
		e.stopPropagation()
		isOpen = !isOpen
	}
</script>

<svelte:body on:click={() => (isOpen = false)} />

<article class="relative w-full sm:w-3/4">
	<button
		on:click={toggleOpenMode}
		class="px-7 py-1 font-Jua select-none flex justify-between items-center w-full bg-base-200 hover:bg-base-100 transition-all duration-300 rounded-xl">
		<span class="text-lg">Tags</span>
		<span class="text-gray-300 text-sm pt-1">Any ></span>
	</button>
	{#if isOpen}
		<div
			on:click={e => e.stopPropagation()}
			aria-hidden="true"
			in:fly={{ duration: 400, x: 500, easing: circOut }}
			out:fly={{ duration: 250, x: -500, easing: circIn }}
			class="absolute top-[50px] left-0 right-0 z-50 bg-base-300 py-2 rounded-xl border-solid">
			<div class="flex flex-col items-start overflow-y-scroll h-[200px]">
				<div class="flex w-full px-5 mb-3">
					<Input bind:value={$FilterTagQueryStore} placeholder="Find tag...">
						<Placeholder placeholder="Find tag..." />
					</Input>
				</div>
				{#each $FilteredTags as tag}
					<div
						aria-hidden="true"
						on:click={() => ToggleTag(tag.tagId, $Tags)}
						class="flex items-center w-full pl-3 pr-5 hover:bg-base-200 cursor-pointer option_wrapper">
						<div class="h-[40px] w-[40px] grid place-items-center">
							<button class="h-[18px] w-[18px]">
								<p
									class={`${GetTagValue(
										tag.selectedMode
									)} min-h-full min-w-full transition-all duration-300 rounded-md border-solid border-base-300 border`} />
							</button>
						</div>
						<button class="py-2 px-5 font-Content">{tag.text}</button>
						<div class="ml-auto">
							{tag.usages || 0}
						</div>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</article>

<style lang="scss">
	.NO {
		@apply bg-base-200;
	}
	.SELECTED {
		@apply bg-green-500;
	}
	.PROHIBITED {
		@apply bg-red-500;
	}
	.option {
		&_wrapper {
			&:hover .NO {
				@apply border-white;
			}
		}
	}
</style>
