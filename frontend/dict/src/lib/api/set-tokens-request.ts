interface Tokens {
	access_token: string
	refresh_token: string
}

export const SetTokensFromResponse = (tokens: Tokens | undefined) => {
	if (!tokens) return

	const { access_token, refresh_token } = tokens

	fetch('/api/set-tokens', {
		method: 'POST',
		body: JSON.stringify({
			access_token: access_token,
			refresh_token: refresh_token
		})
	})
}
