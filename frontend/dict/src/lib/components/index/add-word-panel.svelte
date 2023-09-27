<script>
	import { AddItem, ItemsStore } from '../../../routes/store'
	import Button from '../button.svelte'
	import Input from '../input.svelte'
	import Modal from '../modal.svelte'

	let word = ''
	let translate = ''
	let extra = ''

	$: isEdit = false

	const addWord = () => {
		// get len to calculate id
		const itemsLength = $ItemsStore.length

		// adding new item and request to update db
		AddItem({ id: itemsLength + 1, word, translate, extra })
		// db request here

		// reset
		isEdit = false
		word = ''
		translate = ''
		extra = ''
	}
</script>

<Modal isOpen={isEdit} close={() => (isEdit = !isEdit)}>
	<div class="p-10 flex flex-col gap-y-7 w-full">
		<p class="text-2xl text-center font-Content text-white">
			New <span class="underline">Word</span>
		</p>
		<Input bind:value={word} placeholder="word" />
		<Input bind:value={translate} placeholder="translate" />
		<Input bind:value={extra} placeholder="extra" />
		<div class="w-full mt-10">
			<Button onClick={addWord}>Add</Button>
		</div>
	</div>
</Modal>
<div class="w-[40%]">
	<Button onClick={() => (isEdit = !isEdit)}>
		<span class="text-xl">Add new word</span></Button>
</div>
