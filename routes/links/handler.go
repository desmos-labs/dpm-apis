package links

import (
	"net/http"

	caeruslinks "github.com/desmos-labs/caerus/routes/links"

	"github.com/desmos-labs/dpm-apis/utils"
)

type Handler struct {
	caerus CaerusClient
}

func NewHandler(caerusClient CaerusClient) *Handler {
	return &Handler{
		caerus: caerusClient,
	}
}

// HandleCreateAddressLinkRequest handles the given CreateAddressLinkRequest returning the link address or an error
func (h *Handler) HandleCreateAddressLinkRequest(req *CreateAddressLinkRequest) (*CreateLinkResponse, error) {
	res, err := h.caerus.CreateAddressLink(&caeruslinks.CreateAddressLinkRequest{
		Address: req.Address,
		Chain:   req.ChainType,
	})
	if err != nil {
		return nil, err
	}

	return NewCreateLinkResponse(res.Url), nil
}

// HandleCreateViewProfileLinkRequest handles the given CreateViewProfileLinkRequest returning the link address or an error
func (h *Handler) HandleCreateViewProfileLinkRequest(req *CreateViewProfileLinkRequest) (*CreateLinkResponse, error) {
	res, err := h.caerus.CreateViewProfileLink(&caeruslinks.CreateViewProfileLinkRequest{
		Address: req.Address,
		Chain:   req.ChainType,
	})
	if err != nil {
		return nil, err
	}

	return NewCreateLinkResponse(res.Url), nil
}

// HandleCreateSendLinkRequest handles the given CreateSendLinkRequest returning the link address or an error
func (h *Handler) HandleCreateSendLinkRequest(req *CreateSendLinkRequest) (*CreateLinkResponse, error) {
	res, err := h.caerus.CreateSendLink(&caeruslinks.CreateSendLinkRequest{
		Address: req.Address,
		Amount:  req.Amount,
		Chain:   req.ChainType,
	})
	if err != nil {
		return nil, err
	}

	return NewCreateLinkResponse(res.Url), nil
}

// HandleGetLinkConfigRequest handles the given GetLinkConfigRequest returning the link config or an error
func (h *Handler) HandleGetLinkConfigRequest(url string) (*GetLinkConfigResponse, error) {
	res, err := h.caerus.GetLinkConfig(url)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, utils.WrapErr(http.StatusNotFound, "link not found")
	}

	return NewGetLinkConfigResponse(url, res), nil
}
