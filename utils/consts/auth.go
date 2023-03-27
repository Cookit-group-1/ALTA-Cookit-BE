package consts

// Error Response
const (
	AUTH_ErrorBind            string = "error bind data"
	AUTH_ErrorHash            string = "error hash password"
	AUTH_ErrorComparePassword string = "password not matched"
	AUTH_ErrorCreateToken     string = "error create token"
	AUTH_ErrorRole            string = "role must be VerifiedUser"
	AUTH_ErrorNewPassword     string = "new password and confirmation password is not equal"
	AUTH_ErrorEmptyPassword   string = "old password, new password and confirmation password field cannot be empty"
	AUTH_SecurePassword       string = "password must be at least 8 characters long, must contain uppercase letters, must contain lowercase letters, must contain numbers, must not be too general"
)

// Success Response
const (
	AUTH_SuccessCreate string = "success create account"
	AUTH_SuccessLogin  string = "login success"
)
