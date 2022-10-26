package redfishreceiver

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/valcharry/kubecon2022/receiver/redfishreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

const (
	typeStr   = "redfish"
	stability = component.StabilityLevelInDevelopment
)

var errConfigNotRedfish = errors.New("config was not a redfish receiver config")

func createDefaultConfig() config.Receiver {

	// ok, let's create a configuration that has a 60 second collection interval, full "scrape" (not subscription) against the endpoint URL
	return &Config{
		ScraperControllerSettings: scraperhelper.ScraperControllerSettings{
			ReceiverSettings:   config.NewReceiverSettings(config.NewComponentID(typeStr)),
			CollectionInterval: 60 * time.Second,
		},
		Mode: "pull",
		HTTPClientSettings: confighttp.HTTPClientSettings{
			Endpoint: "http://localhost:8000/redfish/v1",
			Timeout:  10 * time.Second,
		},
		Metrics: metadata.DefaultMetricsSettings(),
	}
}

func createMetricsReceiver(
	_ context.Context,
	params component.ReceiverCreateSettings,
	rConf config.Receiver,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	cfg, ok := rConf.(*Config)
	if !ok {
		return nil, errConfigNotRedfish
	}

	if strings.EqualFold(cfg.Mode, "pull") {

		// we are in a "pull" mode (HTTP GET polling)
		redfishScraper := newRedfishScraper(params, cfg, cfg.Metrics)
		scraper, err := scraperhelper.NewScraper(typeStr, redfishScraper.scrape, scraperhelper.WithStart(redfishScraper.start))
		if err != nil {
			return nil, err
		}

		return scraperhelper.NewScraperControllerReceiver(&cfg.ScraperControllerSettings, params, consumer, scraperhelper.AddScraper(scraper))
	}
	return nil, nil
}

// NewFactory creates a factory for the rrr receiver.
func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiver(createMetricsReceiver, stability))
}
