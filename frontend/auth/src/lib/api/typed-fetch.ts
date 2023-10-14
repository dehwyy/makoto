const CreateTypedFetch =
	<T>(url: string, method = 'POST') =>
	(data: T) => {
		return fetch(url, {
			method: method,
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		})
	}

export const GetFromRequest = async <T extends (...args: any) => any>(req: Request, fn: T) => {
	return (await req.json()) as Awaited<Parameters<typeof fn>[0]>
}

export const SignInFetch = CreateTypedFetch<{
	username: string
	password: string
}>('/api/v1/auth/local/sign-in')

export const SignUpFetch = CreateTypedFetch<{
	username: string
	password: string
	email: string
}>('/api/v1/auth/local/sign-up')

export const SignInByToken = CreateTypedFetch<void>('/api/v1/auth/token')
