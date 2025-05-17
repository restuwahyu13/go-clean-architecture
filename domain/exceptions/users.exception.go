package exc

import inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"

type usersException struct{}

func NewUsersException() inf.IUsersException {
	return usersException{}
}

func (u usersException) Login(key string) string {
	err := make(map[string]string)

	err["user_not_found"] = "User is not exist in our system"
	err["invalid_password"] = "Invalid email or password"

	return err[key]
}
