<script lang="ts">
	//

	import Add from '$lib/components/index/add.svelte'
	import Item from '$lib/components/item.svelte'
	import Input from '$lib/components/input.svelte'

	const items = [
		{ word: '感じる', translate: 'чувствовать', extra: '感' },
		{ word: '選択', translate: 'выбор', extra: '' }
	]
</script>

<main class="py-20 w-[70%] mx-auto flex flex-col gap-y-24 items-center">
	<section class="w-1/2 mx-auto">
		<Input placeholder="Enter word to find..." />
	</section>
	<section class="self-start w-[80%] mx-auto mb-[-3rem]">
		<Add />
	</section>
	<section class="w-[80%] flex flex-col gap-y-10 items-center">
		<div
			style="grid-template-columns: 2fr 6fr 2fr;"
			class="grid w-full bg-base-200 border-2 border-secondary text-center text-2xl rounded-xl px-5 py-5 font-ContentT shadow-md shadow-secondary">
			<p>word</p>
			<p class="ml-[-2%]">translate</p>
			<p>extra</p>
		</div>
		<div class="w-full flex flex-col gap-y-5">
			{#each items as item}
				<Item word={item.word} translate={item.translate} extra={item.extra} />
			{/each}
		</div>
	</section>
</main>
