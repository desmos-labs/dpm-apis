package links

import (
	caeruslinks "github.com/desmos-labs/caerus/routes/links"
)

type CaerusClient interface {
	CreateAddressLink(request *caeruslinks.CreateAddressLinkRequest) (*caeruslinks.CreateLinkResponse, error)
	CreateViewProfileLink(request *caeruslinks.CreateViewProfileLinkRequest) (*caeruslinks.CreateLinkResponse, error)
	CreateSendLink(request *caeruslinks.CreateSendLinkRequest) (*caeruslinks.CreateLinkResponse, error)
	CreateLink(request *caeruslinks.CreateLinkRequest) (*caeruslinks.CreateLinkResponse, error)
}
