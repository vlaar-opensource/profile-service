package db

//IfaceDB blabla
type IfaceDB interface {
	AddUser(username string, password string) (UserAccount, error)
	GetByUsername(username string) (UserAccount, error)
}
