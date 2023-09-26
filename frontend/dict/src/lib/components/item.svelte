<script lang="ts">
	export let word: string
	export let translate: string
	export let extra = ''

	// $: isOpen = false
</script>

<div
	aria-hidden="true"
	class={`flex flex-col select-none items-center bg-base-200 w-full px-5 py-2 rounded-xl border-2 border-base-100 overflow-hidden`}>
	<div
		style="grid-template-columns: 20fr 60fr 20fr 1fr;"
		class="grid font-Jua text-xl text-center w-full">
		<!-- word -->
		<div>{word}</div>

		<!-- translate -->
		<div class="">{translate}</div>

		<!-- extra -->
		<div>{extra}</div>

		<!-- close button -->
		<div>
			<span
				class="cursor-pointer hover:text-red-500 transition-all font-ContentT select-none p-2 mr-[-0.5rem]"
				>X</span>
		</div>
	</div>

	<!-- hidden part -->
	<!-- <div class={`flex justify-center w-full transition-all`}>
		<p>No information yet</p>
	</div> -->
</div>
