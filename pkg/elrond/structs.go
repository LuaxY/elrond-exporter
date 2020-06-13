package elrond

type StatusInfo struct {
	Details map[string]interface{} `json:"details"`
}

type StatisticsInfo struct {
	Statistics struct {
		LiveTPS               *float64 `json:"liveTPS"`
		PeakTPS               *float64 `json:"peakTPS"`
		BlockNumber           *float64 `json:"blockNumber"`
		RoundNumber           *float64 `json:"roundNumber"`
		RoundTime             *float64 `json:"roundTime"`
		AverageBlockTxCount   *float64 `json:"averageBlockTxCount"`
		TotalProcessedTxCount *float64 `json:"totalProcessedTxCount"`
		ShardStatistics       []struct {
			LiveTPS               *float64 `json:"liveTPS"`
			AverageTPS            *float64 `json:"averageTPS"`
			PeakTPS               *float64 `json:"peakTPS"`
			CurrentBlockNonce     *float64 `json:"currentBlockNonce"`
			TotalProcessedTxCount *float64 `json:"totalProcessedTxCount"`
			ShardID               *float64 `json:"shardID"`
			AverageBlockTxCount   *float64 `json:"averageBlockTxCount"`
			LastBlockTxCount      *float64 `json:"lastBlockTxCount"`
		} `json:"shardStatistics"`
		LastBlockTxCount *float64 `json:"lastBlockTxCount"`
		NrOfShards       *float64 `json:"nrOfShards"`
	} `json:"statistics"`
}
