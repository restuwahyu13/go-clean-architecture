package helper

import (
	"math"

	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
)

func Pagination(limit, offset, total int) *opt.Pagination {
	res := new(opt.Pagination)

	res.Limit = limit
	res.Page = offset
	res.TotalPage = math.Ceil(float64(total) / float64(limit))
	res.TotalData = total

	return res
}
