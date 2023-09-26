<script>
	import Input from '../input.svelte'
	import Modal from '../modal.svelte'

	$: isEdit = false
</script>

<Modal isOpen={isEdit} close={() => (isEdit = !isEdit)}>
	<div class="p-5 flex flex-col gap-y-7">
		<p class="text-2xl text-center font-Content mb-2">New <span class="text-primary">word</span></p>
		<Input placeholder="word" />
		<Input placeholder="translate" />
		<Input placeholder="extra" />
		<button class="mt-2 text-xl font-Jua bg-accent px-5 py-2 rounded-xl">Add</button>
	</div>
</Modal>
<button
	on:click={() => (isEdit = !isEdit)}
	class="w-1/2 text-2xl px-5 py-2 bg-accent rounded-xl self-start font-Jua">Add new word</button>
