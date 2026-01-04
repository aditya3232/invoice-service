package clients

import (
	"context"
	"fmt"
	"invoice-service/clients/config"
	"net/http"
)

type CustomerClient struct {
	client config.IClientConfig
}

type ICustomerClient interface {
	FindByID(context.Context, int) (*CustomerData, error)
}

func NewCustomerClient(client config.IClientConfig) ICustomerClient {
	return &CustomerClient{client: client}
}

func (u *CustomerClient) FindByID(ctx context.Context, id int) (*CustomerData, error) {
	var response CustomerResponse
	request := u.client.Client().
		Get(fmt.Sprintf("%s/api/v1/customers/%d", u.client.BaseURL(), id))

	resp, _, errs := request.EndStruct(&response)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("customer response: %s", response.Message)
	}

	return &response.Data, nil
}
