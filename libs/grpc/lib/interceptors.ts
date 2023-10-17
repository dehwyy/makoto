import type { RpcOptions, UnaryCall, NextUnaryFn, MethodInfo } from '@protobuf-ts/runtime-rpc'

interface Cookie {
	get(key: string): string | undefined
}

export class RpcInterceptors {
	static WithToken(cookies: Cookie, key="token") {
		return this.AddAuthorizationHeader(cookies.get(key))
	}
	static AddAuthorizationHeader(header_value: string | undefined) {
		return {
			interceptUnary(
				next: NextUnaryFn,
				method: MethodInfo,
				input: object,
				options: RpcOptions
			): UnaryCall {
				if (!options.meta) {
					options.meta = {}
				}
				if (header_value) {
					options.meta['Authorization'] = 'Bearer ' + header_value
				}

				return next(method, input, options)
			}
		}
	}
}
