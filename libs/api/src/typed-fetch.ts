export class TypedFetch {
  // create a typed fetch
  static Create = <T>(url: string, method = 'POST') =>
	(data: T) => {
		return fetch(url, {
			method: method,
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(data)
		})
	}

  // read body of request (on server side f.e.)
  static Get = async <T extends (...args: any) => any>(req: Request, fn: T) => {
	return (await req.json()) as Awaited<Parameters<typeof fn>[0]>
  }
}
