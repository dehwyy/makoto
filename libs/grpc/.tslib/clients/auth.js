import { TwirpFetchTransport } from '@protobuf-ts/twirp-transport';
import { AuthRPCClient as GeneratedAuthClient } from "@makoto/grpc/generated/auth";
const transport = new TwirpFetchTransport({
    baseUrl: (process.env.TWIRP_GATEWAY_URL || "http://localhost:4000") + "/authorization",
});
const AuthClient = new GeneratedAuthClient(transport);
export { AuthClient };
