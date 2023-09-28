import type { Cache as InternalCache } from './cache/cache';
import type { CacheTypeDef } from './generated';
import { Cache } from './public';
export * from './client';
export * from './lib';
export declare function graphql<_Payload>(str: TemplateStringsArray): _Payload;
export declare const cache: Cache<CacheTypeDef>;
export declare function getCache(): InternalCache;
