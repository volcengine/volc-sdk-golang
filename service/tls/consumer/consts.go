package consumer

const (
	ConsumeFromBegin = "begin"
	ConsumeFromEnd   = "end"

	Delimiter = "-"
)

const (
	nop = iota
	pending
	initializing
	readyToFetch
	fetching
	readyToConsume
	consuming
	backoff
	waitForRestart
)
