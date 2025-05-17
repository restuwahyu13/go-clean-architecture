package inf

type (
	IApi interface {
		Middleware()
		Router()
		Listener()
	}
)
