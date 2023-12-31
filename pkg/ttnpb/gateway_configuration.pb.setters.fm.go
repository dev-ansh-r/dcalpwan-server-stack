// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *GetGatewayConfigurationRequest) SetFields(src *GetGatewayConfigurationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_ids":
			if len(subs) > 0 {
				var newDst, newSrc *GatewayIdentifiers
				if (src == nil || src.GatewayIds == nil) && dst.GatewayIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.GatewayIds
				}
				if dst.GatewayIds != nil {
					newDst = dst.GatewayIds
				} else {
					newDst = &GatewayIdentifiers{}
					dst.GatewayIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.GatewayIds = src.GatewayIds
				} else {
					dst.GatewayIds = nil
				}
			}
		case "format":
			if len(subs) > 0 {
				return fmt.Errorf("'format' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Format = src.Format
			} else {
				var zero string
				dst.Format = zero
			}
		case "type":
			if len(subs) > 0 {
				return fmt.Errorf("'type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Type = src.Type
			} else {
				var zero string
				dst.Type = zero
			}
		case "filename":
			if len(subs) > 0 {
				return fmt.Errorf("'filename' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Filename = src.Filename
			} else {
				var zero string
				dst.Filename = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetGatewayConfigurationResponse) SetFields(src *GetGatewayConfigurationResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "contents":
			if len(subs) > 0 {
				return fmt.Errorf("'contents' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Contents = src.Contents
			} else {
				dst.Contents = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
