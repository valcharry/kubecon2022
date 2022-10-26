package transformer

import redfish "eos2git.cec.lab.emc.com/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/redfish/generated"

// Interface designed to abstract the nuances between the vendor specific details of
// metrics building

type Transformer interface {
	Transform([]redfish.MetricReportV141MetricReport) error
}
