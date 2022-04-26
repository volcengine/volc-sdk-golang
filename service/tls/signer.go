package tls

import "net/http"

func Sign(request *http.Request, cred Credentials) *http.Request {
	query := request.URL.Query()
	request.URL.RawQuery = query.Encode()

	prepareRequestV4(request)

	meta := new(metadata)
	meta.service, meta.region = cred.Service, cred.Region

	// Task 1
	hashedCanonReq := hashedCanonicalRequestV4(request, meta)

	// Task 2
	stringToSign := stringToSign(request, hashedCanonReq, meta)

	// Task 3
	signingKey := signingKey(cred.SecretAccessKey, meta.date, meta.region, meta.service)
	signature := signature(signingKey, stringToSign)

	request.Header.Set("Authorization", buildAuthHeader(signature, meta, cred))

	if cred.SessionToken != "" {
		request.Header.Set("X-Security-Token", cred.SessionToken)
	}

	return request
}
