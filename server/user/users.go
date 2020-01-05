package user

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitCtrl handles controller initialization
func InitCtrl(mysql *sql.DB, r *gin.Engine) {
	db = mysql

	r.GET("/users", GetAllUsers)
}

// GetAllUsers sets GET route for all users
func GetAllUsers(c *gin.Context) {
	selDB, err := db.Query("SELECT * FROM lms_user")
	if err != nil {
		panic(err.Error())
	}

	res := []User{}
	for selDB.Next() {
		var dob mysql.NullTime
		var userID, fname, lname, pwd, mobile string
		err = selDB.Scan(&userID, &pwd, &fname, &lname, &dob, &mobile)
		if err != nil {
			panic(err.Error())
		}
		usr := User{
			UserID:    userID,
			FirstName: fname,
			LastName:  lname,
			Password:  pwd,
			Mobile:    mobile,
			DOB:       dob,
		}
		res = append(res, usr)
	}

	c.JSON(200, gin.H{
		"data": res,
	})
}
