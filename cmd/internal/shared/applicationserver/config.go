package shared

import (
	"fmt"
	"time"

	"go.thethings.network/lorawan-stack/v3/cmd/internal/shared"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/packages"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/web"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
)

// DefaultWebhookTemplatesConfig is the default configuration for the Webhook templates.
var DefaultWebhookTemplatesConfig = web.TemplatesConfig{
	Directory: "/srv/ttn-lorawan/lorawan-webhook-templates",
	URL:       "https://raw.githubusercontent.com/TheThingsNetwork/lorawan-webhook-templates/master",
}

// DefaultApplicationServerConfig is the default configuration for the Application Server.
var DefaultApplicationServerConfig = applicationserver.Config{
	MQTT: config.MQTT{
		Listen:           ":1883",
		ListenTLS:        ":8883",
		PublicAddress:    fmt.Sprintf("%s:1883", shared.DefaultPublicHost),
		PublicTLSAddress: fmt.Sprintf("%s:8883", shared.DefaultPublicHost),
	},
	Webhooks: applicationserver.WebhooksConfig{
		Templates: DefaultWebhookTemplatesConfig,
		Target:    "direct",
		Timeout:   5 * time.Second,
		QueueSize: 1024,
		Workers:   1024,
		Downlinks: web.DownlinksConfig{PublicAddress: shared.DefaultPublicURL + "/api/v3"},
	},
	EndDeviceMetadataStorage: applicationserver.EndDeviceMetadataStorageConfig{
		Location: applicationserver.EndDeviceLocationStorageConfig{
			Timeout: 5 * time.Second,
			Cache: applicationserver.EndDeviceLocationStorageCacheConfig{
				Enable:             true,
				MinRefreshInterval: 15 * time.Minute,
				MaxRefreshInterval: 4 * time.Hour,
				TTL:                14 * 24 * time.Hour,
			},
		},
	},
	Distribution: applicationserver.DistributionConfig{
		Timeout: time.Minute,
		Local: applicationserver.LocalDistributorConfig{
			Broadcast: applicationserver.DistributorConfig{
				SubscriptionBlocks:    true,
				SubscriptionQueueSize: -1,
			},
			Individual: applicationserver.DistributorConfig{
				SubscriptionBlocks:    false,
				SubscriptionQueueSize: io.DefaultBufferSize,
			},
		},
		Global: applicationserver.GlobalDistributorConfig{
			Individual: applicationserver.DistributorConfig{
				SubscriptionBlocks:    false,
				SubscriptionQueueSize: io.DefaultBufferSize,
			},
		},
	},
	PubSub: applicationserver.PubSubConfig{
		Providers: map[string]string{
			"mqtt": "enabled",
			"nats": "enabled",
		},
	},
	Packages: applicationserver.ApplicationPackagesConfig{
		Config: packages.Config{
			Workers: 1024,
			Timeout: 10 * time.Second,
		},
	},
	Formatters: applicationserver.FormattersConfig{
		MaxParameterLength: 40960,
	},
	DeviceLastSeen: applicationserver.LastSeenConfig{
		BatchSize:     1000,
		FlushInterval: 10 * time.Second,
	},
	Downlinks: applicationserver.DownlinksConfig{
		ConfirmationConfig: applicationserver.ConfirmationConfig{
			DefaultRetryAttempts: 8,
			MaxRetryAttempts:     32,
		},
	},
}
