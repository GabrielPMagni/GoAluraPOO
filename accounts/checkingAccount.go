package accounts

type CheckingAccount struct {
	Owner         string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *CheckingAccount) Deposit(value float64) (bool, float64) {
	if value > 0 {
		c.Balance += value
		return true, c.Balance
	}
	return false, c.Balance
}

func (c *CheckingAccount) Withdraw(value float64) (bool, float64) {
	if value > 0 && value < c.Balance {
		c.Balance -= value
		return true, c.Balance
	}
	return false, c.Balance
}

func (c *CheckingAccount) Transfer(value float64, destinyAccount *CheckingAccount) bool {
	if c.Balance > 0 && value < c.Balance {
		c.Balance -= value
		destinyAccount.Balance += value
		return true
	}
	return false
}
