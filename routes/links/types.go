package links

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	caeruslinks "github.com/desmos-labs/caerus/routes/links"
	caerustypes "github.com/desmos-labs/caerus/types"
)

type CreateAddressLinkRequest struct {
	// Address is the address of the user for which to create the link
	Address string

	// ChainType represents the chain type to use to create the link
	ChainType caeruslinks.ChainType
}

func NewCreateAddressLinkRequest(address string, chainType caeruslinks.ChainType) *CreateAddressLinkRequest {
	return &CreateAddressLinkRequest{
		Address:   address,
		ChainType: chainType,
	}
}

type CreateViewProfileLinkRequest struct {
	// Address represents the address of the user for which to create the link
	Address string

	// ChainType represents the chain for which the link should be created
	ChainType caeruslinks.ChainType
}

func NewCreateViewProfileLinkRequest(address string, chainType caeruslinks.ChainType) *CreateViewProfileLinkRequest {
	return &CreateViewProfileLinkRequest{
		Address:   address,
		ChainType: chainType,
	}
}

type CreateSendLinkRequest struct {
	// Address if the address of the user that should receive the funds
	Address string

	// Amount represents the amount of funds to send
	Amount sdk.Coins

	// ChainType represents the chain for which the link should be created
	ChainType caeruslinks.ChainType
}

func NewCreateSendLinkRequest(address string, amount sdk.Coins, chainType caeruslinks.ChainType) *CreateSendLinkRequest {
	return &CreateSendLinkRequest{
		Address:   address,
		Amount:    amount,
		ChainType: chainType,
	}
}

// CreateLinkResponse represents the response returned when a link is created
type CreateLinkResponse struct {
	// DeepLink represents the URL of the generated deep link
	DeepLink string `json:"deep_link"`
}

func NewCreateLinkResponse(deepLink string) *CreateLinkResponse {
	return &CreateLinkResponse{
		DeepLink: deepLink,
	}
}

// GetLinkConfigRequest represents the request sent to get the link configuration
type GetLinkConfigRequest struct {
	Url string
}

func NewGetLinkConfigRequest(url string) *GetLinkConfigRequest {
	return &GetLinkConfigRequest{
		Url: url,
	}
}

// GetLinkConfigResponse represents the response returned when a link configuration is retrieved
type GetLinkConfigResponse struct {
	Url    string
	Config *caerustypes.LinkConfig
}

func NewGetLinkConfigResponse(url string, config *caerustypes.LinkConfig) *GetLinkConfigResponse {
	return &GetLinkConfigResponse{
		Url:    url,
		Config: config,
	}
}
