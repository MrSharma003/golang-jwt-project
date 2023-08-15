package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized access to the resoruce")
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid") // doubt: as we are getting both userId and uid from gin.Context, then what is the purpose?
	err = nil

	//this allowing user can access their on data but if the ADMIN is there then no issue
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized access to this resource")
		return err
	}

	err = CheckUserType(c, userType)
	return err

}
