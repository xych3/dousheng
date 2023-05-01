package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                int32 = 0
	ServiceErrCode             int32 = 10001
	ParamErrCode               int32 = 10002
	UserAlreadyExistErrCode    int32 = 10003
	AuthorizationFailedErrCode int32 = 10004
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
)

func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}

func MConvertErr(err error) map[string]interface{} {
	Err := ConvertErr(err)
	m := make(map[string]interface{})
	m["status_code"] = Err.ErrCode
	m["status_msg"] = Err.ErrMsg
	return m
}
