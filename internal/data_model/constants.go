// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package data_model

import (
	"errors"
	"time"

	"github.com/vkcom/statshouse/internal/vkgo/rpc"
)

func TillStartOfNextSecond(now time.Time) time.Duration { // helper, TODO - move somewhere
	return now.Truncate(time.Second).Add(time.Second).Sub(now)
}

const (
	// Aggregator aggregates data in shards
	// Clients pre calculate shard ID and puts into lower byte of metricID in TL
	// Also clients shuffle shards, so that aggregator is less likely to stop on shard mutex
	// LogAggregationShardsPerSecond should be 8 (Cannot be larger, no point in being smaller)
	LogAggregationShardsPerSecond = 8
	AggregationShardsPerSecond    = 1 << LogAggregationShardsPerSecond

	MaxHistorySendStreams = 24 // Do not change, unless you understand how exactly new conveyor works.
	// Big TODO - using feedback increase the length of this queue in a way it is close to some total memory limit.
	// So if there is single agent with many seconds, queue would grow dramatically (thousand seconds)
	// so that we have enough historic buckets waiting when we finish previous insert
	MaxHistoryInsertContributorsScale = 4 // We will batch less several historic buckets if they contain many contributors

	MaxLivenessResponsesWindowLength = 60
	MaxHistoricBucketsMemorySize     = 50 << 20

	MaxUncompressedBucketSize = 10 << 20 // limits memory for uncompressed data buffer in aggregator, Still dangerous

	MinStringTopCapacity        = 20
	MinStringTopSend            = 5
	MinStringTopInsert          = 5
	AggregatorStringTopCapacity = 1000 // big, but reasonable

	AgentPercentileCompression      = 40 // TODO - will typically have 20-30 centroids for compression 40
	AggregatorPercentileCompression = 80 // TODO - clickhouse has compression of 256 by default

	MaxShortWindow    = 5    // Must be >= 2, 5 seconds to send recent data, if too late - send as historic
	FutureWindow      = 4    // Allow a couple of seconds clocks difference on clients. Plus rounding to multiple of 3
	MaxHistoricWindow = 7200 // 1 day to send historic data, then drop. TODO - return to 86400 after incident

	BelieveTimestampWindow = 86400 + 2*3600 // Margin for crons running once per day.
	// Parts are quickly merged, so all timestamps in [-day..0] will be quickly and thoroughly optimized.

	MinCardinalityWindow = 300 // Our estimators GC depends on this not being too small
	MinMaxCardinality    = 100

	InsertBudgetFixed = 50000
	// fixed budget for BuiltinMetricIDAggKeepAlive and potentially other metrics which can be added with 0 contributors
	// Also helps when # of contributors is very small

	AgentAggregatorDelay = 5  // Typical max
	InsertDelay          = 10 // Typical max

	MaxConveyorDelay = MaxShortWindow + FutureWindow + InsertDelay + AgentAggregatorDelay

	MaxMissedSecondsIntoContributors = 60 // If source sends more MissedSeconds, they will be truncated. Do not make large. We scan 4 arrays of this size on each insert.

	ClientRPCPongTimeout = 30 * time.Second

	AgentMappingTimeout1 = 10 * time.Second
	AgentMappingTimeout2 = 30 * time.Second
	AutoConfigTimeout    = 30 * time.Second

	MaxJournalItemsSent = 1000 // TODO - increase, but limit response size in bytes
	MaxJournalBytesSent = 800 * 1024

	ClickHouseTimeoutConfig = time.Second * 10 // either quickly autoconfig or quickly exit
	ClickhouseConfigRetries = 5
	ClickHouseTimeout       = 5 * time.Minute // reduces chance of duplicates

	KeepAliveMaxBackoff          = 30 * time.Second // for cases when aggregators quickly return error
	JournalDDOSProtectionTimeout = 50 * time.Millisecond

	InternalLogInsertInterval = 5 * time.Second

	RPCErrorMissedRecentConveyor = -5001 // just send again through historic
	RPCErrorInsert               = -5002 // just send again through historic (for recent), or send again after delay (for historic)
	RPCErrorNoAutoCreate         = -5004 // just live with it, this is normal
	RPCErrorTerminateLongpoll    = -5005 // we terminate long polls because rpc.Server does not inform us about hctx disconnects. TODO - remove as soon as Server is updated
	RPCErrorScrapeAgentIP        = -5006 // scrape agent must have IP address

	JournalDiskNamespace        = "metric_journal_v5:"
	TagValueDiskNamespace       = "tag_value_v3:"
	TagValueInvertDiskNamespace = "tag_value_invert_v3:"
	BootstrapDiskNamespace      = "bootstrap:" // stored in aggregator only

	MappingMaxMetricsInQueue = 1000
	MappingMaxMemCacheSize   = 2_000_000
	MappingCacheTTLMinimum   = 7 * 24 * time.Hour
	MappingNegativeCacheTTL  = 5 * time.Second
	MappingMinInterval       = 1 * time.Millisecond

	SimulatorMetricPrefix = "simulator_metric_"

	StatshouseAgentRemoteConfigMetric      = "statshouse_agent_remote_config"
	StatshouseAggregatorRemoteConfigMetric = "statshouse_aggregator_remote_config"
)

func NextBackoffDuration(backoffTimeout time.Duration) time.Duration {
	if backoffTimeout < 0 {
		backoffTimeout = 0
	}
	backoffTimeout = (backoffTimeout + time.Second) * 5 / 4 // exact curve shape is not important
	if backoffTimeout > KeepAliveMaxBackoff {
		backoffTimeout = KeepAliveMaxBackoff
	}
	return backoffTimeout
}

// those can seriously fill our logs, we want to avoid it in a consistent manner
func SilentRPCError(err error) bool {
	var rpcError rpc.Error
	if !errors.As(err, &rpcError) {
		return false
	}
	switch rpcError.Code {
	case RPCErrorMissedRecentConveyor, RPCErrorInsert,
		RPCErrorNoAutoCreate, RPCErrorTerminateLongpoll:
		return true
	default:
		return false
	}
}

func RemoteConfigMetric(name string) bool {
	switch name {
	case StatshouseAgentRemoteConfigMetric, StatshouseAggregatorRemoteConfigMetric:
		return true
	default:
		return false
	}
}
