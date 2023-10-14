import { TwirpFetchTransport } from '@protobuf-ts/twirp-transport'
import { AuthClient as GeneratedAuthClient } from '$rpc/generated/auth/auth.client'
import { CONFIG } from '$lib/const'

const transport = new TwirpFetchTransport({
	baseUrl: `http://localhost:${CONFIG.ports.auth}/twirp`,
	interceptors: [{}]
})

const AuthClient = new GeneratedAuthClient(transport)

export { AuthClient }
