package errno

/**
custom some error code.
error type include: code 、message
error code format:
1 									00 							02
服务级错误						服务模块代码						具体错误代码
(1为系统级错误，2 为普通错误)   一个大型系统的服务模块通常不超过2位数    2位数，防止一个模块定制过多的错误码导致后期不好维护
通常由用户非法操作引起				超过则应该拆分

code = 0 : 正确返回
code > 0 : 错误返回
错误包括：系统级错误码、服务级错误码
按服务模块代码将错误分类
错误码均>=0
*/
var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
)
