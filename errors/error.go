package errors

//Error Error type
type Error int

//Error Error interface
func (e Error) Error() string {
	return errMsg[int(e)]
}


