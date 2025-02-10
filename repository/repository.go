package repository

import "fetch-assessment/rules"

type Repository interface {
	SaveReceipt(rules.Receipt) (string, error)
	LoadReceipt(string) (rules.Receipt, error)
}
