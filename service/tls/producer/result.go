package producer

type Attempt struct {
	SuccessFlag  bool
	RequestId    string
	ErrorCode    string
	ErrorMessage string
	TimestampMs  int64
}

func newAttempt(success bool, requestId, errorCode, errorMessage string, timestampMs int64) *Attempt {
	return &Attempt{
		SuccessFlag:  success,
		RequestId:    requestId,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		TimestampMs:  timestampMs,
	}
}

type Result struct {
	Attempts    []*Attempt
	SuccessFlag bool
}

func newResult() *Result {
	return &Result{
		Attempts:    []*Attempt{},
		SuccessFlag: false,
	}
}
