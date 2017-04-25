package prometheus

import "github.com/turbonomic/monitoring/pkg/client"

type PrometheusMonitor struct {


}

func (monitor *PrometheusMonitor) GetSourceName() client.MONITOR_NAME {
	return client.PROMETHEUS_MESOS
}

func (monitor *PrometheusMonitor) Monitor(target *client.MonitorTarget) (error) {
	// Get the metrics and set the metric values in the repository entities passed in the target object
	return nil
}
