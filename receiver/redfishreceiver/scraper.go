package redfishreceiver

import (
	"context"
	"fmt"
	"strings"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"eos2git.cec.lab.emc.com/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/metadata"
	redfish "eos2git.cec.lab.emc.com/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/redfish/generated"
	"eos2git.cec.lab.emc.com/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/transformer"
)

type redfishScraper struct {
	client           client
	config           *Config
	settings         component.TelemetrySettings
	mb               *metadata.MetricsBuilder
	vendorTranslator transformer.Transformer
}

func newRedfishScraper(settings component.ReceiverCreateSettings, cfg *Config, metricsConfig metadata.MetricsSettings) *redfishScraper {
	return &redfishScraper{
		settings: settings.TelemetrySettings,
		config:   cfg,
		// model:    newRedfishCommonModel(settings),
		mb: metadata.NewMetricsBuilder(metricsConfig, settings.BuildInfo),
	}
}

func (r *redfishScraper) start(_ context.Context, host component.Host) error {

	fmt.Println("rrr - initializing the scraper")
	httpClient, err := newRedfishTelemetryClient(r.config, host, r.settings)
	if err != nil {
		fmt.Printf("redfish scraper start - failed to start the http client %v\n", err)
		return fmt.Errorf("failed to start the telemetry client: %w", err)
	}
	r.client = httpClient

	// need to validate connection + get vendor details
	vendor, err := r.client.GetVendor()
	if err != nil {
		return fmt.Errorf("redfish scraper start - failed to query redfish vendor info: %v\n", err)
	}

	if !strings.EqualFold(vendor, "dell") {
		return fmt.Errorf("unsupported vendor type: %s", vendor)
	} else {
		// Set the dell specific redfish translation
		r.vendorTranslator = transformer.NewDellTransformer(*r.mb)
	}

	// we got this far...
	fmt.Println("rrr - initialized the scraper")
	return nil
}

func (r *redfishScraper) scrape(context.Context) (pmetric.Metrics, error) {

	// ok, we are going to cycle through all of the reports that we want -- and then add metric values to our meters...
	fmt.Println("rrr - starting a scrape")

	// Get Metric Report Collection
	// For every Metric Report in Collection
	// Get Metric Report
	reportCollection, _ := r.client.GetMetricReportsCollection()

	members := reportCollection.GetMembers()
	metricReports := make([]redfish.MetricReportV141MetricReport, 0)
	for _, report := range members {
		metricReport, _ := r.client.GetMetricReport(report.GetOdataId())
		metricReports = append(metricReports, *metricReport)
	}
	// Custom vendor translation

	err := r.vendorTranslator.Transform(metricReports)

	return r.mb.Emit(), err

}
