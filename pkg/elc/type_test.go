package elc_test

type MixLogEvent struct {
	topic     string
	key       MixLogEventKey
	value     map[string]any
	partition float64
	offset    float64
}

type MixLogEventKey struct {
	service string
	id      string
}
