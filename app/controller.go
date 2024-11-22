package app

type (
	Controller interface {
		Routes(*Server)
	}
)
