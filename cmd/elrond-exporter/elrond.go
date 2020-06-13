package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/LuaxY/elrond-exporter/pkg/elrond"
)

func updateStatus(node *elrond.Node) error {
	status, err := node.Status()

	if err != nil {
		return errors.Wrapf(err, "get status for node '%s'", node.Name)
	}

	for key, value := range status.Details {
		metric, ok := metricsStatus[key]

		if !ok {
			switch v := value.(type) {
			case int, int64, float64:
				metric = factory.NewGaugeVec(prometheus.GaugeOpts{
					Name: key,
				}, labels)
			case string:
				metric = factory.NewGaugeVec(prometheus.GaugeOpts{
					Name: key,
				}, append(labels, "value"))
			default:
				fmt.Printf("warn: type unknown %s: %T!\n", key, v)
			}

			metricsStatus[key] = metric
		}

		switch v := value.(type) {
		case int:
			metric.WithLabelValues(node.Name).Set(float64(value.(int)))
		case int64:
			metric.WithLabelValues(node.Name).Set(float64(value.(int64)))
		case float64:
			metric.WithLabelValues(node.Name).Set(value.(float64))
		case string:
			metric.WithLabelValues(node.Name, v).Set(1)
		default:
		}
	}

	return nil
}

func updateStatistics(node *elrond.Node) error {
	statistics, err := node.Statistics()

	if err != nil {
		return errors.Wrapf(err, "get statistics for node '%s'", node.Name)
	}

	if statistics.Statistics.LiveTPS != nil {
		metricsStatistics["erb_stats_live_tps"].WithLabelValues(node.Name).Set(*statistics.Statistics.LiveTPS)
	}

	if statistics.Statistics.PeakTPS != nil {
		metricsStatistics["erb_stats_peak_tps"].WithLabelValues(node.Name).Set(*statistics.Statistics.PeakTPS)
	}

	if statistics.Statistics.BlockNumber != nil {
		metricsStatistics["erb_stats_block_number"].WithLabelValues(node.Name).Set(*statistics.Statistics.BlockNumber)
	}

	if statistics.Statistics.RoundNumber != nil {
		metricsStatistics["erb_stats_round_number"].WithLabelValues(node.Name).Set(*statistics.Statistics.RoundNumber)
	}

	if statistics.Statistics.RoundTime != nil {
		metricsStatistics["erb_stats_round_time"].WithLabelValues(node.Name).Set(*statistics.Statistics.RoundTime)
	}

	if statistics.Statistics.AverageBlockTxCount != nil {
		metricsStatistics["erb_stats_average_block_tx_count"].WithLabelValues(node.Name).Set(*statistics.Statistics.AverageBlockTxCount)
	}

	if statistics.Statistics.TotalProcessedTxCount != nil {
		metricsStatistics["erb_stats_total_processed_tx_count"].WithLabelValues(node.Name).Set(*statistics.Statistics.TotalProcessedTxCount)
	}

	for i, shard := range statistics.Statistics.ShardStatistics {
		id := strconv.Itoa(i)

		if shard.LiveTPS != nil {
			metricsStatistics["erb_shard_stats_live_tps"].WithLabelValues(node.Name, id).Set(*shard.LiveTPS)
		}

		if shard.AverageTPS != nil {
			metricsStatistics["erb_shard_stats_average_tps"].WithLabelValues(node.Name, id).Set(*shard.AverageTPS)
		}

		if shard.PeakTPS != nil {
			metricsStatistics["erb_shard_stats_peak_tps"].WithLabelValues(node.Name, id).Set(*shard.PeakTPS)
		}

		if shard.CurrentBlockNonce != nil {
			metricsStatistics["erb_shard_stats_current_block_nonce"].WithLabelValues(node.Name, id).Set(*shard.CurrentBlockNonce)
		}

		if shard.TotalProcessedTxCount != nil {
			metricsStatistics["erb_shard_stats_total_processed_tx_count"].WithLabelValues(node.Name, id).Set(*shard.TotalProcessedTxCount)
		}

		if shard.ShardID != nil {
			metricsStatistics["erb_shard_stats_shard_id"].WithLabelValues(node.Name, id).Set(*shard.ShardID)
		}

		if shard.AverageBlockTxCount != nil {
			metricsStatistics["erb_shard_stats_average_block_tx_count"].WithLabelValues(node.Name, id).Set(*shard.AverageBlockTxCount)
		}

		if shard.LastBlockTxCount != nil {
			metricsStatistics["erb_shard_stats_last_block_tx_count"].WithLabelValues(node.Name, id).Set(*shard.LastBlockTxCount)
		}
	}

	if statistics.Statistics.LastBlockTxCount != nil {
		metricsStatistics["erb_stats_last_block_tx_count"].WithLabelValues(node.Name).Set(*statistics.Statistics.LastBlockTxCount)
	}

	if statistics.Statistics.NrOfShards != nil {
		metricsStatistics["erb_stats_last_nr_of_shards"].WithLabelValues(node.Name).Set(*statistics.Statistics.NrOfShards)
	}

	return nil
}
