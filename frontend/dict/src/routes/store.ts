import { derived, writable } from 'svelte/store'

interface Item {
	id: number
	word: string
	translate: string
	extra?: string
}

export const ItemsStore = writable<Item[]>([])
export const AddItem = (item: Item) => {
	ItemsStore.update(items => [...items, item])
}
export const RemoveItemById = (id: number) => {
	ItemsStore.update(items => items.filter(item => item.id !== id))
}

export const FilterStore = writable('')

export const FilteredItems = derived([ItemsStore, FilterStore], ([items, filter]) => {
	return items.filter(item => {
		const keyWord = new RegExp(
			filter
				.split('')
				.map(w => w + '[a-z\\s]*')
				.join(''),
			'ig'
		)

		const wordMatch = item.word.match(keyWord)
		const translateMatch = item.translate.match(keyWord)
		const extraMatch = item.extra?.match(keyWord)

		return wordMatch || translateMatch || extraMatch
	})
})
