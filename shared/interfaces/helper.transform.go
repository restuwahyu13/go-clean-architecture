package inf

type ITransform interface {
	ReqToRes(src, dest any) error
	ResToReq(src, dest any) error
}
