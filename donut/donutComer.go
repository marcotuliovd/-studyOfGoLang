package donut

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func ComerDonut () {
	fmt.Println("comendo comendo donut")
	

	erro := checkmail.ValidateFormat("bumfaaaaaaaaa")
	fmt.Println(erro)
}