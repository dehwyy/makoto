import { TwirpFetchTransport } from '@protobuf-ts/twirp-transport'
import { HashmapClient as GeneratedHashmapClient } from "../../generated/hashmap/hashmap.client"
import {SERVER_PORTS} from "@makoto/config"

const transport = new TwirpFetchTransport({
	baseUrl: `http://localhost:${SERVER_PORTS.HASHMAP}/twirp`,
})

const HashmapClient = new GeneratedHashmapClient(transport)

export { HashmapClient}
