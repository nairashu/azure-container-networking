package clustersubnetstate

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

// Constants to describe the error state boolean values for the cluster subnet state
const (
	SubnetExhaustionCRDWatcherFailed  = "SubnetExhaustionCRDWatcherFailed"
	SubnetExhaustionCRDWatcherSuccess = "SubnetExhaustionCRDWatcherSuccess"
)

var cssReconcilerErrorCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "cluster_subnet_state_reconciler_error_count_total",
		Help: "Number of errors in reconciler while watching for subnet exhaustion",
	},
	[]string{"reconcilerWatcherState"},
)

func init() {
	metrics.Registry.MustRegister(
		cssReconcilerErrorCount,
	)
}
