package links

import (
	caeruslinks "github.com/desmos-labs/caerus/routes/links"
	caerustypes "github.com/desmos-labs/caerus/types"
)

type CaerusClient interface {
	CreateAddressLink(request *caeruslinks.CreateAddressLinkRequest) (*caeruslinks.CreateLinkResponse, error)
	CreateViewProfileLink(request *caeruslinks.CreateViewProfileLinkRequest) (*caeruslinks.CreateLinkResponse, error)
	CreateSendLink(request *caeruslinks.CreateSendLinkRequest) (*caeruslinks.CreateLinkResponse, error)
	CreateLink(config *caerustypes.LinkConfig) (*caeruslinks.CreateLinkResponse, error)
	GetLinkConfig(url string) (*caerustypes.LinkConfig, error)
}
