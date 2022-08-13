package user

type UserType int

const (
	USER_TYPE_ADMIN UserType = iota + 1
	USER_TYPE_TEACHER
	USER_TYPE_STUDENT
)

func (userType UserType) String() string {
	switch userType {
	case USER_TYPE_ADMIN:
		return "ADMIN"
	case USER_TYPE_TEACHER:
		return "TEACHER"
	case USER_TYPE_STUDENT:
		return "STUDENT"
	default:
		return "UNKNOWN"
	}
}

func (userType UserType) Index() int {
	if userType > USER_TYPE_STUDENT || userType < USER_TYPE_ADMIN {
		return int(-1)
	}
	return int(userType)
}
