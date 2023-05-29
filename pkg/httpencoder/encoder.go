package httpencoder

import (
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	nethttp "net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// ErrorEncoder encode error response
func ErrorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	// print error stack trace information.
	fmt.Printf("%+v\n", err)
	se := errors.FromError(err)
	res := &Response{
		Code:    1001,
		Success: false,
		Data:    "",
		Message: se.Error(),
	}

	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(res)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	_, _ = w.Write(body)
}

// ResponseEncoder encode response struct
func ResponseEncoder(w http.ResponseWriter, r *http.Request, data interface{}) error {
	res := &Response{
		Code:    1000,
		Success: true,
		Data:    data,
		Message: "success",
	}

	msRes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(msRes)
	return nil
}
