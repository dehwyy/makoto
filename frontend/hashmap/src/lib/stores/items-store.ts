import { derived, writable } from 'svelte/store'
import { Tags, TagsStore, type TagInitial } from './tags-store'

// Initial Store
export interface Item {
	id: number
	key: string
	value: string
	extra: string
	tags: string[]
}

// stores
const ItemsStore = writable<Item[]>([])
export const ItemsFilterQueryStore = writable('') // (input) query

export class Items {
	static Set(items: Item[]) {
		ItemsStore.set(Array.from(items).reverse())
	}

	static SetIdToItemAfterCreate(id: number) {
		ItemsStore.update(items =>
			items.map(it => {
				if (it.id == 0) {
					return { ...it, id }
				} else {
					return it
				}
			})
		)
	}

	static Edit(item: Item) {
		item.tags.forEach(tag => Tags.Add(tag))

		ItemsStore.update(items =>
			items.map(it => {
				const is_needed_item = it.id == item.id
				is_needed_item && Tags.DescreaseCount(it.tags)

				return is_needed_item ? item : it
			})
		)
	}

	static Add(item: Item) {
		item.tags.forEach(tag => Tags.Add(tag))

		ItemsStore.update(items => [item, ...items])
	}

	static RemoveById(id: number) {
		ItemsStore.update(items =>
			items.filter(item => {
				const is_needed_item = item.id == id
				is_needed_item && Tags.DescreaseCount(item.tags)

				return !is_needed_item
			})
		)
	}
}

// Filtered Items which depends on $ItemsStore, $ItemsFilterQueryStore and $TagsStore
export const FilteredItems = derived(
	[ItemsStore, ItemsFilterQueryStore, TagsStore],
	([items, filter, specified_tags]) => {
		//

		// ? Searching using reserved search params like `item?key=word` or `item?id=123`
		const reserved_search_params = {
			// id: number
			// key: string
			//
		} as Record<string, string>
		const common_words: string[] = []

		filter.split(' ').forEach(w => {
			// if word doesn't match the `reserved search params`
			if (!w.match(/^(item)\?[(id)|(key)|(value)|(extra)]+=/)) {
				common_words.push(w)
				return
			}

			const [prefix, value] = w.split('=')
			const key = prefix.split('?')[1]
			reserved_search_params[key] = value
		})

		const match_word = new RegExp(
			common_words
				.join(' ')
				.split('')
				.map(w => w + '[a-zа-я\\s]*')
				.join(''),
			'ig'
		)

		return items.filter(item => {
			//

			// if ID (unique) exists in query BUT not match current Item => return false
			const id = reserved_search_params.id
			if (id && Number(id) != item.id) {
				return false
			}

			const k = reserved_search_params.key
			if (k && k != item.key) {
				return false
			}

			const v = reserved_search_params.value
			if (v && v != item.value) {
				return false
			}

			const e = reserved_search_params.extra
			if (e && e != item.extra) {
				return false
			}

			const wordMatch = item.key.match(match_word)
			const translateMatch = item.value.match(match_word)
			const extraMatch = item.extra?.match(match_word)

			// creating dict to importve performance from O(n^2) to O(n)
			const tags = {} as Record<string, boolean>
			for (const tag of item.tags) {
				tags[tag.toLowerCase()] = true
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
