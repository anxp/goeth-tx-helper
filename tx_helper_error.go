package goeth_tx_helper

// https://stackoverflow.com/questions/61407054/using-standard-library-only-how-to-wrap-error-in-custom-error

/*
	NOTE! If you want to use approach like:

	if err != nil {

		// Why pointer to pointer?
		// https://www.reddit.com/r/golang/comments/txi397/cant_understand_parameters_of_errorsas/?rdt=56934
		// https://stackoverflow.com/questions/69447919/panic-errors-target-must-be-interface-or-implement-error-in-go
		externalError := &goeth_tx_helper.ExternalErrorWrapper{}

		if errors.As(err, &externalError) {
			fmt.Printf("\n"+
				"\n========== EXTERNAL ERROR =========="+
				"\n%s"+
				"\n===================================="+
				"\n", err.Error())
		}
	}

	DO NOT embed ExternalErrorWrapper error into another error like fmt.Errorf("failed to get tick (%d) data: %s", tick, err), where err is ExternalErrorWrapper
	INSTEAD use AddContext method() (see below)
*/

import "fmt"

type ExternalErrorWrapper struct {
	OriginalError     error
	OurMessage        string
	AdditionalContext string
}

// Error is mark the struct as an error.
func (e *ExternalErrorWrapper) Error() string {
	template := "" +
		"TxHELPER ERROR: %v\n" +
		"EXTERNAL ERROR: %v"

	errMessage := fmt.Sprintf(template, e.OurMessage, e.OriginalError)

	if e.AdditionalContext != "" {
		template += "\nADDITIONAL CONTEXT: %v"
		errMessage = fmt.Sprintf(template, e.OurMessage, e.OriginalError, e.AdditionalContext)
	}

	return errMessage
}

// Unwrap is used to make it work with errors.Is, errors.As.
func (e *ExternalErrorWrapper) Unwrap() error {
	return e.OriginalError
}

func (e *ExternalErrorWrapper) AddContext(additionalInfo string) {
	e.AdditionalContext += additionalInfo
}

// WrapExternalError to easily create a new error which wraps the given error.
func WrapExternalError(originalErr error, message string) error {
	return &ExternalErrorWrapper{
		OriginalError: originalErr,
		OurMessage:    message,
	}
}
