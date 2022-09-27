package sdk

import (
	"github.com/golangtips/yuque/sdk/internal"
	"github.com/golangtips/yuque/sdk/intf"
)

type IYuQue = intf.IYuQue

func NewYuQue(baseURL, token, userAgent, namespace string) intf.IYuQue {
	return internal.NewYuQue(baseURL, token, userAgent, namespace)
}
