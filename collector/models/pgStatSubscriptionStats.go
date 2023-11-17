package models

import (
	"time"

	"github.com/uptrace/bun"
)

// +metric=slice
type PgStatSubscriptionStats struct {
	bun.BaseModel   `bun:"pg_stat_subscription_stats"`
	SubID           int64     `bun:"subid" metric:"subscription_id,type:label"`
	Subname         string    `bun:"subname" metric:"subscription,type:label"`
	ApplyErrorCount int64     `bun:"apply_error_count" help:"Number of errors that have occurred while applying changes" metric:",type:counter"`
	SyncErrorCount  int64     `bun:"sync_error_count" help:"Number of times an error occurred during the initial table synchronization" metric:",type:counter"`
	StatsReset      time.Time `bun:"stats_reset" help:"Time at which these statistics were last reset"`
}
