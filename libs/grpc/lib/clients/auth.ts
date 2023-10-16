import { TwirpFetchTransport } from '@protobuf-ts/twirp-transport'
import { AuthClient as GeneratedAuthClient } from "../../generated/auth/auth.client"

const transport = new TwirpFetchTransport({
	baseUrl: `http://localhost:5001/twirp`,
})

const AuthClient = new GeneratedAuthClient(transport)

export { AuthClient }
