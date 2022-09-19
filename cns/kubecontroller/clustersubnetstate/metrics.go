package clustersubnetstate

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	cssReconcilerErrorState = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cluster_subnet_state_reconciler_error_state",
			Help: "Reconciler Error",
		},
	)
)

func init() {
	metrics.Registry.MustRegister(
		cssReconcilerErrorState,
	)
}
