import artifact from '$houdini/artifacts/SignUp'
import { MutationStore } from '../runtime/stores/mutation'

export class SignUpStore extends MutationStore {
	constructor() {
		super({
			artifact,
		})
	}
}
