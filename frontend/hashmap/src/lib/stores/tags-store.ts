import { derived, writable } from 'svelte/store'

interface OptionInitial {
	tagId: number
	text: string
	usages?: number
}

interface Option extends OptionInitial {
	selectedMode: number
}

export const TagsStore = writable<Option[]>([])
export class Tags {
	private static OptionMode = {
		/**
		 * NO - not stated
		 * SELECTED - add this tag to query
		 * PROHIBITED - prohibit tag from query
		 */
		values: ['NO', 'SELECTED', 'PROHIBITED'] as const,
		startValue: 0,
		getValue(value: number): (typeof this.values)[number] {
			return this.values[value % this.values.length]
		},
		increased(current_value: number): number {
			return (current_value + 1) % this.values.length
		}
	}

	static Set(tags: OptionInitial[]) {
		TagsStore.set(tags.map(tag => ({ ...tag, selectedMode: this.OptionMode.startValue })))
	}

	// TODO: оно вообще надо?
	static GetTagValue(value: number) {
		return this.OptionMode.getValue(value)
	}

	static Toggle(tag_id: number, tags: Option[]) {
		// "!" at the end cuz it cannot be undefined
		const current_option_index = tags.indexOf(tags.find(tag => tag.tagId === tag_id)!)

		// toggle `selectMode`
		// iterating using `.map` and if index == ClickedTagIndex (fn arg) => update `selectMode` (values in `this.OptionMode`) ELSE => do nothing
		TagsStore.set(
			tags.map((tag, i) =>
				i != current_option_index
					? tag
					: { ...tag, selectedMode: this.OptionMode.increased(tag.selectedMode) }
			)
		)
	}
}
// export const SetTags = (tags: OptionInitial[]) => {
// 	TagsStore.set(tags.map(tag => ({ ...tag, selectedMode: OptionMode.startValue })))
// }

// export const GetTagValue = (value: number) => OptionMode.getValue(value)

// export const ToggleTag = (tag_id: number, tags: Option[]) => {
// 	// "!" at the end cuz it cannot be undefined
// 	const current_option_index = tags.indexOf(tags.find(tag => tag.tagId === tag_id)!)

// 	// toggle `selectMode`
// 	TagsStore.set(
// 		tags.map((tag, i) =>
// 			i != current_option_index
// 				? tag
// 				: { ...tag, selectedMode: OptionMode.increased(tag.selectedMode) }
// 		)
// 	)
// }

// Filtered Tags
export const FilterTagQueryStore = writable('') // (input) query

// Computed based on $TagsStore and $FilterTagsQueryStore
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
