package user

import (
	"strings"
	"time"

	"github.com/mjaydip/go-lms/server/fileio"
)

const userFile = "users.txt"
const layoutISO = "2006-01-02"

// Users type provides interaction of user
type Users struct {
	userList []User
}

// LoadUsers creates user list and load data from data source
func (u *Users) LoadUsers() {
	u.userList = make([]User, 0)
	userStr := fileio.ReadFile(userFile)
	lines := strings.Split(userStr, "\n")
	for _, line := range lines {
		elems := strings.Split(line, ",")
		dob, _ := time.Parse(layoutISO, elems[5])
		usr := User{
			UserID:    elems[0],
			FirstName: elems[1],
			LastName:  elems[2],
			Password:  elems[3],
			Mobile:    elems[4],
			DOB:       dob,
		}
		u.userList = append(u.userList, usr)
	}
}

// SaveUsers saves current state of userList to file
func (u *Users) SaveUsers() {
	s := ""
	for _, user := range u.userList {
		s += user.String()
	}

	fileio.WriteFile(userFile, []byte(s))
}

// AddUser adds a new user to the list
func (u *Users) AddUser(usr *User) {
	u.userList = append(u.userList, *usr)
	u.SaveUsers()
}

// GetUsers provides list of users
func (u *Users) GetUsers() []User {
	return u.userList
}
