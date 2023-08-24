package caerus

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	caerusauth "github.com/desmos-labs/caerus/authentication"
	caeruslinks "github.com/desmos-labs/caerus/routes/links"
	"github.com/desmos-labs/caerus/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addressPrefix = regexp.MustCompile(`^https?://`)
)

// Client represents a client that can be used to communicate with the Caerus server
type Client struct {
	apiKey       string
	linksService caeruslinks.LinksServiceClient
}

// NewClient returns a new Client instance with the given gRPC connection
func NewClient(apiKey string, caerusGrpcConn *grpc.ClientConn) *Client {
	return &Client{
		apiKey:       apiKey,
		linksService: caeruslinks.NewLinksServiceClient(caerusGrpcConn),
	}
}

// NewClientFromEnvVariables returns a new Client instance from the environment variables
func NewClientFromEnvVariables() (*Client, error) {
	caerusGrpcAddress := utils.GetEnvOr(EnvCaerusGRPCAddress, "")
	if caerusGrpcAddress == "" {
		return nil, fmt.Errorf("missing %s", EnvCaerusGRPCAddress)
	}

	apiKey := utils.GetEnvOr(EnvCaerusAPIKey, "")
	if apiKey == "" {
		return nil, fmt.Errorf("missing %s", EnvCaerusAPIKey)
	}

	// Build the transport credentials based on the HTTP protocol specified inside the URL
	transportCredential := insecure.NewCredentials()
	if strings.HasPrefix(caerusGrpcAddress, "https://") {
		transportCredential = credentials.NewClientTLSFromCert(nil, "")
	}

	// Trim the https?:// prefix
	caerusGrpcAddress = addressPrefix.ReplaceAllString(caerusGrpcAddress, "")

	// Build the connection
	grpcConn, err := grpc.Dial(caerusGrpcAddress, grpc.WithTransportCredentials(transportCredential))
	if err != nil {
		return nil, err
	}

	return NewClient(apiKey, grpcConn), nil
}

func (client *Client) getAuthenticatedContext() context.Context {
	return caerusauth.SetupContextWithAuthorization(context.Background(), client.apiKey)
}

// --------------------------------------------------------------------------------------------------------------------

// CreateAddressLink allows to generate a new deep link that allows to open
// the given address on the given chain and perform the action decided by the
// user
func (client *Client) CreateAddressLink(request *caeruslinks.CreateAddressLinkRequest) (*caeruslinks.CreateLinkResponse, error) {
	return client.linksService.CreateAddressLink(client.getAuthenticatedContext(), request)
}

// CreateViewProfileLink allows to generate a new deep link that allows to
// view the profile of the given user
func (client *Client) CreateViewProfileLink(request *caeruslinks.CreateViewProfileLinkRequest) (*caeruslinks.CreateLinkResponse, error) {
	return client.linksService.CreateViewProfileLink(client.getAuthenticatedContext(), request)
}

// CreateSendLink allows to generate a new deep link that allows to send
// tokens to the given address
func (client *Client) CreateSendLink(request *caeruslinks.CreateSendLinkRequest) (*caeruslinks.CreateLinkResponse, error) {
	return client.linksService.CreateSendLink(client.getAuthenticatedContext(), request)
}

// CreateLink allows to generated a new deep link based on the given
// configuration
func (client *Client) CreateLink(request *caeruslinks.CreateLinkRequest) (*caeruslinks.CreateLinkResponse, error) {
	return client.linksService.CreateLink(client.getAuthenticatedContext(), request)
}
