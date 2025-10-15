package pair

type ResponseReason string

const (
	ResponseReasonUserDenial ResponseReason = "userdenial"
	ResponseReasonTimeOut    ResponseReason = "timeout"
	ResponseReasonUnknown    ResponseReason = "unknown"
)
