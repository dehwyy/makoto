import path from 'path'
import { config } from '$config'

export enum Tokens {
	access = 'auth-token',
	refresh = 'refresh-token'
}

export const USER_ID = 'user_id'
export const CONFIG = config(path.resolve('../../config'))
