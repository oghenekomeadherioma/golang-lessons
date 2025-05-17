package main

import (
	"fmt"
	"time"
)

type Account struct {
	Id string
	Currency string
	Balance int
	Owner string
	AccType string
	Opened time.Time
}

func main() {
 sharks_account, err := NewAccount(5000, "$", "Mr Sharks", "34", "Standard", time.Now())
 if err != nil { fmt.Println(err)}
 fmt.Println(sharks_account.GetAccount())
//  fmt.Println(sharks_account.GetBalance())
//  fmt.Println(sharks_account.Withdraw(2000))
//  fmt.Println(sharks_account.GetBalance())
//  fmt.Println(sharks_account.Deposit(90_000))
//  fmt.Println(sharks_account.GetBalance)
}

func NewAccount(balance int,currency,  owner, id, acctype string, opened time.Time) (*Account, error) {
	//_, err := time.Parse(time.RFC3339, opened)
	// if err !=nil {
	// 	return nil, fmt.Errorf("invalid date format %v", err)
	// }
	if currency == "" || balance <= 0 || owner == "" || acctype == "" || opened.After(time.Now()){
		return nil, fmt.Errorf("all fields most be of type required")
	}
	newAccount := Account{
		Currency: currency,
		Balance: balance,
		Owner: owner,
		AccType: acctype,
		Opened: opened,
	}
	return &newAccount, nil
	}

func (acc *Account) Deposit(amount int) error{
	if amount <= 0 { return fmt.Errorf("Ammount can't be zero")}
	return nil
}
func (acc *Account) Withdraw(amount int) error{
	if amount <= 0 { return fmt.Errorf("Ammount can't be zero")}
	if acc.Balance < amount { return fmt.Errorf("Insufficient funds")}
	acc.Balance -= amount; return nil 
}
func (acc *Account) Transfer(amount int, to *Account) error{
	if err:= validate_transaction(amount, acc); err != nil { return err}
	to.Balance += amount
	acc.Balance -= amount
	return nil 
}
func (acc *Account) GetAccount() string {
	return fmt.Sprintf("Owner: %s\nBalance: %d %s\nAccount Type: %s\nOpened: %s\n Account ID: %s\n",
		 acc.Owner, acc.Balance, acc.Currency, acc.AccType, acc.Opened.Format(time.RFC3339), acc.Id,)
}
func (acc *Account) GetBalance() int {
	return acc.Balance
}

func validate_transaction(ammount int, account *Account) error{
	if ammount <= 0 {return fmt.Errorf("Ammount can't be zero")}
	if ammount > account.Balance { return fmt.Errorf("Insufficient funds")}
	return nil
}
	
	