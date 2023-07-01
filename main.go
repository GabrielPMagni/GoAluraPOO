package main

import (
	"fmt"

	c "github.com/GabrielPMagni/GoAluraPOO/accounts"
)

func main() {
	fooAccount := c.CheckingAccount{Owner: "Foo", AgencyNumber: 100, AccountNumber: 123456, Balance: 120}
	barAccount := c.CheckingAccount{Owner: "Bar", AgencyNumber: 100, AccountNumber: 123456, Balance: 120}

	fooAccount.Deposit(200.0)
	fooAccount.Withdraw(50.0)
	fooAccount.Transfer(60.0, &barAccount)

	fmt.Println(fooAccount)
}
