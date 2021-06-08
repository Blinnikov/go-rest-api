package helpers

type ctxKey int8

const (
	CtxKeyUser ctxKey = iota
	CtxKeyRequestID
)
