package main

import "testing"

func TestNewBasicUser(t *testing.T) {
	expectedName := "John Doe"
	have := []string{"read"}
	notHave := []string{"edit", "ban_user", "delete", "manage_roles"}
	u1 := User(NewBasicUser(expectedName))

	for _, v := range have {
		if !u1.HasPermission(v) {
			t.Errorf("BasicUser don't have expected permission %s", v)
		}
	}
	for _, v := range notHave {
		if u1.HasPermission(v) {
			t.Errorf("BasicUser have unexpected permission %s", v)
		}
	}
	if u1.GetUserName() != "John Doe" {
		t.Errorf("Unexpected user name %s, expected %s", u1.GetUserName(), expectedName)
	}
}
func TestNewModerator(t *testing.T) {
	expectedName := "John Doe"
	have := []string{"read", "edit", "ban_user"}
	notHave := []string{"delete", "manage_roles"}
	u1 := User(NewModerator(expectedName))

	for _, v := range have {
		if !u1.HasPermission(v) {
			t.Errorf("Moderator don't have expected permission %s", v)
		}
	}
	for _, v := range notHave {
		if u1.HasPermission(v) {
			t.Errorf("Moderator have unexpected permission %s", v)
		}
	}
	if u1.GetUserName() != "John Doe" {
		t.Errorf("Unexpected user name %s, expected %s", u1.GetUserName(), expectedName)
	}
}
func TestNewAdmin(t *testing.T) {
	expectedName := "John Doe"
	have := []string{"read", "edit", "ban_user", "delete", "manage_roles"}
	notHave := []string{}
	u1 := User(NewAdmin(expectedName))

	for _, v := range have {
		if !u1.HasPermission(v) {
			t.Errorf("Admin don't have expected permission %s", v)
		}
	}
	for _, v := range notHave {
		if u1.HasPermission(v) {
			t.Errorf("Admin have unexpected permission %s", v)
		}
	}
	if u1.GetUserName() != "John Doe" {
		t.Errorf("Unexpected user name %s, expected %s", u1.GetUserName(), expectedName)
	}
}

func TestHasPermission(t *testing.T) {

	tests := []struct {
		name string
		user User
		want bool
	}{
		{
			name: "read",
			user: User(NewBasicUser("John Doe")),
			want: true,
		},
		{
			name: "edit",
			user: User(NewBasicUser("John Doe")),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.HasPermission(tt.name); got != tt.want {
				t.Errorf("HasPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserName(t *testing.T) {
	tests := []struct {
		name string
		user User
		want string
	}{
		{
			name: "BasicUser",
			user: User(NewBasicUser("John Doe")),
			want: "John Doe",
		},
		{
			name: "Moderator",
			user: User(NewModerator("Joh Doe")),
			want: "Joh Doe",
		},
		{
			name: "Admin",
			user: User(NewAdmin("Doe John")),
			want: "Doe John",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetUserName(); got != tt.want {
				t.Errorf("GetUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRole(t *testing.T) {
	tests := []struct {
		name string
		user User
		want string
	}{
		{
			name: "BasicUser",
			user: User(NewBasicUser("John Doe")),
			want: "BasicUser",
		},
		{
			name: "Moderator",
			user: User(NewModerator("Joh Doe")),
			want: "Moderator",
		},
		{
			name: "Admin",
			user: User(NewAdmin("Doe John")),
			want: "Admin",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.GetRole(); got != tt.want {
				t.Errorf("GetUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
