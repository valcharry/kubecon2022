package redfishreceiver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	redfish "github.com/valcharry/kubecon2022/receiver/redfishreceiver/internal/redfish/generated"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

// client defines the basic HTTP client interface.
type client interface {
	Get(path string) ([]byte, error)
	GetVendor() (string, error)
	GetMetricReport(report string) (*redfish.MetricReportV141MetricReport, error)
	GetMetricDefinition(metric string) (*redfish.MetricDefinitionV111MetricDefinition, error)
	GetMetricReportsCollection() (*redfish.MetricReportCollectionMetricReportCollection, error)
}

var _ client = (*redfishTelemetryClient)(nil)

type redfishTelemetryClient struct {
	client *http.Client
	cfg    *Config
	logger *zap.Logger
}

// newRedfishTelemetryClient
func newRedfishTelemetryClient(cfg *Config, host component.Host, settings component.TelemetrySettings) (client, error) {
	client, err := cfg.ToClient(host, settings)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP Client: %w", err)
	}

	return &redfishTelemetryClient{
		client: client,
		cfg:    cfg,
		logger: settings.Logger,
	}, nil
}

// Get issues an authorized Get request to the specified url.
func (c *redfishTelemetryClient) Get(path string) ([]byte, error) {
	req, err := c.buildReq(path)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			c.logger.Warn("failed to close response body", zap.Error(err))
		}
	}()

	if resp.StatusCode != http.StatusOK {
		/*
			if resp.StatusCode >= 400 {
				c.logger.Error("redfish telemetry", zap.Error(err), zap.String("status_code", strconv.Itoa(resp.StatusCode)))
			}
		*/
		return nil, fmt.Errorf("request GET %s failed - %q", req.URL.String(), resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body %w", err)
	}

	return body, nil
}

// GetVendor fetches the vendor information of the Redfish compatible endpoint
func (c *redfishTelemetryClient) GetVendor() (string, error) {
	path := fmt.Sprintf("/redfish/v1")
	body, err := c.Get(path)
	if err != nil {
		return "", err
	}
	// c.logger.Info("here we go", zap.String("Body", string(body)))
	serviceRoot := &redfish.ServiceRootV190ServiceRoot{}

	err = json.Unmarshal(body, serviceRoot)
	if err != nil {
		return "", err
	}

	vendor, ok := serviceRoot.GetVendorOk()
	if !ok {
		return "", fmt.Errorf("Unable to query vendor information")
	}

	return *vendor, nil
}

// GetMetricReport gets a single metric report
func (c *redfishTelemetryClient) GetMetricReport(report string) (*redfish.MetricReportV141MetricReport, error) {
	defaultPath := "/redfish/v1/TelemetryService/MetricReports/"
	path := ""
	if strings.Contains(report, defaultPath) {
		path = report
	} else {
		path = fmt.Sprintf("/redfish/v1/TelemetryService/MetricReports/%s", report)
	}

	body, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	// c.logger.Info("here we go", zap.String("Body", string(body)))
	metricReport := &redfish.MetricReportV141MetricReport{}

	err = json.Unmarshal(body, metricReport)
	if err != nil {
		return nil, err
	}

	return metricReport, nil
}

// GetMetricDefinition gets the metric definition for a given metric
func (c *redfishTelemetryClient) GetMetricDefinition(metric string) (*redfish.MetricDefinitionV111MetricDefinition, error) {
	path := fmt.Sprintf("/redfish/v1/TelemetryService/MetricDefinitions/%s", metric)
	body, err := c.Get(path)
	if err != nil {
		return nil, err
	}

	metricDefinition := &redfish.MetricDefinitionV111MetricDefinition{}

	err = json.Unmarshal(body, metricDefinition)
	if err != nil {
		return nil, err
	}

	return metricDefinition, nil
}

// GetMetricReportsCollection gets the collection of metric reports
func (c *redfishTelemetryClient) GetMetricReportsCollection() (*redfish.MetricReportCollectionMetricReportCollection, error) {
	path := "/redfish/v1/TelemetryService/MetricReports"
	body, err := c.Get(path)
	if err != nil {
		return nil, err
	}

	metricReports := &redfish.MetricReportCollectionMetricReportCollection{}

	err = json.Unmarshal(body, metricReports)
	if err != nil {
		return nil, err
	}

	return metricReports, nil
}

// buildReq builds the request with the appropriate accept header and basic authentication
func (c *redfishTelemetryClient) buildReq(path string) (*http.Request, error) {
	url := c.cfg.Endpoint + path
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// we must have this accept header or we will get a 406
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.cfg.Username, c.cfg.Password)
	return req, nil
}
