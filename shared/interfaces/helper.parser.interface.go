package inf

import (
	"io"

	"github.com/google/uuid"
)

type IParser interface {
	ToString(v any) string
	ToInt(v any) (int, error)
	ToFloat(v any) (float64, error)
	ToByte(v any) ([]byte, error)
	Marshal(source any) ([]byte, error)
	Unmarshal(src []byte, dest any) error
	Decode(src io.Reader, dest any) error
	Encode(src io.Writer, dest any) error
	FromUUID(s string) uuid.UUID
	FromNullUUID(s string) uuid.NullUUID
	DecimalToFloat(n int64) float64
}
