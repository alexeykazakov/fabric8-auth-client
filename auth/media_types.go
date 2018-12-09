// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "auth": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-auth/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-auth-client
// --pkg=auth
// --version=v1.3.0

package auth

import (
	"github.com/goadesign/goa"
	"net/http"
)

// Role Assignment Array (default view)
//
// Identifier: application/vnd.assign-role-array+json; view=default
type AssignRoleArray struct {
	Data []*AssignRoleData `form:"data" json:"data" xml:"data"`
}

// Validate validates the AssignRoleArray media type instance.
func (mt *AssignRoleArray) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeAssignRoleArray decodes the AssignRoleArray instance encoded in resp body.
func (c *Client) DecodeAssignRoleArray(resp *http.Response) (*AssignRoleArray, error) {
	var decoded AssignRoleArray
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// JWT Token (default view)
//
// Identifier: application/vnd.authtoken+json; view=default
type AuthToken struct {
	Token *TokenData `form:"token" json:"token" xml:"token"`
}

// Validate validates the AuthToken media type instance.
func (mt *AuthToken) Validate() (err error) {
	if mt.Token == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	return
}

// DecodeAuthToken decodes the AuthToken instance encoded in resp body.
func (c *Client) DecodeAuthToken(resp *http.Response) (*AuthToken, error) {
	var decoded AuthToken
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Holds the response to a cluster list request (default view)
//
// Identifier: application/vnd.clusterlist+json; view=default
type ClusterList struct {
	Data []*ClusterData `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// Validate validates the ClusterList media type instance.
func (mt *ClusterList) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeClusterList decodes the ClusterList instance encoded in resp body.
func (c *Client) DecodeClusterList(resp *http.Response) (*ClusterList, error) {
	var decoded ClusterList
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Request payload required to create new invitations (default view)
//
// Identifier: application/vnd.create_invitation_request+json; view=default
type CreateInvitationRequest struct {
	// An array of users invited to become members or to accept a role
	Data []*Invitee `form:"data" json:"data" xml:"data"`
	// links to redirect after accepting invitation sucessfully or in case of error
	Links *RedirectURL `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// Validate validates the CreateInvitationRequest media type instance.
func (mt *CreateInvitationRequest) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	return
}

// DecodeCreateInvitationRequest decodes the CreateInvitationRequest instance encoded in resp body.
func (c *Client) DecodeCreateInvitationRequest(resp *http.Response) (*CreateInvitationRequest, error) {
	var decoded CreateInvitationRequest
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Request payload required to create a new organization (default view)
//
// Identifier: application/vnd.create_organization_request+json; view=default
type CreateOrganizationRequest struct {
	// The name of the new organization
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// DecodeCreateOrganizationRequest decodes the CreateOrganizationRequest instance encoded in resp body.
func (c *Client) DecodeCreateOrganizationRequest(resp *http.Response) (*CreateOrganizationRequest, error) {
	var decoded CreateOrganizationRequest
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Response returned when creating a new organization (default view)
//
// Identifier: application/vnd.create_organization_response+json; view=default
type CreateOrganizationResponse struct {
	// The identifier of the new organization
	OrganizationID *string `form:"organization_id,omitempty" json:"organization_id,omitempty" xml:"organization_id,omitempty"`
}

// DecodeCreateOrganizationResponse decodes the CreateOrganizationResponse instance encoded in resp body.
func (c *Client) DecodeCreateOrganizationResponse(resp *http.Response) (*CreateOrganizationResponse, error) {
	var decoded CreateOrganizationResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Request payload required to create a new team (default view)
//
// Identifier: application/vnd.create_team_request+json; view=default
type CreateTeamRequest struct {
	// The name of the new team
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The identifier of the space in which to create the team
	SpaceID *string `form:"space_id,omitempty" json:"space_id,omitempty" xml:"space_id,omitempty"`
}

// DecodeCreateTeamRequest decodes the CreateTeamRequest instance encoded in resp body.
func (c *Client) DecodeCreateTeamRequest(resp *http.Response) (*CreateTeamRequest, error) {
	var decoded CreateTeamRequest
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Response returned when creating a new team (default view)
//
// Identifier: application/vnd.create_team_response+json; view=default
type CreateTeamResponse struct {
	// The identifier of the new team
	TeamID *string `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
}

// DecodeCreateTeamResponse decodes the CreateTeamResponse instance encoded in resp body.
func (c *Client) DecodeCreateTeamResponse(resp *http.Response) (*CreateTeamResponse, error) {
	var decoded CreateTeamResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// User Create (default view)
//
// Identifier: application/vnd.createuser+json; view=default
type CreateUser struct {
	Data *CreateUserData `form:"data" json:"data" xml:"data"`
}

// Validate validates the CreateUser media type instance.
func (mt *CreateUser) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if mt.Data != nil {
		if err2 := mt.Data.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeCreateUser decodes the CreateUser instance encoded in resp body.
func (c *Client) DecodeCreateUser(resp *http.Response) (*CreateUser, error) {
	var decoded CreateUser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Tokens from external providers such as GitHub or OpenShift (default view)
//
// Identifier: application/vnd.externaltoken+json; view=default
type ExternalToken struct {
	// The token associated with the identity for the specific external provider
	AccessToken string `form:"access_token" json:"access_token" xml:"access_token"`
	// The external provider URL.
	ProviderAPIURL string `form:"provider_api_url" json:"provider_api_url" xml:"provider_api_url"`
	// The scope associated with the token
	Scope string `form:"scope" json:"scope" xml:"scope"`
	// The type of the toke, example : bearer
	TokenType string `form:"token_type" json:"token_type" xml:"token_type"`
	// The username of the identity loaded from the specific external provider
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the ExternalToken media type instance.
func (mt *ExternalToken) Validate() (err error) {
	if mt.AccessToken == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "access_token"))
	}
	if mt.Scope == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "scope"))
	}
	if mt.TokenType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token_type"))
	}
	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if mt.ProviderAPIURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "provider_api_url"))
	}
	return
}

// DecodeExternalToken decodes the ExternalToken instance encoded in resp body.
func (c *Client) DecodeExternalToken(resp *http.Response) (*ExternalToken, error) {
	var decoded ExternalToken
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// This endpoint can be used to obtain a status of the available token from external providers such as GitHub or OpenShift (default view)
//
// Identifier: application/vnd.externaltokenstatus+json; view=default
type ExternalTokenStatus struct {
	// The external provider URL.
	ProviderAPIURL string `form:"provider_api_url" json:"provider_api_url" xml:"provider_api_url"`
	// The username of the identity loaded from the specific external provider.
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the ExternalTokenStatus media type instance.
func (mt *ExternalTokenStatus) Validate() (err error) {
	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if mt.ProviderAPIURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "provider_api_url"))
	}
	return
}

// DecodeExternalTokenStatus decodes the ExternalTokenStatus instance encoded in resp body.
func (c *Client) DecodeExternalTokenStatus(resp *http.Response) (*ExternalTokenStatus, error) {
	var decoded ExternalTokenStatus
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Identity Team Array (default view)
//
// Identifier: application/vnd.identity-team-array+json; view=default
type IdentityTeamArray struct {
	Data []*IdentityTeamData `form:"data" json:"data" xml:"data"`
}

// Validate validates the IdentityTeamArray media type instance.
func (mt *IdentityTeamArray) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeIdentityTeamArray decodes the IdentityTeamArray instance encoded in resp body.
func (c *Client) DecodeIdentityTeamArray(resp *http.Response) (*IdentityTeamArray, error) {
	var decoded IdentityTeamArray
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Assigned Roles in a Protected Resource (default view)
//
// Identifier: application/vnd.identityroles+json; view=default
type Identityroles struct {
	Data []*IdentityRolesData `form:"data" json:"data" xml:"data"`
}

// Validate validates the Identityroles media type instance.
func (mt *Identityroles) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeIdentityroles decodes the Identityroles instance encoded in resp body.
func (c *Client) DecodeIdentityroles(resp *http.Response) (*Identityroles, error) {
	var decoded Identityroles
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// JSONAPIErrors media type (default view)
//
// Identifier: application/vnd.jsonapierrors+json; view=default
type JSONAPIErrors struct {
	Errors []*JSONAPIError `form:"errors" json:"errors" xml:"errors"`
}

// Validate validates the JSONAPIErrors media type instance.
func (mt *JSONAPIErrors) Validate() (err error) {
	if mt.Errors == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "errors"))
	}
	for _, e := range mt.Errors {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeJSONAPIErrors decodes the JSONAPIErrors instance encoded in resp body.
func (c *Client) DecodeJSONAPIErrors(resp *http.Response) (*JSONAPIErrors, error) {
	var decoded JSONAPIErrors
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Oauth 2.0 token payload (default view)
//
// Identifier: application/vnd.oauthtoken+json; view=default
type OauthToken struct {
	// Access token
	AccessToken *string `form:"access_token,omitempty" json:"access_token,omitempty" xml:"access_token,omitempty"`
	// Access token expires in seconds
	ExpiresIn *string `form:"expires_in,omitempty" json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// RefreshToken
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// Token type
	TokenType *string `form:"token_type,omitempty" json:"token_type,omitempty" xml:"token_type,omitempty"`
}

// DecodeOauthToken decodes the OauthToken instance encoded in resp body.
func (c *Client) DecodeOauthToken(resp *http.Response) (*OauthToken, error) {
	var decoded OauthToken
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Indentity Provider Configuration. It list all endpoints supported by Auth Service (default view)
//
// Identifier: application/vnd.openidconfiguration+json; view=default
type OpenIDConfiguration struct {
	// REQUIRED. URL of the OpenID Provider's OAuth 2.0 Authorization Endpoint
	AuthorizationEndpoint *string `form:"authorization_endpoint,omitempty" json:"authorization_endpoint,omitempty" xml:"authorization_endpoint,omitempty"`
	// RECOMMENDED. JSON array containing a list of the Claim Names of the Claims that the OpenID Provider MAY be able to supply values for. Note that for privacy or other reasons, this might not be an exhaustive list.
	ClaimsSupported []string `form:"claims_supported,omitempty" json:"claims_supported,omitempty" xml:"claims_supported,omitempty"`
	// URL of the OpenID Provider's Logout Endpoint
	EndSessionEndpoint *string `form:"end_session_endpoint,omitempty" json:"end_session_endpoint,omitempty" xml:"end_session_endpoint,omitempty"`
	// OPTIONAL. JSON array containing a list of the OAuth 2.0 Grant Type values that this OP supports. Dynamic OpenID Providers MUST support the authorization_code and implicit Grant Type values and MAY support other Grant Types.
	GrantTypesSupported []string `form:"grant_types_supported,omitempty" json:"grant_types_supported,omitempty" xml:"grant_types_supported,omitempty"`
	// REQUIRED. JSON array containing a list of the JWS signing algorithms (alg values) supported by the OpenID Provider for the ID Token to encode the Claims in a JWT. The algorithm RS256 MUST be included.
	IDTokenSigningAlgValuesSupported []string `form:"id_token_signing_alg_values_supported,omitempty" json:"id_token_signing_alg_values_supported,omitempty" xml:"id_token_signing_alg_values_supported,omitempty"`
	// REQUIRED. URL using the https scheme with no query or fragment component that the OpenID Provider asserts as its Issuer Identifier. If Issuer discovery is supported, this value MUST be identical to the issuer value returned by WebFinger. This also MUST be identical to the iss Claim value in ID Tokens issued from this Issuer.
	Issuer *string `form:"issuer,omitempty" json:"issuer,omitempty" xml:"issuer,omitempty"`
	// REQUIRED. URL of the OpenID Provider's JSON Web Key Set [JWK] document. This contains the signing key(s) the Relying Parties uses to validate signatures from the OpenID Provider. The JWK Set MAY also contain the Server's encryption key(s), which are used by Relying Parties to encrypt requests to the Server. When both signing and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage.
	JwksURI *string `form:"jwks_uri,omitempty" json:"jwks_uri,omitempty" xml:"jwks_uri,omitempty"`
	// REQUIRED. JSON array containing a list of the OAuth 2.0 response_type values that this OpenID Provider supports.
	ResponseTypesSupported []string `form:"response_types_supported,omitempty" json:"response_types_supported,omitempty" xml:"response_types_supported,omitempty"`
	// RECOMMENDED. JSON array containing a list of the OAuth 2.0 scope values that this server supports. The server MUST support the `openid` scope value.
	ScopesSupported []string `form:"scopes_supported,omitempty" json:"scopes_supported,omitempty" xml:"scopes_supported,omitempty"`
	// REQUIRED. JSON array containing a list of the Subject Identifier types that this OpenID Provider supports. Valid types include pairwise and public.
	SubjectTypesSupported []string `form:"subject_types_supported,omitempty" json:"subject_types_supported,omitempty" xml:"subject_types_supported,omitempty"`
	// URL of the OpenID Provider's OAuth 2.0 Token Endpoint. This is REQUIRED unless only the Implicit Flow is used.
	TokenEndpoint *string `form:"token_endpoint,omitempty" json:"token_endpoint,omitempty" xml:"token_endpoint,omitempty"`
	// OPTIONAL. JSON array containing a list of Client Authentication methods supported by this Token Endpoint. The options are client_secret_post, client_secret_basic, client_secret_jwt, and private_key_jwt etc.
	TokenEndpointAuthMethodsSupported []string `form:"token_endpoint_auth_methods_supported,omitempty" json:"token_endpoint_auth_methods_supported,omitempty" xml:"token_endpoint_auth_methods_supported,omitempty"`
	// RECOMMENDED. URL of the OpenID Provider's UserInfo Endpoint
	UserinfoEndpoint *string `form:"userinfo_endpoint,omitempty" json:"userinfo_endpoint,omitempty" xml:"userinfo_endpoint,omitempty"`
}

// DecodeOpenIDConfiguration decodes the OpenIDConfiguration instance encoded in resp body.
func (c *Client) DecodeOpenIDConfiguration(resp *http.Response) (*OpenIDConfiguration, error) {
	var decoded OpenIDConfiguration
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Organization Array (default view)
//
// Identifier: application/vnd.organization-array+json; view=default
type OrganizationArray struct {
	Data []*OrganizationData `form:"data" json:"data" xml:"data"`
}

// Validate validates the OrganizationArray media type instance.
func (mt *OrganizationArray) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeOrganizationArray decodes the OrganizationArray instance encoded in resp body.
func (c *Client) DecodeOrganizationArray(resp *http.Response) (*OrganizationArray, error) {
	var decoded OrganizationArray
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Public Keys (default view)
//
// Identifier: application/vnd.publickeys+json; view=default
type PublicKeys struct {
	Keys []interface{} `form:"keys" json:"keys" xml:"keys"`
}

// Validate validates the PublicKeys media type instance.
func (mt *PublicKeys) Validate() (err error) {
	if mt.Keys == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "keys"))
	}
	return
}

// DecodePublicKeys decodes the PublicKeys instance encoded in resp body.
func (c *Client) DecodePublicKeys(resp *http.Response) (*PublicKeys, error) {
	var decoded PublicKeys
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Redirect Location (default view)
//
// Identifier: application/vnd.redirectlocation+json; view=default
type RedirectLocation struct {
	// The location which should be used to redirect browser
	RedirectLocation string `form:"redirect_location" json:"redirect_location" xml:"redirect_location"`
}

// Validate validates the RedirectLocation media type instance.
func (mt *RedirectLocation) Validate() (err error) {
	if mt.RedirectLocation == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "redirect_location"))
	}
	return
}

// DecodeRedirectLocation decodes the RedirectLocation instance encoded in resp body.
func (c *Client) DecodeRedirectLocation(resp *http.Response) (*RedirectLocation, error) {
	var decoded RedirectLocation
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Payload for registering a resource (default view)
//
// Identifier: application/vnd.register_resource+json; view=default
type RegisterResource struct {
	// The parent resource (of the same type) to which this resource belongs
	ParentResourceID *string `form:"parent_resource_id,omitempty" json:"parent_resource_id,omitempty" xml:"parent_resource_id,omitempty"`
	// The identifier for this resource. If left blank, one will be generated
	ResourceID *string `form:"resource_id,omitempty" json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// The type of resource
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the RegisterResource media type instance.
func (mt *RegisterResource) Validate() (err error) {
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	return
}

// DecodeRegisterResource decodes the RegisterResource instance encoded in resp body.
func (c *Client) DecodeRegisterResource(resp *http.Response) (*RegisterResource, error) {
	var decoded RegisterResource
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Response returned when a resource is registered (default view)
//
// Identifier: application/vnd.register_resource_response+json; view=default
type RegisterResourceResponse struct {
	// The identifier for the registered resource
	ResourceID *string `form:"resource_id,omitempty" json:"resource_id,omitempty" xml:"resource_id,omitempty"`
}

// DecodeRegisterResourceResponse decodes the RegisterResourceResponse instance encoded in resp body.
func (c *Client) DecodeRegisterResourceResponse(resp *http.Response) (*RegisterResourceResponse, error) {
	var decoded RegisterResourceResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A Protected Resource (default view)
//
// Identifier: application/vnd.resource+json; view=default
type Resource struct {
	// The parent resource (of the same type) to which this resource belongs
	ParentResourceID *string `form:"parent_resource_id,omitempty" json:"parent_resource_id,omitempty" xml:"parent_resource_id,omitempty"`
	// The identifier for this resource. If left blank, one will be generated
	ResourceID *string `form:"resource_id,omitempty" json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// The valid scopes for this resource
	ResourceScopes []string `form:"resource_scopes,omitempty" json:"resource_scopes,omitempty" xml:"resource_scopes,omitempty"`
	// The type of resource
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// DecodeResource decodes the Resource instance encoded in resp body.
func (c *Client) DecodeResource(resp *http.Response) (*Resource, error) {
	var decoded Resource
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// HasScopes for a user on a resource (default view)
//
// Identifier: application/vnd.resource.scopes+json; view=default
type IdentityResourceScope struct {
	Data *IdentityResourceScopeData `form:"data" json:"data" xml:"data"`
}

// Validate validates the IdentityResourceScope media type instance.
func (mt *IdentityResourceScope) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if mt.Data != nil {
		if err2 := mt.Data.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeIdentityResourceScope decodes the IdentityResourceScope instance encoded in resp body.
func (c *Client) DecodeIdentityResourceScope(resp *http.Response) (*IdentityResourceScope, error) {
	var decoded IdentityResourceScope
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Resource scopes payload (default view)
//
// Identifier: application/vnd.resource_scopes+json; view=default
type ResourceScopes struct {
	// Name of the scope
	ID string `form:"id" json:"id" xml:"id"`
	// Type of resource
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the ResourceScopes media type instance.
func (mt *ResourceScopes) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	return
}

// DecodeResourceScopes decodes the ResourceScopes instance encoded in resp body.
func (c *Client) DecodeResourceScopes(resp *http.Response) (*ResourceScopes, error) {
	var decoded ResourceScopes
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Resource scopes data wrapper (default view)
//
// Identifier: application/vnd.resource_scopes_data+json; view=default
type ResourceScopesData struct {
	// The data wrapper for the response
	Data []*ResourceScopes `form:"data" json:"data" xml:"data"`
}

// Validate validates the ResourceScopesData media type instance.
func (mt *ResourceScopesData) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeResourceScopesData decodes the ResourceScopesData instance encoded in resp body.
func (c *Client) DecodeResourceScopesData(resp *http.Response) (*ResourceScopesData, error) {
	var decoded ResourceScopesData
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Available Roles for a Resource Type (default view)
//
// Identifier: application/vnd.roles+json; view=default
type Roles struct {
	Data []*RolesData `form:"data" json:"data" xml:"data"`
}

// Validate validates the Roles media type instance.
func (mt *Roles) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeRoles decodes the Roles instance encoded in resp body.
func (c *Client) DecodeRoles(resp *http.Response) (*Roles, error) {
	var decoded Roles
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// JWT Token (default view)
//
// Identifier: application/vnd.rpttoken+json; view=default
type RPTToken struct {
	// RPT token
	RptToken *string `form:"rpt_token,omitempty" json:"rpt_token,omitempty" xml:"rpt_token,omitempty"`
}

// DecodeRPTToken decodes the RPTToken instance encoded in resp body.
func (c *Client) DecodeRPTToken(resp *http.Response) (*RPTToken, error) {
	var decoded RPTToken
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Space Resource (default view)
//
// Identifier: application/vnd.spaceresource+json; view=default
type SpaceResource struct {
	Data *SpaceResourceData `form:"data" json:"data" xml:"data"`
}

// Validate validates the SpaceResource media type instance.
func (mt *SpaceResource) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if mt.Data != nil {
		if err2 := mt.Data.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeSpaceResource decodes the SpaceResource instance encoded in resp body.
func (c *Client) DecodeSpaceResource(resp *http.Response) (*SpaceResource, error) {
	var decoded SpaceResource
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The status of the current running instance (default view)
//
// Identifier: application/vnd.status+json; view=default
type Status struct {
	// The time when built
	BuildTime string `form:"buildTime" json:"buildTime" xml:"buildTime"`
	// Commit SHA this build is based on
	Commit string `form:"commit" json:"commit" xml:"commit"`
	// The status of the used configuration. 'OK' or an error message if there is something wrong with the configuration used by service.
	ConfigurationStatus string `form:"configurationStatus" json:"configurationStatus" xml:"configurationStatus"`
	// The status of Database connection. 'OK' or an error message is displayed.
	DatabaseStatus string `form:"databaseStatus" json:"databaseStatus" xml:"databaseStatus"`
	// 'True' if the Developer Mode is enabled
	DevMode *bool `form:"devMode,omitempty" json:"devMode,omitempty" xml:"devMode,omitempty"`
	// The time when started
	StartTime string `form:"startTime" json:"startTime" xml:"startTime"`
}

// Validate validates the Status media type instance.
func (mt *Status) Validate() (err error) {
	if mt.Commit == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "commit"))
	}
	if mt.BuildTime == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "buildTime"))
	}
	if mt.StartTime == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "startTime"))
	}
	if mt.DatabaseStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "databaseStatus"))
	}
	if mt.ConfigurationStatus == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "configurationStatus"))
	}
	return
}

// DecodeStatus decodes the Status instance encoded in resp body.
func (c *Client) DecodeStatus(resp *http.Response) (*Status, error) {
	var decoded Status
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Team Array (default view)
//
// Identifier: application/vnd.team-array+json; view=default
type TeamArray struct {
	Data []*TeamData `form:"data" json:"data" xml:"data"`
}

// Validate validates the TeamArray media type instance.
func (mt *TeamArray) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeTeamArray decodes the TeamArray instance encoded in resp body.
func (c *Client) DecodeTeamArray(resp *http.Response) (*TeamArray, error) {
	var decoded TeamArray
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// User Update (default view)
//
// Identifier: application/vnd.updateuser+json; view=default
type UpdateUser struct {
	Data *UpdateUserData `form:"data" json:"data" xml:"data"`
}

// Validate validates the UpdateUser media type instance.
func (mt *UpdateUser) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if mt.Data != nil {
		if err2 := mt.Data.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeUpdateUser decodes the UpdateUser instance encoded in resp body.
func (c *Client) DecodeUpdateUser(resp *http.Response) (*UpdateUser, error) {
	var decoded UpdateUser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Holds the response of user idenitity IDs for updating list of user IDs (default view)
//
// Identifier: application/vnd.updateuseridlist+json; view=default
type UpdateUserIDList struct {
	Data []*UpdateUserID `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// Validate validates the UpdateUserIDList media type instance.
func (mt *UpdateUserIDList) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeUpdateUserIDList decodes the UpdateUserIDList instance encoded in resp body.
func (c *Client) DecodeUpdateUserIDList(resp *http.Response) (*UpdateUserIDList, error) {
	var decoded UpdateUserIDList
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// User Identity (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	Data *UserData `form:"data" json:"data" xml:"data"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if mt.Data != nil {
		if err2 := mt.Data.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// User Array (default view)
//
// Identifier: application/vnd.user-array+json; view=default
type UserArray struct {
	Data []*UserData `form:"data" json:"data" xml:"data"`
}

// Validate validates the UserArray media type instance.
func (mt *UserArray) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeUserArray decodes the UserArray instance encoded in resp body.
func (c *Client) DecodeUserArray(resp *http.Response) (*UserArray, error) {
	var decoded UserArray
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// User Information (default view)
//
// Identifier: application/vnd.userinfo+json; view=default
type UserInfo struct {
	// email of the user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// last name, can be achieved from fullName from the user table
	FamilyName *string `form:"family_name,omitempty" json:"family_name,omitempty" xml:"family_name,omitempty"`
	// first name, can be achieved from fullName rrom the user table
	GivenName *string `form:"given_name,omitempty" json:"given_name,omitempty" xml:"given_name,omitempty"`
	// username, each identity has a username
	PreferredUsername *string `form:"preferred_username,omitempty" json:"preferred_username,omitempty" xml:"preferred_username,omitempty"`
	// subject (identity is subject)
	Sub *string `form:"sub,omitempty" json:"sub,omitempty" xml:"sub,omitempty"`
}

// DecodeUserInfo decodes the UserInfo instance encoded in resp body.
func (c *Client) DecodeUserInfo(resp *http.Response) (*UserInfo, error) {
	var decoded UserInfo
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Holds the paginated response to a user list request (default view)
//
// Identifier: application/vnd.userlist+json; view=default
type UserList struct {
	Data []*UserData `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
	Links    *PagingLinks  `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta     *UserListMeta `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// Validate validates the UserList media type instance.
func (mt *UserList) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeUserList decodes the UserList instance encoded in resp body.
func (c *Client) DecodeUserList(resp *http.Response) (*UserList, error) {
	var decoded UserList
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Holds the paginated response to a user spaces request (default view)
//
// Identifier: application/vnd.userresourceslist+json; view=default
type UserResourcesList struct {
	Data []*UserResourceData `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
	Links    *PagingLinks  `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	Meta     *UserListMeta `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// Validate validates the UserResourcesList media type instance.
func (mt *UserResourcesList) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeUserResourcesList decodes the UserResourcesList instance encoded in resp body.
func (c *Client) DecodeUserResourcesList(resp *http.Response) (*UserResourcesList, error) {
	var decoded UserResourcesList
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
