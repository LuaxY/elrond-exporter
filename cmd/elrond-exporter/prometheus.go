package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	registry *prometheus.Registry
	factory  promauto.Factory
	labels   = []string{"node"}

	metricsStatus     map[string]*prometheus.GaugeVec
	metricsStatistics map[string]*prometheus.GaugeVec
)

func initPrometheus() error {
	registry = prometheus.NewRegistry()
	factory = promauto.With(registry)

	metricsStatus = make(map[string]*prometheus.GaugeVec)
	metricsStatistics = make(map[string]*prometheus.GaugeVec)

	metricsStatistics["erd_stats_live_tps"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_stats_live_tps",
	}, labels)

	metricsStatistics["erd_stats_peak_tps"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_stats_peak_tps",
	}, labels)

	metricsStatistics["erd_stats_block_number"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_stats_block_number",
	}, labels)

	metricsStatistics["erd_stats_round_number"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_stats_round_number",
	}, labels)

	metricsStatistics["erd_stats_round_time"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_stats_round_time",
	}, labels)

	metricsStatistics["erd_stats_average_block_tx_count"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_stats_average_block_tx_count",
	}, labels)

	metricsStatistics["erd_stats_total_processed_tx_count"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_stats_total_processed_tx_count",
	}, labels)

	metricsStatistics["erd_shard_stats_live_tps"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_shard_stats_live_tps",
	}, append(labels, "shard"))

	metricsStatistics["erd_shard_stats_average_tps"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_shard_stats_average_tps",
	}, append(labels, "shard"))

	metricsStatistics["erd_shard_stats_peak_tps"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_shard_stats_peak_tps",
	}, append(labels, "shard"))

	metricsStatistics["erd_shard_stats_current_block_nonce"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_shard_stats_current_block_nonce",
	}, append(labels, "shard"))

	metricsStatistics["erd_shard_stats_total_processed_tx_count"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_shard_stats_total_processed_tx_count",
	}, append(labels, "shard"))

	metricsStatistics["erd_shard_stats_shard_id"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_shard_stats_shard_id",
	}, append(labels, "shard"))

	metricsStatistics["erd_shard_stats_average_block_tx_count"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_shard_stats_average_block_tx_count",
	}, append(labels, "shard"))

	metricsStatistics["erd_shard_stats_last_block_tx_count"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_shard_stats_last_block_tx_count",
	}, append(labels, "shard"))

	metricsStatistics["erd_stats_last_block_tx_count"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_stats_last_block_tx_count",
	}, labels)

	metricsStatistics["erd_stats_last_nr_of_shards"] = factory.NewGaugeVec(prometheus.GaugeOpts{
		Name: "erd_stats_last_nr_of_shards",
	}, labels)

	return nil
}
