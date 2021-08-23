package metrics

import "github.com/kuvalda989/prom-exporter/config"

// структура пром метрики
type PromMetric struct {
	Name  string
	Tags  map[string]string
	Value float64
}

func GetMetrics(config config.Config) []PromMetric {
	metric_slice := []PromMetric{}
	if config.Source_type == "file" {
		metric_slice = GetMetricFromFile(config.Source)
	}
	return metric_slice
}
