import { TypedFetch as tp } from '@makoto/lib/typed-fetch'

export const SignInFetch = tp.Create<{
	username: string
	password: string
}>('/api/v1/auth/local/sign-in')

export const SignUpFetch = tp.Create<{
	username: string
	password: string
	email: string
}>('/api/v1/auth/local/sign-up')

export const SignInByToken = tp.Create<void>('/api/v1/auth/token')
