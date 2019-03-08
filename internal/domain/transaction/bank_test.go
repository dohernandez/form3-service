package transaction_test

import (
	"testing"

	"github.com/dohernandez/form3-service/internal/domain/transaction"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestBankAccountValidate(t *testing.T) {
	testCases := []struct {
		scenario        string
		bankAccountFunc func(bankAccount transaction.BankAccount) transaction.BankAccount
		err             error
	}{
		{
			scenario: "Bank account is valid",
			bankAccountFunc: func(bankAccount transaction.BankAccount) transaction.BankAccount {
				return bankAccount
			},
		},
		{
			scenario: "Bank account is invalid due to empty person name",
			bankAccountFunc: func(bankAccount transaction.BankAccount) transaction.BankAccount {
				bankAccount.Name = ""

				return bankAccount
			},
			err: errors.Errorf("name: cannot be blank."),
		},
		{
			scenario: "Bank account is invalid due to empty account number",
			bankAccountFunc: func(bankAccount transaction.BankAccount) transaction.BankAccount {
				bankAccount.Number = ""

				return bankAccount
			},
			err: errors.Errorf("account_number: cannot be blank."),
		},
		{
			scenario: "Bank account is invalid due to empty account name",
			bankAccountFunc: func(bankAccount transaction.BankAccount) transaction.BankAccount {
				bankAccount.AccountName = ""

				return bankAccount
			},
			err: errors.Errorf("account_name: cannot be blank."),
		},
		{
			scenario: "Bank account is invalid due to empty account number code",
			bankAccountFunc: func(bankAccount transaction.BankAccount) transaction.BankAccount {
				bankAccount.AccountNumberCode = ""

				return bankAccount
			},
			err: errors.Errorf("account_number_code: cannot be blank."),
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			bankAccount := tc.bankAccountFunc(transaction.NewBankAccountMock())

			err := bankAccount.Validate()
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAccountValidate(t *testing.T) {
	testCases := []struct {
		scenario    string
		accountFunc func(account transaction.Account) transaction.Account
		err         error
	}{
		{
			scenario: "Account is valid",
			accountFunc: func(account transaction.Account) transaction.Account {
				return account
			},
		},
		{
			scenario: "Account is invalid due to zero bank id",
			accountFunc: func(account transaction.Account) transaction.Account {
				account.ID = 0

				return account
			},
			err: errors.Errorf("bank_id: cannot be zero."),
		},
		{
			scenario: "Account is invalid due to empty bank id code",
			accountFunc: func(account transaction.Account) transaction.Account {
				account.IDCode = ""

				return account
			},
			err: errors.Errorf("bank_id_code: cannot be blank."),
		},
		{
			scenario: "Account is invalid due to empty account number",
			accountFunc: func(account transaction.Account) transaction.Account {
				account.Number = ""

				return account
			},
			err: errors.Errorf("account_number: cannot be blank."),
		},
	}

	for _, tc := range testCases {
		tc := tc // Pinning ranged variable, more info: https://github.com/kyoh86/scopelint
		t.Run(tc.scenario, func(t *testing.T) {
			account := tc.accountFunc(transaction.NewAccountMock())

			err := account.Validate()
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
