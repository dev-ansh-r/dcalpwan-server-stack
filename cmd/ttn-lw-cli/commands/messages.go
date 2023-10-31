package commands

var (
	applicationPurgeWarning = []string{
		"This action will permanently delete the application and all related data (API keys, rights, attributes etc.)",
	}
	clientPurgeWarning = []string{
		"This action will permanently delete the client and all related data (rights, attributes etc.)",
	}
	gatewayPurgeWarning = []string{
		"This action will permanently delete the gateway and all related data (API keys, antennas, attributes etc.)",
	}
	organizationPurgeWarning = []string{
		"This action will permanently delete the organization and all related data (API keys, rights, attributes etc.)",
		"It might also cause entities to be orphaned if this organization is the only one that has full rights on the entity.",
	}
	userPurgeWarning = []string{
		"This action will permanently delete the user and all related data (API keys, entity rights, attributes etc.).",
		"It might also cause entities to be orphaned if this user is the only one that has full rights on the entity.",
	}
)
