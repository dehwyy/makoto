import { AuthClient as AC } from "./auth";
import { HashmapClient as HS } from "./hashmap";
import { UserInfoClient as UIC } from "./user";
import { MakotoCookies } from "@makoto/lib/cookies";
const CreateSafeClient = (client, cookies) => new Proxy(client, {
    get: (target, prop, rec) => {
        // if this is a RpcServiceMethod
        if (target["methods"].map(m => m.localName).includes(prop)) {
            // making a Proxy which would listen on fn call
            return new Proxy(Reflect.get(target, prop, rec), {
                apply: async (target, thisArg, args) => {
                    try {
                        // try to request
                        const response = await Reflect.apply(target, thisArg, args);
                        const { authorization: authorization_header, ...headers } = response.headers;
                        const new_response = { ...response, headers };
                        let token = "";
                        // if header is not empty
                        if (authorization_header != "") {
                            const split_token = authorization_header?.split(" ");
                            // if after keyword there is token (f.e. Bearer <token>)
                            if (split_token && split_token.length > 1) {
                                token = split_token[1];
                            }
                        }
                        if (token.length) {
                            MakotoCookies.setGlobal(cookies, "token", token);
                        }
                        return new_response;
                    }
                    catch (e) {
                        // if err occured
                        return {
                            response: {},
                            status: {
                                code: "400",
                                detail: String(e)
                            }
                        };
                    }
                }
            });
        }
        return Reflect.get(target, prop, rec);
    }
});
const TwirpClient = (cookies) => ({
    Authorization: CreateSafeClient(AC, cookies),
    Hashmap: CreateSafeClient(HS, cookies),
    UserInfo: CreateSafeClient(UIC, cookies)
});
export { TwirpClient as SafeTwirpClient };
