package response

import (
	"github.com/bytedance/sonic"
	"webapi/pkg/exception"
)

type DataUnwrapper interface {
	UnwrapData(interface{}) error
}

// Standard Response
type CommonResponse struct {
	ResponseCode    int                        `json:"code"`
	ResponseMessage string                     `json:"message"`
	Errors          *exception.ExceptionErrors `json:"errors,omitempty"`
	Data            any                        `json:"data,omitempty"`
	RequestID       string                     `json:"request_id,omitempty"`
	Path            string                     `json:"path,omitempty"`
}

func (resp *CommonResponse) UnwrapData(target interface{}) error {
	bs, err := sonic.Marshal(resp.Data)
	if err != nil {
		return err
	}

	if err := sonic.Unmarshal(bs, target); err != nil {
		return err
	}

	return nil
}
