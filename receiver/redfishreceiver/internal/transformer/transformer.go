package transformer

import redfish "github.com/valcharry/kubecon2022/receiver/redfishreceiver/internal/redfish/generated"

// Interface designed to abstract the nuances between the vendor specific details of
// metrics building

type Transformer interface {
	Transform([]redfish.MetricReportV141MetricReport) error
}
