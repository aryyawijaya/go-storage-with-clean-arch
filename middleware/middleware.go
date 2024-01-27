package middleware

type Middleware struct {
	wrapper Wrapper
}

func NewMiddleware(wrapper Wrapper) *Middleware {
	return &Middleware{
		wrapper: wrapper,
	}
}
