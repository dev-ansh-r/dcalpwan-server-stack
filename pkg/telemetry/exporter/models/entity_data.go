package models

// AccountsCount telemetry data about the amount of accounts, total and the amount on each type.
type AccountsCount struct {
	Users         UsersCount         `json:"users"`
	Organizations OrganizationsCount `json:"organizations"`
}

// UsersCount telemetry data about the amount of users, total and the amount on each type.
type UsersCount struct {
	Standard uint64 `json:"standard"`
	Admin    uint64 `json:"admin"`
}

// OrganizationsCount telemetry data about the total amount of orgainizations.
type OrganizationsCount struct {
	Total uint64 `json:"total"`
}

// ApplicationsCount telemetry data about the amount of applications.
type ApplicationsCount struct {
	Total uint64 `json:"total"`
}

// ActivateEndDevicesCount contains the amount of devices that are active respecting the described time frame.
type ActivateEndDevicesCount struct {
	Total     uint64 `json:"total"`
	LastDay   uint64 `json:"last_day"`
	LastWeek  uint64 `json:"last_week"`
	LastMonth uint64 `json:"last_month"`
}

// EndDevicesCount contains telemetry data regarding the amount of end devices and its different types.
type EndDevicesCount struct {
	Total              uint64                  `json:"total"`
	ActivateEndDevices ActivateEndDevicesCount `json:"activate_end_devices"`
}

// GatewaysCount contains telemetry data regarding the amount of gateways and some extra insights.
type GatewaysCount struct {
	Total                     uint64            `json:"total"`
	GatewaysByFrequencyPlanID map[string]uint64 `json:"gateways_by_frequency_plan_id"`
}

// EntitiesCount contains telemetry data regarding the amount of each entity and its different types.
type EntitiesCount struct {
	Gateways     GatewaysCount     `json:"gateways"`
	EndDevices   EndDevicesCount   `json:"end_devices"`
	Applications ApplicationsCount `json:"applications"`
	Accounts     AccountsCount     `json:"accounts"`
}
