package ttnpb

import "context"

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *ListOAuthClientAuthorizationsRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}
