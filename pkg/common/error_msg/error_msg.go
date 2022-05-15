package error_msg

const (
	ErrorUndefined = iota + 10000
	ErrorUnauthorized
	ErrorDBOperation
	ErrorLogin
	ErrorInternal
	UploadImageFailed
	ErrorMinIO
)

const (
	Success = iota + 200
)
