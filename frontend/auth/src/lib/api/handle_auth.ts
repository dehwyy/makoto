import type { RpcStatus } from '@protobuf-ts/runtime-rpc'
import { redirect } from '@sveltejs/kit'

interface HandlerProps {
	status: RpcStatus
}

export class HandleAuth {
	static Handle({ status }: HandlerProps): Response | never {
		const { code, detail } = status || { code: '500', detail: 'Internal Server Error' }

		if (code.startsWith('5') || code.startsWith('4')) {
			return new Response(null, {
				status: Number(code),
				statusText: detail.split('\n')[0] || 'Internal Server Error'
			})
		}

		throw redirect(307, '/?redirect=/')
	}
}
