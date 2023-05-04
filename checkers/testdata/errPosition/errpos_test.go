package checker_test

type tt struct {
	Field string
}

/*! The last return parameter is not an error (checker_test/WrongOrderPublic) */
func WrongOrderPublic() (error, tt) {
	return nil, tt{}
}

/*! The last return parameter is not an error (checker_test/wrongOrder) */
func wrongOrder() (error, bool) {
	return nil, false
}

/*! The last return parameter is not an error (checker_test/wrongOrderWithStruct) */
func wrongOrderWithStruct() (error, tt) {
	return nil, tt{}
}

func correctOrder() (tt, error) {
	return tt{}, nil
}

func CorrectOrder() (tt, error) {
	return tt{}, nil
}
