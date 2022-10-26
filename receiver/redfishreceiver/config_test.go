package redfishreceiver

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/confighttp"
)

func TestConfigValidation(t *testing.T) {
	cases := []struct {
		desc        string
		cfg         Config
		expectedErr error
	}{
		{
			desc: "valid (no merge stanza)",
			cfg: Config{
				HTTPClientSettings: confighttp.HTTPClientSettings{
					Endpoint: "http://some.endpoint",
				},
				Username: "otelu",
				Password: "otelp",
				Mode:     "pull",
			},
		},
		{
			desc: "valid (w/ merge stanza)",
			cfg: Config{
				HTTPClientSettings: confighttp.HTTPClientSettings{
					Endpoint: "http://some.endpoint",
				},
				Username: "otelu",
				Password: "otelp",
				Mode:     "pull",
				Merge: []Merge{
					{
						Report:      "thisreport",
						HighMetric:  "highmetric",
						LowMetric:   "lowmetric",
						NewMetric:   "newmetric",
						Description: "descriptive",
					},
				},
			},
		},
		{
			desc: "empty endpoint",
			cfg: Config{
				HTTPClientSettings: confighttp.HTTPClientSettings{
					Endpoint: "",
				},
			},
			expectedErr: errMissingEndpoint,
		},
		{
			desc: "unparseable URL",
			cfg: Config{
				HTTPClientSettings: confighttp.HTTPClientSettings{
					Endpoint: "h" + string(rune(0x7f)),
					// TLSClientSetting: configtls.TLSClientSetting{},
				},
			},
			expectedErr: errInvalidEndpoint,
		},
		{
			desc: "no username",
			cfg: Config{
				HTTPClientSettings: confighttp.HTTPClientSettings{
					Endpoint: "http://some.endpoint",
				},
				Password: "otelp",
			},
			expectedErr: errMissingUsername,
		},
		{
			desc: "no password",
			cfg: Config{
				HTTPClientSettings: confighttp.HTTPClientSettings{
					Endpoint: "http://some.endpoint",
				},
				Username: "otelu",
			},
			expectedErr: errMissingPassword,
		},
		{
			desc: "missing merge stanza field",
			cfg: Config{
				HTTPClientSettings: confighttp.HTTPClientSettings{
					Endpoint: "http://some.endpoint",
				},
				Merge: []Merge{
					{
						Report: "thisreport",
					},
				},
			},
			expectedErr: errMissingMergeField,
		},
		{
			desc: "invalid receiver mode",
			cfg: Config{
				HTTPClientSettings: confighttp.HTTPClientSettings{
					Endpoint: "http://some.endpoint",
				},
				Mode: "",
			},
			expectedErr: errInvalidReceiverMode,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.cfg.Validate()
			if tc.expectedErr != nil {
				require.ErrorContains(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
