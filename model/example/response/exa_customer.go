package response

import "github.com/zouchangfu/QanLong/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
