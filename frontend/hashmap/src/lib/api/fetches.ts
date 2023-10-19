import { TypedFetch as tp } from '@makoto/lib/typed-fetch'

export const RemoveItem = tp.Create<{
	itemId: number
}>('/api/items/remove')

export const EditItem = tp.Create<{
	itemId: number
	key: string
	value: string
	extra: string
	tags: string[]
}>('/api/items/edit')

export const CreateItem = tp.Create<{
	key: string
	value: string
	extra: string
	tags: string[]
}>('/api/items/create')
