package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func _TestAllUser(t *testing.T) {
	users, err := AllUser()
	fmt.Println(users, err)
	
	b,_ := json.MarshalIndent(users,"","		")
	fmt.Println(string(b))
}

func TestFindUserRoleByUserId(t *testing.T) {
	userrole, err := FindUserRoleByUserId(2)
	fmt.Println(userrole, err)
	
	b,_ := json.MarshalIndent(userrole,"","		")
	fmt.Println(string(b))
}

func _TestAllUserRole(t *testing.T) {
	userrole, err := AllUserRole()
	fmt.Println(userrole, err)
	
	b,_ := json.MarshalIndent(userrole,"","		")
	fmt.Println(string(b))
}