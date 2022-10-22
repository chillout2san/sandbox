package types

type ErrType = int

const (
	NO_ERROR int = iota
	REQUEST_ERROR
	RESPONSE_ERROR
)
