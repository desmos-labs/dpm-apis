package links

import (
	"net/http"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	caeruslinks "github.com/desmos-labs/caerus/routes/links"
	"github.com/gin-gonic/gin"

	"github.com/desmos-labs/dpm-apis/routes"
	"github.com/desmos-labs/dpm-apis/utils"
)

const (
	ChainTypeKey = "chain_type"
	AmountKey    = "amount"
)

func RegisterWithContext(ctx routes.Context) {
	Register(ctx.Router, NewHandler(ctx.Caerus))
}

// Register registers all the routes that allow to perform links-related operations
func Register(router *gin.Engine, handler *Handler) {
	router.
		GET("/deep-links/config", func(context *gin.Context) {
			deepLinkURL, exists := context.GetQuery("url")
			if !exists {
				utils.HandleError(context, utils.WrapErr(http.StatusBadRequest, "missing url param"))
				return
			}

			res, err := handler.HandleGetLinkConfigRequest(deepLinkURL)
			if err != nil {
				utils.HandleError(context, err)
				return
			}

			context.JSON(http.StatusOK, res)
		})

	router.Group("/deep-links/:address").
		GET("", func(c *gin.Context) {
			// Build the request
			address, err := parseAddress(c)
			if err != nil {
				utils.HandleError(c, err)
				return
			}
			chainType, err := parseChainType(c)
			if err != nil {
				utils.HandleError(c, err)
				return
			}
			req := NewCreateAddressLinkRequest(address, chainType)

			// Handle the request
			res, err := handler.HandleCreateAddressLinkRequest(req)
			if err != nil {
				utils.HandleError(c, err)
				return
			}

			// Return the response
			c.JSON(http.StatusOK, res)
		}).
		GET("/view-profile", func(c *gin.Context) {
			// Build the request
			address, err := parseAddress(c)
			if err != nil {
				utils.HandleError(c, err)
				return
			}
			chainType, err := parseChainType(c)
			if err != nil {
				utils.HandleError(c, err)
				return
			}
			req := NewCreateViewProfileLinkRequest(address, chainType)

			// Handle the request
			res, err := handler.HandleCreateViewProfileLinkRequest(req)
			if err != nil {
				utils.HandleError(c, err)
				return
			}

			// Return the response
			c.JSON(http.StatusOK, res)
		}).
		GET("/send", func(c *gin.Context) {
			// Build the request
			address, err := parseAddress(c)
			if err != nil {
				utils.HandleError(c, err)
				return
			}
			chainType, err := parseChainType(c)
			if err != nil {
				utils.HandleError(c, err)
				return
			}
			amount, err := parseAmount(c)
			if err != nil {
				utils.HandleError(c, err)
				return
			}
			req := NewCreateSendLinkRequest(address, amount, chainType)

			// Handle the request
			res, err := handler.HandleCreateSendLinkRequest(req)
			if err != nil {
				utils.HandleError(c, err)
				return
			}

			// Return the response
			c.JSON(http.StatusOK, res)
		})
}

// parseAddress returns the address that has been specified inside the given context.
// It expects the address to be specified using the address query param in the form of a
// string (es. "desmos1...").
// If the specified address is not valid, it returns an error
func parseAddress(context *gin.Context) (string, error) {
	address := context.Param("address")
	if address == "" {
		return "", utils.WrapErr(http.StatusBadRequest, "invalid address")
	}

	_, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return "", utils.WrapErr(http.StatusBadRequest, "invalid address")
	}

	return address, nil
}

// parseChainType returns the chain type that has been specified inside the given context.
// It expects the chain type to be specified using the ChainTypeKey in the form of a
// string (either "mainnet" or "testnet").
func parseChainType(context *gin.Context) (caeruslinks.ChainType, error) {
	chainType, exists := context.GetQuery(ChainTypeKey)
	if !exists {
		return caeruslinks.ChainType_UNDEFINED, utils.WrapErr(http.StatusBadRequest, "invalid chain type")
	}

	chainTypeValue, ok := caeruslinks.ChainType_value[strings.ToUpper(chainType)]
	if !ok {
		return caeruslinks.ChainType_UNDEFINED, utils.WrapErr(http.StatusBadRequest, "invalid chain type")
	}

	return caeruslinks.ChainType(chainTypeValue), nil
}

// parseAmount returns the amount that has been specified inside the given context.
// It expects the amount to be specified using the AmountKey in the form of a
// string (e.g. "1000udaric").
// If the specified amount is not valid, it returns an error
func parseAmount(context *gin.Context) (sdk.Coins, error) {
	amountValue, exists := context.GetQuery(AmountKey)
	if !exists {
		return sdk.NewCoins(), nil
	}

	amount, err := sdk.ParseCoinsNormalized(amountValue)
	if err != nil {
		return sdk.NewCoins(), utils.WrapErr(http.StatusBadRequest, "invalid amount")
	}

	return amount, nil
}
