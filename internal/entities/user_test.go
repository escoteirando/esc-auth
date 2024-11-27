package entities

import "testing"

func TestRoleType_StringParse(t *testing.T) {
	tests := []struct {
		name string
		r    RoleType
		want string
	}{
		{"user", RoleUser, RoleUserStr},
		{"admin", RoleAdmin, RoleAdminStr},
		{"unknown", RoleUnknown, RoleUnknownStr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("RoleType.String() = %v, want %v", got, tt.want)
			}
			var role RoleType
			if _ = role.Parse(tt.want); role != tt.r {
				t.Errorf("RoleType.Parse(%s) = %v, want %v", tt.want, role, tt.r)
			}
		})
	}
}
