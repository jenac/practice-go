package expersist

import (
	"testing"
)

func Test(t *testing.T) {
	users := []User{
		{
			Username: "james",
			Email:    "james@s.com",
			Role:     "admin",
		},
		{
			Username: "tim",
			Email:    "tim@S.com",
			Role:     "user",
		},
	}

	SaveTxt(users)
	SaveJSON(users)
	SaveCSV(users)
}
