package order

import "errors"

const (
	NotValidate = iota
	Pending
	IsPaid
	Process
	Completed
	Cancelled
)

var (
	ErrWrongFormat     = errors.New("has a wrong format")
	ErrOrderIsCreated  = errors.New("order is created")
	ErrOrderIsCanceled = errors.New("order is canceled")
)

const MaxLimit = 50
