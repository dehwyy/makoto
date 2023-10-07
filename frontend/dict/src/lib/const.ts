export enum Tokens {
	access = 'auth-token',
	refresh = 'refresh-token'
}

export const USER_ID = 'user_id'

export enum ContextKeys {
	_NOT_ZERO, // reserved value for 0 as it is the default Number val
	isAuthed
}

export const SERVICES = {
	AUTH: 'http://localhost:3001'
} as const
