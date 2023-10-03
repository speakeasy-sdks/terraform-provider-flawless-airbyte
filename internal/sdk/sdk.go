// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"airbyte/internal/sdk/pkg/utils"
	"fmt"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"http://localhost:8000/api",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Airbyte - Airbyte Configuration API: Airbyte Configuration API
// [https://airbyte.io](https://airbyte.io).
//
// The Configuration API is an internal Airbyte API that is designed for communications between different Airbyte components.
// * Its main purpose is to enable the Airbyte Engineering team to configure the internal state of [Airbyte Cloud](https://airbyte.com/airbyte-cloud)
// * It is also sometimes used by OSS users to configure their own Self-Hosted Airbyte deployment (internal state, etc)
//
// WARNING
// * Airbyte does NOT have active commitments to support this API long-term.
// * OSS users can utilize the Configuration API, but at their own risk.
// * This API is utilized internally by the Airbyte Engineering team and may be modified in the future if the need arises.
// * Modifications by the Airbyte Engineering team could create breaking changes and OSS users would need to update their code to catch up to any backwards incompatible changes in the API.
//
// This API is a collection of HTTP RPC-style methods. While it is not a REST API, those familiar with REST should find the conventions of this API recognizable.
//
// Here are some conventions that this API follows:
// * All endpoints are http POST methods.
// * All endpoints accept data via `application/json` request bodies. The API does not accept any data via query params.
// * The naming convention for endpoints is: localhost:8000/api/{VERSION}/{METHOD_FAMILY}/{METHOD_NAME} e.g. `localhost:8000/api/v1/connections/create`.
// * For all `update` methods, the whole object must be passed in, even the fields that did not change.
//
// Authentication (OSS):
// * When authenticating to the Configuration API, you must use Basic Authentication by setting the Authentication Header to Basic and base64 encoding the username and password (which are `airbyte` and `password` by default - so base64 encoding `airbyte:password` results in `YWlyYnl0ZTpwYXNzd29yZA==`). So the full header reads `'Authorization': "Basic YWlyYnl0ZTpwYXNzd29yZA=="`
//
// https://airbyte.io - Find out more about Airbyte
type Airbyte struct {
	// Attempt - Interactions with attempt related resources.
	Attempt *attempt
	// Connection - Connection between sources and destinations.
	Connection                   *connection
	ConnectorBuilderProject      *connectorBuilderProject
	DeclarativeSourceDefinitions *declarativeSourceDefinitions
	// Destination - Destination related resources.
	Destination *destination
	// DestinationDefinition - DestinationDefinition related resources.
	DestinationDefinition *destinationDefinition
	// DestinationDefinitionSpecification - DestinationDefinitionSpecification related resources.
	DestinationDefinitionSpecification *destinationDefinitionSpecification
	// DestinationOauth - Source OAuth related resources to delegate access from user.
	DestinationOauth *destinationOauth
	// Health - Healthchecks
	Health        *health
	Internal      *internal
	Jobs          *jobs
	Logs          *logs
	Notifications *notifications
	Openapi       *openapi
	Operation     *operation
	Scheduler     *scheduler
	// Source - Source related resources.
	Source *source
	// SourceDefinition - SourceDefinition related resources.
	SourceDefinition *sourceDefinition
	// SourceDefinitionSpecification - SourceDefinition specification related resources.
	SourceDefinitionSpecification *sourceDefinitionSpecification
	// SourceOauth - Source OAuth related resources to delegate access from user.
	SourceOauth *sourceOauth
	// State - Interactions with state related resources.
	State          *state
	StreamStatuses *streamStatuses
	Streams        *streams
	// WebBackend - Endpoints for the Airbyte web application. Those APIs should not be called outside the web application implementation and are not
	// guaranteeing any backwards compatibility.
	//
	WebBackend *webBackend
	// Workspace - Workspace related resources.
	Workspace *workspace

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Airbyte)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Airbyte) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Airbyte) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Airbyte) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Airbyte) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Airbyte) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Airbyte {
	sdk := &Airbyte{
		sdkConfiguration: sdkConfiguration{
			Language:          "terraform",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "0.5.0",
			GenVersion:        "2.144.7",
			UserAgent:         "speakeasy-sdk/terraform 0.5.0 2.144.7 1.0.0 airbyte",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Attempt = newAttempt(sdk.sdkConfiguration)

	sdk.Connection = newConnection(sdk.sdkConfiguration)

	sdk.ConnectorBuilderProject = newConnectorBuilderProject(sdk.sdkConfiguration)

	sdk.DeclarativeSourceDefinitions = newDeclarativeSourceDefinitions(sdk.sdkConfiguration)

	sdk.Destination = newDestination(sdk.sdkConfiguration)

	sdk.DestinationDefinition = newDestinationDefinition(sdk.sdkConfiguration)

	sdk.DestinationDefinitionSpecification = newDestinationDefinitionSpecification(sdk.sdkConfiguration)

	sdk.DestinationOauth = newDestinationOauth(sdk.sdkConfiguration)

	sdk.Health = newHealth(sdk.sdkConfiguration)

	sdk.Internal = newInternal(sdk.sdkConfiguration)

	sdk.Jobs = newJobs(sdk.sdkConfiguration)

	sdk.Logs = newLogs(sdk.sdkConfiguration)

	sdk.Notifications = newNotifications(sdk.sdkConfiguration)

	sdk.Openapi = newOpenapi(sdk.sdkConfiguration)

	sdk.Operation = newOperation(sdk.sdkConfiguration)

	sdk.Scheduler = newScheduler(sdk.sdkConfiguration)

	sdk.Source = newSource(sdk.sdkConfiguration)

	sdk.SourceDefinition = newSourceDefinition(sdk.sdkConfiguration)

	sdk.SourceDefinitionSpecification = newSourceDefinitionSpecification(sdk.sdkConfiguration)

	sdk.SourceOauth = newSourceOauth(sdk.sdkConfiguration)

	sdk.State = newState(sdk.sdkConfiguration)

	sdk.StreamStatuses = newStreamStatuses(sdk.sdkConfiguration)

	sdk.Streams = newStreams(sdk.sdkConfiguration)

	sdk.WebBackend = newWebBackend(sdk.sdkConfiguration)

	sdk.Workspace = newWorkspace(sdk.sdkConfiguration)

	return sdk
}
