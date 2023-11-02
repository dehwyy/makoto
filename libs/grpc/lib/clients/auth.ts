import { TwirpFetchTransport } from '@protobuf-ts/twirp-transport'
import { AuthRPCClient as GeneratedAuthClient } from "../../generated/auth/auth.client"

const transport = new TwirpFetchTransport({
	baseUrl: (process.env.TWIRP_GATEWAY_URL || "localhost:4000") + "/authorization",
})

const AuthClient = new GeneratedAuthClient(transport)

export { AuthClient }
