package opt

type (
	Error struct {
		Name    string `json:"name,omitempty"`
		Message string `json:"message"`
		Code    int    `json:"code,omitempty"`
		Stack   any    `json:"stack,omitempty"`
	}

	Response struct {
		StatCode   float64 `json:"stat_code"`
		Message    any     `json:"message,omitempty"`
		ErrCode    any     `json:"err_code,omitempty"`
		ErrMsg     any     `json:"err_msg,omitempty"`
		Pagination any     `json:"pagination,omitempty"`
		Data       any     `json:"data,omitempty"`
		Errors     any     `json:"errors,omitempty"`
		Info       Info    `json:"info"`
	}

	Info struct {
		Host         any `json:"host"`
		Path         any `json:"path"`
		Method       any `json:"method"`
		Protocol     any `json:"protocol"`
		IPAddress    any `json:"ip_address"`
		UserAgent    any `json:"user_agent"`
		Timestamp    any `json:"timestamp"`
		ResponseTime any `json:"response_time"`
	}
)
