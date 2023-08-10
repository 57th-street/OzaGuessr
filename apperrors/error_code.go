package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	// サービス層で発生したエラー
	GetDataFailed         ErrCode = "S001"
	InsertDataFailed      ErrCode = "S002"
	HashPasswordFailed    ErrCode = "S003"
	ComparePasswordFailed ErrCode = "S004"
	CheckDataExistFailed  ErrCode = "S005"
	NAData                ErrCode = "S006"
	GenerateTokenFailed   ErrCode = "S007"
	UpdateDataFailed      ErrCode = "S008"

	// リクエストエラー
	ReqBodyDecodeFailed ErrCode = "R001"
	DataAlreadyExists   ErrCode = "R002"

	//　認証エラー
	RequiredAuthorizationHeader ErrCode = "A001"
	Unauthorized                ErrCode = "A002"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &AppError{ErrCode: code, Message: message, Err: err}
}
