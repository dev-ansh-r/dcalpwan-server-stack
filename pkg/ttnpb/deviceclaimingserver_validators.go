package ttnpb

import "context"

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *AuthorizeApplicationRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *AuthorizeGatewayRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}
