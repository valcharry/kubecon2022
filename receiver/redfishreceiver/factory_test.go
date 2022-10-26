package redfishreceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
)

func TestType(t *testing.T) {
	factory := NewFactory()
	require.EqualValues(t, "redfish", string(factory.Type()))
}

func TestCreateMetricsReceiver(t *testing.T) {
	testCases := []struct {
		desc string
		run  func(t *testing.T)
	}{
		{
			desc: "Default config",
			run: func(t *testing.T) {
				t.Parallel()

				_, err := createMetricsReceiver(
					context.Background(),
					componenttest.NewNopReceiverCreateSettings(),
					createDefaultConfig(),
					consumertest.NewNop(),
				)

				require.NoError(t, err)
			},
		},
		{
			desc: "Nil config",
			run: func(t *testing.T) {
				t.Parallel()

				_, err := createMetricsReceiver(
					context.Background(),
					componenttest.NewNopReceiverCreateSettings(),
					nil,
					consumertest.NewNop(),
				)
				require.ErrorIs(t, err, errConfigNotRedfish)
			},
		},
		{
			desc: "Nil consumer",
			run: func(t *testing.T) {
				t.Parallel()
				_, err := createMetricsReceiver(
					context.Background(),
					componenttest.NewNopReceiverCreateSettings(),
					createDefaultConfig(),
					nil,
				)
				require.ErrorIs(t, err, component.ErrNilNextConsumer)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, testCase.run)
	}
}
