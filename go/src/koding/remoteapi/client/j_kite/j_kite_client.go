package j_kite

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new j kite API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for j kite API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostRemoteAPIJKiteCreate post remote API j kite create API
*/
func (a *Client) PostRemoteAPIJKiteCreate(params *PostRemoteAPIJKiteCreateParams) (*PostRemoteAPIJKiteCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJKiteCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJKiteCreate",
		Method:             "POST",
		PathPattern:        "/remote.api/JKite.create",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJKiteCreateReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJKiteCreateOK), nil

}

/*
PostRemoteAPIJKiteCreatePlanID post remote API j kite create plan ID API
*/
func (a *Client) PostRemoteAPIJKiteCreatePlanID(params *PostRemoteAPIJKiteCreatePlanIDParams) (*PostRemoteAPIJKiteCreatePlanIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJKiteCreatePlanIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJKiteCreatePlanID",
		Method:             "POST",
		PathPattern:        "/remote.api/JKite.createPlan/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJKiteCreatePlanIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJKiteCreatePlanIDOK), nil

}

/*
PostRemoteAPIJKiteDeleteKiteID post remote API j kite delete kite ID API
*/
func (a *Client) PostRemoteAPIJKiteDeleteKiteID(params *PostRemoteAPIJKiteDeleteKiteIDParams) (*PostRemoteAPIJKiteDeleteKiteIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJKiteDeleteKiteIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJKiteDeleteKiteID",
		Method:             "POST",
		PathPattern:        "/remote.api/JKite.deleteKite/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJKiteDeleteKiteIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJKiteDeleteKiteIDOK), nil

}

/*
PostRemoteAPIJKiteDeletePlanID post remote API j kite delete plan ID API
*/
func (a *Client) PostRemoteAPIJKiteDeletePlanID(params *PostRemoteAPIJKiteDeletePlanIDParams) (*PostRemoteAPIJKiteDeletePlanIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJKiteDeletePlanIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJKiteDeletePlanID",
		Method:             "POST",
		PathPattern:        "/remote.api/JKite.deletePlan/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJKiteDeletePlanIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJKiteDeletePlanIDOK), nil

}

/*
PostRemoteAPIJKiteFetchPlansID post remote API j kite fetch plans ID API
*/
func (a *Client) PostRemoteAPIJKiteFetchPlansID(params *PostRemoteAPIJKiteFetchPlansIDParams) (*PostRemoteAPIJKiteFetchPlansIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJKiteFetchPlansIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJKiteFetchPlansID",
		Method:             "POST",
		PathPattern:        "/remote.api/JKite.fetchPlans/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJKiteFetchPlansIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJKiteFetchPlansIDOK), nil

}

/*
PostRemoteAPIJKiteList post remote API j kite list API
*/
func (a *Client) PostRemoteAPIJKiteList(params *PostRemoteAPIJKiteListParams) (*PostRemoteAPIJKiteListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJKiteListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJKiteList",
		Method:             "POST",
		PathPattern:        "/remote.api/JKite.list",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJKiteListReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJKiteListOK), nil

}

/*
PostRemoteAPIJKiteModifyID post remote API j kite modify ID API
*/
func (a *Client) PostRemoteAPIJKiteModifyID(params *PostRemoteAPIJKiteModifyIDParams) (*PostRemoteAPIJKiteModifyIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJKiteModifyIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJKiteModifyID",
		Method:             "POST",
		PathPattern:        "/remote.api/JKite.modify/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJKiteModifyIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJKiteModifyIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
