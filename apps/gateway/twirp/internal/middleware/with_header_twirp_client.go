// reads the `Authorization` header from response (as twirp doesn't give access to it directly)
package middleware

import "net/http"

type twirpClientRoundTripper struct {
	AuthorizationHeader string
}

func newTwirpClientRoundTripper() *twirpClientRoundTripper {
	return &twirpClientRoundTripper{}
}

func (tripper *twirpClientRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	response, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		return response, err
	}

	tripper.AuthorizationHeader = response.Header.Get(_AuthorizationHeader)

	return response, nil
}
