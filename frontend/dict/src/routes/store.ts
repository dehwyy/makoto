import { derived, writable } from 'svelte/store'

interface Item {
	wordId: string
	word: string
	value: string
	extra: string
	tags: {
		tagId: string
		text: string
	}[]
}
const ItemsStore = writable<Item[]>([])

export const SetItems = (items: Item[]) => {
	items.reverse()
	ItemsStore.set(items)
}
export const AddItem = (item: Item) => {
	ItemsStore.update(items => [item, ...items])
}
export const RemoveItemById = (id: string) => {
	ItemsStore.update(items => items.filter(item => item.wordId !== id))
}

export const FilterStore = writable('')

export const FilteredItems = derived([ItemsStore, FilterStore], ([items, filter]) => {
	return items.filter(item => {
		const keyWord = new RegExp(
			filter
				.split('')
				.map(w => w + '[a-zа-я\\s]*')
				.join(''),
			'ig'
		)

		const wordMatch = item.word.match(keyWord)
		const translateMatch = item.value.match(keyWord)
		const extraMatch = item.extra?.match(keyWord)

		return wordMatch || translateMatch || extraMatch
	})
})
