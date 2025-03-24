package data

func NewCheckResult(pass bool, msg string, err error) *CheckResult {
	return &CheckResult{
		Pass: pass,
		Msg:  msg,
		Err:  err,
	}
}

func NewPassResult() *CheckResult {
	return &CheckResult{
		Pass: true,
		Msg:  "",
		Err:  nil,
	}
}

func NewFailResult(msg string, err error) *CheckResult {
	return &CheckResult{
		Pass: false,
		Msg:  msg,
		Err:  err,
	}
}
