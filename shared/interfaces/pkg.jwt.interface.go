package inf

import (
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
)

type IJsonWebToken interface {
	Sign(prefix string, body any) (*opt.SignMetadata, error)
}
