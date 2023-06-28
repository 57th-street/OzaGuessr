package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	GetDataFailed         ErrCode = "S001"
	InsertDataFailed      ErrCode = "S002"
	HashPasswordFailed    ErrCode = "S003"
	ComparePasswordFailed ErrCode = "S004"
	CheckDataExistFailed  ErrCode = "S005"
	NAData                ErrCode = "S006"

	ReqBodyDecodeFailed ErrCode = "R001"
	DataAlreadyExists   ErrCode = "R002"
	LoginFailed         ErrCode = "R003"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &AppError{ErrCode: code, Message: message, Err: err}
}
