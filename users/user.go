package users

import (
	"fmt"
	"time"

	"github.com/salernolucas93/godesdecero/modelos"
)

func AltaUsuario() {
	user := new(modelos.User)
	user.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(user)
}
