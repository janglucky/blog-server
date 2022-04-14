package web

type Status int32

func (s Status) Error() string {
	return s.String()
}

func (s Status) String() string {
	if v, ok := StatusName[s]; ok {
		return v
	}
	return ""
}

const (
	// 0-999 common status
	STATUS_OK               Status = 0
	STATUS_PARAMS_INVALID   Status = 1
	STATUS_NOT_LOGIN        Status = 2
	STATUS_ERROR			Status = 3
	STATUS_NO_CONTENT       Status = 204
	STATUS_INTERNAL_ERROR   Status = 500
	STATUS_FUNCTION_TIMEOUT Status = 504
)

var StatusName = map[Status]string {
	// common status
	STATUS_OK:               "ok",
	STATUS_PARAMS_INVALID:   "params.invalid",
	STATUS_NOT_LOGIN:        "user.not login",
	STATUS_NO_CONTENT:       "no.content",
	STATUS_ERROR:            "error",
	STATUS_INTERNAL_ERROR:   "internal.error",
	STATUS_FUNCTION_TIMEOUT: "internal.function.timeout",
}