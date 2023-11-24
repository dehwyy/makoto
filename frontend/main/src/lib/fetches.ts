import { TypedFetch as tp } from '@makoto/lib/typed-fetch'

const BASE = '/api/v1'

const u = (url: string) => `${BASE}${url}`

export const UpdateUserInfo = tp.Create<{
  userId: string
  picture: string
  description: string
  darkBg: string
  lightBg: string
  languages: string[]
}>(u('/userinfo/update'))
