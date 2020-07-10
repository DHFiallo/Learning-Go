//Darius Fiallo
//June 16, 2020
//Bank Application

package main

type owner struct {
	id         string
	address    string
	entityType string
}

type bankAccount struct {
	accNum   string
	accOwner owner
	bal      float64
	intRate  float64
	accType  string
}

type wallet struct {
	walletID    string
	accounts    []bankAccount
	walletOwner owner
}

func (a *bankAccount) withdraw(amount float64) {
	if a.balance >= amount {
		a.balance = a.balance - amount
		return true
	}
	return false
}

func (a *bankAccount) deposit(amount float64) {
	
	if amt > 0 {
		a.balance = a.balance + amt
		return true;
	}
	return false;

}

func (a *bankAccount) applyInterest() {
	acctype := a.accOwner.entityType
	if strings.ToUpper(acctype) == "BUSINESS" {
		typeAcc = strings.ToUpper(a.accType)
		
		switch typeAcc {
		case "CHECKING":
			a.balance = a.balance * 1.005
		case "INVESTMENT":
			a.balance = a.balance * 1.01
		case "SAVING":
			a.balance = a.balance * 1.02
		}

	}
	if strings.ToUpper(acctype) == "INDIVIDUAL" {
		typeAcc = strings.ToUpper(a.accType)
		
		switch typeAcc {
		case "CHECKING":
			a.balance = a.balance * 1.01
		case "INVESTMENT":
			a.balance = a.balance * 1.02
		case "SAVING":
			a.balance = a.balance * 1.05
		}
	}
}

func (a *bankAccount) wire(b *bankAccount, amt float64) {
	a.withdraw(amt)
	b.deposit(amt)
}

func (e *)
func main() {

	//ind1 = rahul, darius
	//bus1 = syed

	var user string
	fmt.Println("What's your name?")
	fmt.Scanln(&user)
	user = strings.ToUpper(user)

	switch user {
	case "RAHUL":
		ptr1 := *ind1;
	case "SYED":
		ptr1 := *bus1;
	case "DARIUS":
		ptr1 := *ind2;
	}

	for {
		var num string
		fmt.Println("Would you like to..")
		fmt.Println("1. View accounts")
		fmt.Println("2. View wallet")
		fmt.Scanln(&num)

		if(num == 1) {
			for i, s := range displayAccounts() {
				
			}
		}
}
