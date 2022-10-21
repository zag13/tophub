package httpz

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

// NopCloseRequestDecoder 增加 io.NopCloser 方法。
func NopCloseRequestDecoder(r *http.Request, v interface{}) error {
	codec, ok := khttp.CodecForRequest(r, "Content-Type")
	if !ok {
		return errors.BadRequest("CODEC", fmt.Sprintf("unregister Content-Type: %s", r.Header.Get("Content-Type")))
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.BadRequest("CODEC", err.Error())
	}
	if len(data) == 0 {
		return nil
	}
	if err = codec.Unmarshal(data, v); err != nil {
		return errors.BadRequest("CODEC", fmt.Sprintf("body unmarshal %s", err.Error()))
	}
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	return nil
}

// GetRequestFromCtx 获取http.request
func GetRequestFromCtx(ctx context.Context) (request *http.Request, getReqOk bool) {
	if transPort, ok := transport.FromServerContext(ctx); ok {
		if transPort.Kind() == transport.KindHTTP {
			if info, ok := transPort.(*khttp.Transport); ok {
				request = info.Request()
				getReqOk = true
			}
		}
	}
	return
}
