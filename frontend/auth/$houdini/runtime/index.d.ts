import { SignUpStore } from "../plugins/houdini-svelte/stores/SignUp";
import type { Cache as InternalCache } from "./cache/cache";
import type { CacheTypeDef } from "./generated";
import { Cache } from "./public";
export * from "./client";
export * from "./lib";

export function graphql(
    str: "\n\t\tmutation SignUp($username: String!, $password: String!, $question: String!, $answer: String!) {\n\t\t\tsignUp(\n\t\t\t\tinput: { username: $username, password: $password, question: $question, answer: $answer }\n\t\t\t) {\n\t\t\t\tuserId\n\t\t\t\ttokens {\n\t\t\t\t\taccess_token\n\t\t\t\t\trefresh_token\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t"
): SignUpStore;

export declare function graphql<_Payload>(str: TemplateStringsArray): _Payload;
export declare const cache: Cache<CacheTypeDef>;
export declare function getCache(): InternalCache;