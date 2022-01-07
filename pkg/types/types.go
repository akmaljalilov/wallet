package types

type Account struct {
	ID int64
}

type ErrAccountNotFound struct {

}

func (e ErrAccountNotFound) Error() string {
	return "account not found"
}