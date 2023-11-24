import type { RpcStatus } from '@protobuf-ts/runtime-rpc'
import { redirect } from '@sveltejs/kit'

interface HandlerProps {
	status: RpcStatus
}

export class HandleAuth {
	static Handle({ status }: HandlerProps): Response | never {
		const { code, detail } = status || { code: 'Error', detail: 'Internal Server Error' }

		console.log(`code: ${code} with detail ${detail}`)

		if (code != 'OK') {
			return new Response(null, {
				status: 500,
				statusText: detail.split('\n')[0] || 'Internal Server Error'
			})
		}

		throw redirect(307, '/?redirect=/')
	}
}
