// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.1.0
// - protoc              v4.22.2
// source: ttn/lorawan/v3/organization.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	golang "github.com/TheThingsIndustries/protoc-gen-go-flags/golang"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForOrganization adds flags to select fields in Organization.
func AddSelectFlagsForOrganization(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("deleted-at", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("name", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("name", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("description", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("description", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("attributes", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("attributes", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("contact-info", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("contact-info", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("administrative-contact", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("administrative-contact", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("administrative-contact", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("technical-contact", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("technical-contact", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("technical-contact", prefix), hidden)
}

// SelectFromFlags outputs the fieldmask paths forOrganization message from select flags.
func PathsFromSelectFlagsForOrganization(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted_at", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("deleted_at", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("name", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("name", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("description", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("description", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("attributes", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("attributes", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("contact_info", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("contact_info", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("administrative_contact", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("administrative_contact", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("administrative_contact", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("technical_contact", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("technical_contact", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("technical_contact", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	return paths, nil
}

// AddSetFlagsForOrganization adds flags to select fields in Organization.
func AddSetFlagsForOrganization(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForOrganizationIdentifiers(flags, flagsplugin.Prefix("ids", prefix), true)
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes", prefix), "", flagsplugin.WithHidden(hidden)))
	// FIXME: Skipping ContactInfo because repeated messages are currently not supported.
	AddSetFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("administrative-contact", prefix), hidden)
	AddSetFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("technical-contact", prefix), hidden)
}

// SetFromFlags sets the Organization message from flags.
func (m *Organization) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("ids", prefix)); changed {
		if m.Ids == nil {
			m.Ids = &OrganizationIdentifiers{}
		}
		if setPaths, err := m.Ids.SetFromFlags(flags, flagsplugin.Prefix("ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Name = val
		paths = append(paths, flagsplugin.Prefix("name", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Description = val
		paths = append(paths, flagsplugin.Prefix("description", prefix))
	}
	if val, changed, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Attributes = val
		paths = append(paths, flagsplugin.Prefix("attributes", prefix))
	}
	// FIXME: Skipping ContactInfo because it does not seem to implement AddSetFlags.
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("administrative_contact", prefix)); changed {
		if m.AdministrativeContact == nil {
			m.AdministrativeContact = &OrganizationOrUserIdentifiers{}
		}
		if setPaths, err := m.AdministrativeContact.SetFromFlags(flags, flagsplugin.Prefix("administrative_contact", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("technical_contact", prefix)); changed {
		if m.TechnicalContact == nil {
			m.TechnicalContact = &OrganizationOrUserIdentifiers{}
		}
		if setPaths, err := m.TechnicalContact.SetFromFlags(flags, flagsplugin.Prefix("technical_contact", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	return paths, nil
}

// AddSetFlagsForListOrganizationsRequest adds flags to select fields in ListOrganizationsRequest.
func AddSetFlagsForListOrganizationsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("collaborator", prefix), true)
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the ListOrganizationsRequest message from flags.
func (m *ListOrganizationsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("collaborator", prefix)); changed {
		if m.Collaborator == nil {
			m.Collaborator = &OrganizationOrUserIdentifiers{}
		}
		if setPaths, err := m.Collaborator.SetFromFlags(flags, flagsplugin.Prefix("collaborator", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field_mask", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FieldMask = golang.SetFieldMask(val)
		paths = append(paths, flagsplugin.Prefix("field_mask", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Deleted = val
		paths = append(paths, flagsplugin.Prefix("deleted", prefix))
	}
	return paths, nil
}

// AddSetFlagsForListOrganizationAPIKeysRequest adds flags to select fields in ListOrganizationAPIKeysRequest.
func AddSetFlagsForListOrganizationAPIKeysRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForOrganizationIdentifiers(flags, flagsplugin.Prefix("organization-ids", prefix), hidden)
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the ListOrganizationAPIKeysRequest message from flags.
func (m *ListOrganizationAPIKeysRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("organization_ids", prefix)); changed {
		if m.OrganizationIds == nil {
			m.OrganizationIds = &OrganizationIdentifiers{}
		}
		if setPaths, err := m.OrganizationIds.SetFromFlags(flags, flagsplugin.Prefix("organization_ids", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Order = val
		paths = append(paths, flagsplugin.Prefix("order", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Limit = val
		paths = append(paths, flagsplugin.Prefix("limit", prefix))
	}
	if val, changed, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Page = val
		paths = append(paths, flagsplugin.Prefix("page", prefix))
	}
	return paths, nil
}
