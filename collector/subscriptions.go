package collector

import (
	"context"

	"github.com/1and1/pg-exporter/collector/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"
)

const (
	// subsystem
	subscriptions = "subscriptions"
)

// ScrapeSubscriptions scrapes from pg_stat_subscriptions_stats
type ScrapeSubscriptions struct{}

// Name of the Scraper
func (ScrapeSubscriptions) Name() string {
	return "pg_stat_subscriptions_stats"
}

// Help describes the role of the Scraper
func (ScrapeSubscriptions) Help() string {
	return "Collect from pg_stat_subscriptions_stats"
}

// Version returns minimum PostgreSQL version
func (ScrapeSubscriptions) Version() int {
	return 150000
}

// Type returns the scrape type
func (ScrapeSubscriptions) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeSubscriptions) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	statSubscriptions := &models.PgStatSubscriptionStatsSlice{}
	if err := db.NewSelect().Model(statSubscriptions).Scan(ctx); err != nil {
		return err
	}

	return statSubscriptions.ToMetrics(namespace, subscriptions, ch)
}
