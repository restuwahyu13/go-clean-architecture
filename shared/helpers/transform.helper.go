package helper

import (
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
)

type transform struct{}

func NewTransform() inf.ITransform {
	return transform{}
}

func (h transform) ReqToRes(src, dest any) error {
	helper := NewParser()

	srcByte, err := helper.Marshal(src)
	if err != nil {
		return err
	}

	if err = helper.Unmarshal(srcByte, dest); err != nil {
		return err
	}

	return nil
}

func (h transform) ResToReq(src, dest any) error {
	helper := NewParser()

	srcByte, err := helper.Marshal(src)
	if err != nil {
		return err
	}

	if err = helper.Unmarshal(srcByte, dest); err != nil {
		return err
	}

	return nil
}
