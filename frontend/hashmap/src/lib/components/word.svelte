<script lang="ts">
	import { fade } from 'svelte/transition'

	export let word: string
	export let translate: string
	export let extra = ''

	export let onCloseButtonClick = () => {}

	// $: isOpen = false
</script>

<div transition:fade={{ duration: 500, delay: 100 }} aria-hidden="true" class="word_wrapper">
	<div class="word_wrapper_content">
		<!-- word -->
		<div>{word}</div>

		<!-- translate -->
		<div>{translate}</div>

		<!-- extra -->
		<div>{extra}</div>

		<!-- close button -->
		<div>
			<button on:click={onCloseButtonClick} class="word_close-btn">X</button>
		</div>
	</div>
</div>

<style lang="scss">
	.word {
		&_wrapper {
			@apply flex flex-col select-none items-center bg-base-200 w-full px-5 py-2 rounded-xl border-2 border-base-100 overflow-hidden;

			&_content {
				@apply grid font-Jua text-xl place-items-center  w-full;
				grid-template-columns: 20fr 60fr 20fr 1fr;
			}
		}
		&_close-btn {
			@apply cursor-pointer hover:text-red-500 transition-all font-ContentT select-none p-2 mr-[-0.5rem];
		}
	}
</style>
