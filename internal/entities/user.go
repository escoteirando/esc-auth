package entities

type (
	UserEntity struct {
		Id       int
		UserName string
		Password string
		PersonId int
		Role     RoleType
	}
	RoleType uint8
)

const (
	RoleUser    RoleType = 1
	RoleAdmin   RoleType = 255
	RoleUnknown RoleType = 0

	RoleUserStr    = "user"
	RoleAdminStr   = "admin"
	RoleUnknownStr = "unknown"
)

func (r RoleType) String() string {
	switch r {
	case RoleUser:
		return RoleUserStr
	case RoleAdmin:
		return RoleAdminStr
	default:
		return RoleUnknownStr
	}
}

func (r *RoleType) Parse(role string) RoleType {
	switch role {
	case RoleUserStr:
		*r = RoleUser
	case RoleAdminStr:
		*r = RoleAdmin
	default:
		*r = RoleUnknown
	}
	return *r
}
