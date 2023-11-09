import { derived, writable } from 'svelte/store'
import { Tags, TagsStore, type TagInitial } from './tags-store'
import { GetItems } from '$lib/api/fetches'

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
	async ([_, filter, specified_tags]) => {
		try {
			const { data } = await GetItems({
				userId: '',
				part: 0,
				partSize: 50,
				query: filter,
				tags: specified_tags
					.filter(tag => tag.selectedMode != 0) // only specified tags
					.map(tag => ({ text: tag.text, include: tag.selectedMode == 1 })) // if tag.selectedMode == 1 -> include. Else -> exclude
			})

			return data?.items ?? []
		} catch {
			// error would appear when SSR (it's ok)
			return []
		}
	}
)
