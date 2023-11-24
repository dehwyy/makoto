import type { GetItemsResponse, Item } from '@makoto/grpc/.ts/generated/hashmap/hashmap'
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

export const CreateItem = tp.Create<
	{
		userId?: string // TODO:
		key: string
		value: string
		extra: string
		tags: string[]
	},
	{
		itemId: number
	}
>('/api/items/create')

export const GetItems = tp.Create<
	{
		userId: string
		part: number
		partSize: number
		query: string
		tags: {
			text: string
			include: boolean
		}[]
	},
	{ items: Item[] }
>('/api/items')
