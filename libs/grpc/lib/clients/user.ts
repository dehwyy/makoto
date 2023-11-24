import { TwirpFetchTransport } from '@protobuf-ts/twirp-transport'
import {  UserRPCClient as GeneratedUserClient } from "@makoto/grpc/generated/user"

const transport = new TwirpFetchTransport({
	baseUrl: (process.env.TWIRP_GATEWAY_URL || "http://localhost:4000") + "/user",
})

const UserInfoClient = new GeneratedUserClient(transport)

export { UserInfoClient}
