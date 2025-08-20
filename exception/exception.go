package http_error

// AuthException 身份验证异常
type AuthException struct {
	Message string
}

// DbException 数据库事务异常
type DbException struct {
	Message string
}

// ConfirmException 需要点击确认的异常
type ConfirmException struct {
	Message string
}

// BackException 页面后退异常
type BackException struct {
	Message string
}
