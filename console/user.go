package console

import (
	"bufio"
	"fmt"
	"time"

	"github.com/mjaydip/go-lms/server/user"
)

func addUserOp(s *bufio.Scanner) {
	u := user.User{}
	fmt.Println("Enter User Details: ")
	fmt.Print("UserID: ")
	s.Scan()
	u.UserID = s.Text()
	fmt.Print("First Name: ")
	s.Scan()
	u.FirstName = s.Text()
	fmt.Print("Last Name: ")
	s.Scan()
	u.LastName = s.Text()
	fmt.Print("Password: ")
	s.Scan()
	u.Password = s.Text()
	fmt.Print("Mobile: ")
	s.Scan()
	u.Mobile = s.Text()
	fmt.Print("Date of birth (yyyy-mm-dd): ")
	s.Scan()
	dob, _ := time.Parse(user.LayoutISO, s.Text())
	u.DOB = dob

	user.GetInst().AddUser(&u)
	fmt.Println("User added successfully!")
}

// HandleUserOp handles all user actions for console app
func HandleUserOp(op string, s *bufio.Scanner) {
	switch op {
	case "a", "A":
		addUserOp(s)
	}
}
