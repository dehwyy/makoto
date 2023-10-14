import type { RpcOptions, UnaryCall, NextUnaryFn, MethodInfo } from '@protobuf-ts/runtime-rpc'

export class RpcInterceptors {
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
