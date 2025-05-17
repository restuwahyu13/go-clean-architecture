package helper

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	cons "github.com/restuwahyu13/go-clean-architecture/shared/constants"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
)

type responseTimer struct {
	startTime time.Time
	http.ResponseWriter
}

var errorCodeMapping = map[int]string{
	http.StatusBadGateway:          "SERVICE_ERROR",
	http.StatusServiceUnavailable:  "SERVICE_UNAVAILABLE",
	http.StatusGatewayTimeout:      "SERVICE_TIMEOUT",
	http.StatusConflict:            "DUPLICATE_RESOURCE",
	http.StatusBadRequest:          "INVALID_REQUEST",
	http.StatusUnprocessableEntity: "INVALID_REQUEST",
	http.StatusPreconditionFailed:  "REQUEST_COULD_NOT_BE_PROCESSED",
	http.StatusForbidden:           "ACCESS_DENIED",
	http.StatusUnauthorized:        "UNAUTHORIZED_TOKEN",
	http.StatusNotFound:            "UNKNOWN_RESOURCE",
	http.StatusInternalServerError: "GENERAL_ERROR",
}

func Version(path string) string {
	return fmt.Sprintf("%s/%s", cons.API, path)
}

func Api(rw http.ResponseWriter, r *http.Request, options opt.Response) {
	rt := &responseTimer{startTime: time.Now(), ResponseWriter: rw}

	response := buildResponse(options, r, rt)
	writeResponse(rt, NewParser(), response)
}

func (rt *responseTimer) WriteHeader(code int) {
	rt.ResponseWriter.WriteHeader(code)
}

func getErrorCode(statusCode int) string {
	if code, exists := errorCodeMapping[statusCode]; exists {
		return code
	}

	return errorCodeMapping[http.StatusInternalServerError]
}

func isEmptyResponse(resp opt.Response) bool {
	return reflect.DeepEqual(resp, opt.Response{})
}

func getProtocol(r *http.Request) string {
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		return "https"
	}
	return "http"
}

func getIPAddress(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")

	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}

	if ip == "" {
		ip = r.RemoteAddr
	}

	return ip
}

func buildResponse(options opt.Response, r *http.Request, rt *responseTimer) opt.Response {
	response := opt.Response{
		StatCode: http.StatusInternalServerError,
		ErrMsg:   cons.DEFAULT_ERR_MSG,
	}

	if isEmptyResponse(options) {
		defaultErrCode := getErrorCode(http.StatusInternalServerError)
		response.ErrCode = &defaultErrCode

		return response
	}

	response = copyResponseFields(options, response)
	setResponseDefaults(&response)

	response.Info = opt.Info{
		Host:         r.Host,
		Protocol:     getProtocol(r),
		Path:         r.URL.Path,
		Method:       r.Method,
		Timestamp:    time.Now().Format(time.RFC3339),
		ResponseTime: fmt.Sprintf("%d ms", time.Since(rt.startTime).Milliseconds()),
		UserAgent:    r.UserAgent(),
		IPAddress:    getIPAddress(r),
	}

	return response
}

func copyResponseFields(source, target opt.Response) opt.Response {
	if source.StatCode != 0 {
		target.StatCode = source.StatCode
	}

	if source.Message != nil {
		target.Message = source.Message
	}

	if source.ErrCode != nil {
		target.ErrCode = source.ErrCode
	}

	if source.ErrMsg != "" {
		target.ErrMsg = source.ErrMsg
	}

	if source.Data != nil {
		target.Data = source.Data
	}

	if source.Errors != nil {
		target.Errors = source.Errors
	}

	if source.Pagination != nil {
		target.Pagination = source.Pagination
	}

	target = opt.Response{
		StatCode:   target.StatCode,
		Message:    target.Message,
		ErrCode:    target.ErrCode,
		ErrMsg:     target.ErrMsg,
		Data:       target.Data,
		Errors:     target.Errors,
		Pagination: target.Pagination,
	}

	return target
}

func setResponseDefaults(response *opt.Response) {
	if response.StatCode == 0 {
		response.StatCode = http.StatusInternalServerError
	}

	if response.StatCode >= http.StatusBadRequest && response.ErrCode == nil {
		defaultErrCode := getErrorCode(int(response.StatCode))
		response.ErrCode = &defaultErrCode
	}

	if response.StatCode >= http.StatusInternalServerError && response.ErrMsg == cons.DEFAULT_ERR_MSG {
		response.ErrMsg = cons.DEFAULT_ERR_MSG
	}
}

func writeResponse(rw http.ResponseWriter, parser inf.IParser, response opt.Response) {
	rw.Header().Set("Content-Type", "application/json")

	statusCode := response.StatCode
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	rw.WriteHeader(int(statusCode))

	if err := parser.Encode(rw, response); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)

		errorResponse := fmt.Sprintf(`{"stat_code":%d, "err_code":"%s", "err_msg":"%s"}`, http.StatusInternalServerError, errorCodeMapping[http.StatusInternalServerError], cons.DEFAULT_ERR_MSG)
		fmt.Fprint(rw, errorResponse)
	}
}
