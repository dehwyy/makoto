import type { RpcOptions, UnaryCall, NextUnaryFn, MethodInfo } from '@protobuf-ts/runtime-rpc';
interface Cookie {
    get(key: string): string | undefined;
}
export declare class RpcInterceptors {
    static WithToken(cookies: Cookie, key?: string): {
        interceptUnary(next: NextUnaryFn, method: MethodInfo<any, any>, input: object, options: RpcOptions): UnaryCall<object, object>;
    };
    static AddAuthorizationHeader(header_value: string | undefined): {
        interceptUnary(next: NextUnaryFn, method: MethodInfo, input: object, options: RpcOptions): UnaryCall;
    };
}
export {};
