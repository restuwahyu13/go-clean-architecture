package con

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	cons "github.com/restuwahyu13/go-clean-architecture/shared/constants"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
)

func SqlConnection(ctx context.Context, env dto.Environtment) (*sqlx.DB, error) {
	sqlx.BindDriver(cons.POSTGRES, sqlx.DOLLAR)

	return sqlx.ConnectContext(ctx, cons.POSTGRES, env.POSTGRES.URL)
}
