package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/uptrace/bun"

	"github.com/1and1/pg-exporter/collector/models"
)

const (
	// subsystem
	io = "io"
)

// ScrapeIO scrapes from pg_stat_io
type ScrapeIO struct{}

// Name of the Scraper
func (ScrapeIO) Name() string {
	return "pg_stat_io"
}

// Help describes the role of the Scraper
func (ScrapeIO) Help() string {
	return "Collect from pg_stat_io"
}

// Version returns minimum PostgreSQL version
func (ScrapeIO) Version() int {
	return 160000
}

// Type returns the scrape type
func (ScrapeIO) Type() ScrapeType {
	return SCRAPEGLOBAL
}

// Scrape collects data from database connection and sends it over channel as prometheus metric.
func (ScrapeIO) Scrape(ctx context.Context, db *bun.DB, ch chan<- prometheus.Metric) error {
	var statIO models.PgStatIO
	if err := db.NewSelect().Model(&statIO).Scan(ctx); err != nil {
		return err
	}

	return statIO.ToMetrics(namespace, io, ch)
}
