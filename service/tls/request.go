package tls

func retryErrorCheck(err error) (bool, error) {
	if err == nil {
		return false, nil
	}

	switch e := err.(type) {
	case *BadResponseError:
		if e.HTTPCode == 500 || e.HTTPCode == 502 || e.HTTPCode == 503 {
			return true, e
		}
	default:
		return false, e
	}

	return false, err
}
