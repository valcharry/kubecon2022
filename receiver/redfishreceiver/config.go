package redfishreceiver

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"eos2git.cec.lab.emc.com/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/metadata"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/multierr"
)

var _ config.Receiver = (*Config)(nil)

var (
	// Errors for missing required config fields.
	errMissingUsername     = errors.New(`no "username" specified in config`)
	errMissingPassword     = errors.New(`no "password" specified in config`)
	errMissingEndpoint     = errors.New(`no "endpoint" specified in config`)
	errMissingMergeField   = errors.New(`configuration - missing at least one required values in the merge stanza (report, highmetric, lowmetric, newmetric, description)`)
	errInvalidReceiverMode = errors.New(`configuration - invalid receiver mode specified (must be sse or pull)`)

	// Errors for invalid url components in the endpoint.
	errInvalidEndpoint = errors.New(`"endpoint" must be in the form of <scheme>://<hostname>:<port>`)
)

type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	confighttp.HTTPClientSettings           `mapstructure:",squash"`
	Username                                string                   `mapstructure:"username"`
	Password                                string                   `mapstructure:"password"`
	Mode                                    string                   `mapstructure:"mode"`
	Reports                                 []string                 `mapstructure:"reports"`
	AdjustTime                              bool                     `mapstructure:"adjusttime"`
	Attributes                              map[string]string        `mapstructure:"attributes"`
	Merge                                   []Merge                  `mapstructure:"merge"`
	Metrics                                 metadata.MetricsSettings `mapstructure:"metrics"`
}

type Merge struct {
	Report      string `mapstructure:"report"`
	HighMetric  string `mapstructure:"highmetric"`
	LowMetric   string `mapstructure:"lowmetric"`
	NewMetric   string `mapstructure:"newmetric"`
	Description string `mapstructure:"description"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {

	var err error
	if receiverErr := cfg.ReceiverSettings.Validate(); receiverErr != nil {
		err = multierr.Append(err, receiverErr)
	}
	if cfg.Username == "" {
		err = multierr.Append(err, errMissingUsername)
	}
	if cfg.Password == "" {
		err = multierr.Append(err, errMissingPassword)
	}
	if cfg.Endpoint == "" {
		err = multierr.Append(err, errMissingEndpoint)
	}
	if !strings.EqualFold(cfg.Mode, "sse") && !strings.EqualFold(cfg.Mode, "pull") {
		err = multierr.Append(err, errInvalidReceiverMode)
	}

	if _, parseErr := url.Parse(cfg.Endpoint); parseErr != nil {
		errInvalidEndpoint = multierr.Append(errInvalidEndpoint, parseErr)
		err = multierr.Append(err, errInvalidEndpoint)
	}
	// cycle through the "merge" stanzas and ensure they have all the values -- will check after initialization if they are valid
	for _, merged := range cfg.Merge {
		if merged.Report == "" || merged.HighMetric == "" || merged.LowMetric == "" || merged.NewMetric == "" || merged.Description == "" {
			return fmt.Errorf("configuration - missing at least one required values in the merge stanza (report, highmetric, lowmetric, newmetric, description)")
		}
	}

	return err
}
