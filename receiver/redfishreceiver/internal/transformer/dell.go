package transformer

import (
	"strconv"

	"eos2git.cec.lab.emc.com/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/metadata"
	redfish "eos2git.cec.lab.emc.com/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/redfish/generated"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

var _ Transformer = (*dellTransformer)(nil)

type dellTransformer struct {
	mb metadata.MetricsBuilder
}

func NewDellTransformer(mb metadata.MetricsBuilder) *dellTransformer {
	return &dellTransformer{mb: mb}
}

func (d *dellTransformer) Transform(reports []redfish.MetricReportV141MetricReport) error {
	for _, report := range reports {
		d.transform(report)
	}
	return nil
}

func (d *dellTransformer) transform(report redfish.MetricReportV141MetricReport) {
	// attempt to not process anything we've seen before

	// process metrics
	for _, metric := range report.GetMetricValues() {
		metricId, ok := metric.GetMetricIdOk()
		if ok {
			switch *metricId {
			case "RPMReading":
				d.buildFanSpeedMetric(metric)
			case "TemperatureReading":
				d.buildTemperatureMetric(metric)
			case "RxBytes":
				d.buildNetworkIoReceive(metric)
			case "TxBytes":
				d.buildNetworkIoTransmit(metric)
			}
		}

	}
}

func (d *dellTransformer) buildFanSpeedMetric(metric redfish.MetricReportV141MetricValue) {
	val, _ := strconv.ParseInt(metric.GetMetricValue(), 10, 64)
	oem := metric.GetOem()

	// id and name can be FQDD
	fqdd := oem["Dell"]["FQDD"].(string)
	d.mb.RecordHwFanSpeedDataPoint(pcommon.NewTimestampFromTime(metric.GetTimestamp()), val, fqdd, fqdd, "")
}

func (d *dellTransformer) buildTemperatureMetric(metric redfish.MetricReportV141MetricValue) {
	temp, _ := strconv.ParseFloat(metric.GetMetricValue(), 64)
	// fmt.Println("writing tempmetric")
	oem := metric.GetOem()

	// id and name can be FQDD
	fqdd := oem["Dell"]["FQDD"].(string)
	d.mb.RecordHwTemperatureDataPoint(pcommon.NewTimestampFromTime(metric.GetTimestamp()), temp, fqdd, fqdd, "")
}

func (d *dellTransformer) buildNetworkIoReceive(metric redfish.MetricReportV141MetricValue) {
	val, _ := strconv.ParseInt(metric.GetMetricValue(), 10, 64)
	// fmt.Println("writing tempmetric")
	oem := metric.GetOem()

	// id and name can be FQDD
	fqdd := oem["Dell"]["FQDD"].(string)
	d.mb.RecordHwNetworkIoReceiveDataPoint(pcommon.NewTimestampFromTime(metric.GetTimestamp()), val, fqdd, fqdd, "")
}

func (d *dellTransformer) buildNetworkIoTransmit(metric redfish.MetricReportV141MetricValue) {
	val, _ := strconv.ParseInt(metric.GetMetricValue(), 10, 64)
	// fmt.Println("writing tempmetric")
	oem := metric.GetOem()

	// id and name can be FQDD
	fqdd := oem["Dell"]["FQDD"].(string)
	d.mb.RecordHwNetworkIoTransmitDataPoint(pcommon.NewTimestampFromTime(metric.GetTimestamp()), val, fqdd, fqdd, "")
}
