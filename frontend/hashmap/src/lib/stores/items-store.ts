import { derived, writable } from 'svelte/store'
import { TagsStore } from './tags-store'

// Initial Store
interface Item {
	id: number
	key: string
	value: string
	extra: string
	tags: {
		tagId: number
		text: string
	}[]
}

// stores
const ItemsStore = writable<Item[]>([])
export const ItemsFilterQueryStore = writable('') // (input) query

export class Items {
	static Set(items: Item[]) {
		ItemsStore.set(items)
	}

	static Edit(item: Item) {
		ItemsStore.update(items => items.map(it => (it.id === item.id ? item : it)))
	}

	static Add(item: Item) {
		ItemsStore.update(items => [item, ...items])
	}

	static RemoveById(id: number) {
		ItemsStore.update(items => items.filter(item => item.id !== id))
	}
}

// Filtered Items which depends on $ItemsStore, $ItemsFilterQueryStore and $TagsStore
export const FilteredItems = derived(
	[ItemsStore, ItemsFilterQueryStore, TagsStore],
	([items, filter, specified_tags]) => {
		return items.filter(item => {
			const keyWord = new RegExp(
				filter
					.split('')
					.map(w => w + '[a-zа-я\\s]*')
					.join(''),
				'ig'
			)

			const wordMatch = item.key.match(keyWord)
			const translateMatch = item.value.match(keyWord)
			const extraMatch = item.extra?.match(keyWord)

			// creating dict to importve performance from O(n^2) to O(n)
			const tags = {} as Record<string, boolean>
			for (const tag of item.tags) {
				tags[tag.text.toLowerCase()] = true
			}

			for (const tag of specified_tags) {
				const tag_text = tag.text.toLowerCase()

				switch (tag.selectedMode) {
					case 1:
						// clarify whether SelectedTag is in ItemTags
						// if not => invalidate word
						if (!tags[tag_text]) {
							return false
						}
						break
					case 2:
						// SelectedTag is NOT in ItemTags
						if (tags[tag_text]) {
							return false
						}
						break
				}
			}

			return wordMatch || translateMatch || extraMatch
		})
	}
)
