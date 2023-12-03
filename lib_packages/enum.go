package my_lib

type Role int

const (
	User Role = iota
	Manager
	Admin
)

// const (
// 	User Role = 1 << iota
// 	Manager
// 	Admin
// 	RoleMask = (1 << (iota)) - 1
// )

func (r Role) String() string {
	switch r {
	case User:
		return "User"
	case Manager:
		return "Manager"
	case Admin:
		return "Admin"
	default:
		return "Invalid role"
	}
}
