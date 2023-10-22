import { derived, writable } from 'svelte/store'

export interface TagInitial {
	tagId: number
	text: string
	usages: number
}

interface Tag extends TagInitial {
	selectedMode: number
}

export const TagsStore = writable<Tag[]>([])
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

	static Set(tags: TagInitial[]) {
		TagsStore.set(tags.map(tag => ({ ...tag, selectedMode: this.OptionMode.startValue })))
	}

	static DescreaseCount(tags: string[]) {
		TagsStore.update(ts =>
			ts.map(t => ({ ...t, usages: tags.includes(t.text) ? t.usages - 1 : t.usages }))
		)
	}

	static Add(tag: string) {
		// adding only tag doesn't exist yet
		TagsStore.update(old_tags => {
			const idx = old_tags.findIndex(t => t.text === tag)
			if (idx === -1) {
				return [
					...old_tags,
					{ text: tag, tagId: 0, usages: 1, selectedMode: this.OptionMode.startValue }
				]
			}

			old_tags[idx].usages++

			return old_tags
		})
	}

	// used for dynamic CSS class
	static GetTagValue(value: number) {
		return this.OptionMode.getValue(value)
	}

	static Toggle(tag_id: number, tags: Tag[]) {
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

		return tag.text.match(query) && tag.usages > 0
	})
})
