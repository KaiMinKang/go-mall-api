package errcode

// 定义业务错误码常量。
const (
	// CodeSuccess 表示成功
	CodeSuccess int = 0

	// CodeInvalidParam 表示参数错误
	CodeInvalidParam int = 40001

	// CodeGoodsNotFound 表示商品不存在
	CodeGoodsNotFound int = 40401

	// CodeInternalServerError 表示服务器内部错误
	CodeInternalServerError int = 50000

	//CodeUsernameExists 表示用户名已存在
	CodeUsernameExists int = 40901

	//CodeInvalidCredentials 表示用户名或密码错误
	CodeInvalidCredentials int = 40101

	//CodeUnauthorized 表示未授权
	CodeUnauthorized int = 40102
)
