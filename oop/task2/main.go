package main

import "fmt"

type User interface {
	GetUserName() string
	HasPermission(permission string) bool
	GetRole() string
}

type Permissions struct {
	permissions map[string]struct{}
}

func (p *Permissions) HasPermission(permission string) bool {
	_, ok := p.permissions[permission]
	return ok
}

type BasicUser struct {
	username string
	role     string
	Permissions
}

func (u *BasicUser) GetUserName() string {
	return u.username
}
func (u *BasicUser) GetRole() string {
	return u.role
}
func NewBasicUser(username string) *BasicUser {
	return &BasicUser{
		username:    username,
		role:        "BasicUser",
		Permissions: Permissions{map[string]struct{}{"read": {}}},
	}
}

type Moderator struct {
	BasicUser
}

func NewModerator(username string) *Moderator {
	u := NewBasicUser(username)
	u.role = "Moderator"
	u.permissions["edit"] = struct{}{}
	u.permissions["ban_user"] = struct{}{}
	return &Moderator{*u}
}

type Admin struct {
	Moderator
}

func NewAdmin(username string) *Admin {
	u := NewModerator(username)
	u.role = "Admin"
	u.permissions["delete"] = struct{}{}
	u.permissions["manage_roles"] = struct{}{}
	return &Admin{*u}
}

func main() {
	var u1, u2, u3 User
	u1 = NewBasicUser("John Doe")
	u2 = NewModerator("Jane Doe")
	u3 = NewAdmin("Job Dom")
	fmt.Println(u1, u2, u3)
	fmt.Println(u1.GetUserName(), u2.GetUserName(), u3.GetUserName())
	fmt.Println(u1.GetRole(), u2.GetRole(), u3.GetRole())
	fmt.Println(u1.HasPermission("edit"), u2.HasPermission("edit"), u3.HasPermission("edit"))
}
