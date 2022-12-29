package consumer

const (
	ConsumeFromBegin = "begin"
	ConsumeFromEnd   = "end"
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
)
