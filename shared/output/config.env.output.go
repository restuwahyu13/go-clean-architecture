package opt

type (
	Application struct {
		ENV          string
		PORT         string
		INBOUND_SIZE int
	}

	Redis struct {
		URL string
	}

	Postgres struct {
		URL string
	}

	Jwt struct {
		SECRET  string
		EXPIRED int
	}

	Environtment struct {
		APP      *Application
		REDIS    *Redis
		POSTGRES *Postgres
		JWT      *Jwt
	}
)
