package user_test

import (
	"os"
	"testing"

	"github.com/AlessandroLorenzi/giretti/user"
	"github.com/stretchr/testify/assert"
)

func TestGetByUsername(t *testing.T) {
	os.Chdir("../example_site")
	a := assert.New(t)

	err := user.InitRepo()
	a.NoError(err)
	users := user.GetAll()

	a.Equal(1, len(users))
}
