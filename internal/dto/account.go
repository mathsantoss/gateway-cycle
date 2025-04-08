package dto

import (
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

type CreateAccountInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
} // CreateAccount represents the data required to create a new account.

type AccountOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} // AccountOutput represents the data returned when an account is retrieved.

func ToAccount(input CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
} // toAccount converts a CreateAccountInput to a domain.Account.

func FromAccount(account *domain.Account) AccountOutput {
	return AccountOutput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
		APIKey:    account.APIKey,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
} // FromAccount converts a domain.Account to an AccountOutput.
