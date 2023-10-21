export class TypedFetch {
  // create a typed fetch
  static Create = <T, R = void>(url: string, method = 'POST') =>
	async (data: T) => {
		const response = await fetch(url, {
			method: method,
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		})

		let response_data: R | undefined = undefined
		if (!response.ok) {
			response_data = await response.json()
		}


		return {
			...response,
			data: response_data
		}
	}

  // read body of request (on server side f.e.)
  static Get = async <T extends (...args: any) => any>(req: Request, fn: T) => {
		return (await req.json()) as Awaited<Parameters<typeof fn>[0]>
  }
}
