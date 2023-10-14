import type { SignInRequest } from '$rpc/generated/auth/auth'

// TODO: API would change
type OAuth2Provider = 'local' | 'google'

export class RpcPayloads {
	static SignIn(args: { provider: OAuth2Provider; code: string }): SignInRequest {
		const { code, provider } = args

		return {
			authMethod: {
				oneofKind: 'oauth2',
				oauth2: {
					provider,
					code
				}
			}
		}
	}
}
