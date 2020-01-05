package user

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// LayoutISO is a format to parse date time values
const LayoutISO = "2006-01-02"

// User type describes a system user like admin, librarian
type User struct {
	UserID    string         `json:"userID"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Password  string         `json:"password"`
	Mobile    string         `json:"mobile"`
	DOB       mysql.NullTime `json:"dateOfBirth"`
}

func (u *User) String() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%v", u.UserID, u.FirstName, u.LastName, u.Password, u.Mobile, u.DOB.Time.Format(LayoutISO))
}

// Print prints user in a nice format
func (u *User) Print() {
	fmt.Printf("%s, %s, %s, %s, %s, %v\n", u.UserID, u.FirstName, u.LastName, u.Password, u.Mobile, u.DOB.Time.Format(LayoutISO))
}
