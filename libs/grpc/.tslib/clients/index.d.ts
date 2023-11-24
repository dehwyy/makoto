import { MakotoCookiesInterface as Cookies } from "@makoto/lib/cookies";
declare const TwirpClient: (cookies: Cookies) => {
    Authorization: import("../../.ts/generated/auth/auth.client").AuthRPCClient;
    Hashmap: import("../../.ts/generated/hashmap/hashmap.client").HashmapRPCClient;
    UserInfo: import("../../.ts/generated/user/user.client").UserRPCClient;
};
export { TwirpClient as SafeTwirpClient };
