import { TwirpFetchTransport } from '@protobuf-ts/twirp-transport';
import { HashmapRPCClient as GeneratedHashmapClient } from "@makoto/grpc/generated/hashmap";
const transport = new TwirpFetchTransport({
    baseUrl: (process.env.TWIRP_GATEWAY_URL || "http://localhost:4000") + "/hashmap",
});
const HashmapClient = new GeneratedHashmapClient(transport);
export { HashmapClient };
