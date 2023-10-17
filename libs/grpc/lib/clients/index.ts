import { RpcStatus } from "@protobuf-ts/runtime-rpc"
import {AuthClient as AC} from "./auth"
import  {HashmapClient as HS} from "./hashmap"

interface TwirpClient {
  methods: {
    localName: string
  }[]
}

const CreateSafeClient = <T extends TwirpClient>(client: T) => new Proxy(client, {
  get: (target, prop: string, rec) => {
    // if this is a RpcServiceMethod
    if (target["methods"].map(m => m.localName).includes(prop)) {

        // making a Proxy which would listen on fn call
        return new Proxy(Reflect.get(target, prop, rec) as Function, {
          apply: async (target, thisArg, args) => {
            try {

              // try to request
              return await Reflect.apply(target as any, thisArg, args)
            } catch (e) {

              // if err occured
              return {
                response: {},
                status: {
                  code: "400",
                  detail: String(e)
                } as RpcStatus
              }
            }
          }
        })
    }

    return Reflect.get(target, prop, rec)
  }
})


const AuthClient = CreateSafeClient(AC)
const HashmapClient = CreateSafeClient(HS)

export {AuthClient, HashmapClient, AuthClient as SafeAuthClient, HashmapClient as SafeHashmapClient}
