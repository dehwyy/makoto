import { derived, writable } from 'svelte/store'

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
const ItemsStore = writable<Item[]>([])

export const SetItems = (items: Item[]) => {
	console.log(items)
	items.reverse()
	ItemsStore.set(items)
}
export const AddItem = (item: Item) => {
	ItemsStore.update(items => [item, ...items])
}
export const RemoveItemById = (id: number) => {
	ItemsStore.update(items => items.filter(item => item.id !== id))
}

// Tags Store
interface OptionInitial {
	tagId: number
	text: string
	usages?: number
}

interface Option extends OptionInitial {
	selectedMode: number
}

const OptionMode = {
	/**
	 * NO - not stated
	 * SELECTED - add this tag to query
	 * PROHIBITED - prohibit tag from query
	 */
	_values: ['NO', 'SELECTED', 'PROHIBITED'] as const,
	startValue: 0,
	getValue(value: number): (typeof this._values)[number] {
		return this._values[value % this._values.length]
	},
	increased(current_value: number): number {
		return (current_value + 1) % this._values.length
	}
}

export const TagsStore = writable<Option[]>([])
export const SetTags = (tags: OptionInitial[]) => {
	TagsStore.set(tags.map(tag => ({ ...tag, selectedMode: OptionMode.startValue })))
}

export const GetTagValue = (value: number) => OptionMode.getValue(value)

export const ToggleTag = (tag_id: number, tags: Option[]) => {
	// "!" at the end cuz it cannot be undefined
	const current_option_index = tags.indexOf(tags.find(tag => tag.tagId === tag_id)!)

	// toggle `selectMode`
	TagsStore.set(
		tags.map((tag, i) =>
			i != current_option_index
				? tag
				: { ...tag, selectedMode: OptionMode.increased(tag.selectedMode) }
		)
	)
}

// Filtered Tags
export const FilterTagQueryStore = writable('')

export const FilteredTags = derived([TagsStore, FilterTagQueryStore], ([tags, filter_query]) => {
	return tags.filter(tag => {
		const query = new RegExp(
			filter_query
				.split('')
				.map(w => w + '[a-zа-я\\s]*')
				.join(''),
			'ig'
		)

		return tag.text.match(query)
	})
})

// Filtered Words
export const FilterStore = writable('')

export const FilteredItems = derived(
	[ItemsStore, FilterStore, TagsStore],
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
						// clarify whether SelectedTag is in WordTags
						// if not => invalidate word
						if (!tags[tag_text]) {
							return false
						}
						break
					case 2:
						// SelectedTag is NOT in WordTags
						if (tags[tag_text]) {
							return false
						}
				}
			}

			return wordMatch || translateMatch || extraMatch
		})
	}
)
