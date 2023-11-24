import { redirect } from '@sveltejs/kit'

// todo: to -> @makoto/libs/lib
export class HandleResponse {
	static Handle(status_code: string): never | boolean {
		if (status_code != 'OK') {
			// TODO:
			throw redirect(307, 'http://localhost:3001')
		}

		return true
	}
}
