package models

import (
	"time"

	"github.com/uptrace/bun"
)

// +metric=slice
type PgStatIO struct {
	bun.BaseModel `bun:"pg_stat_io"`

	BackendType   string    `bun:"backend_type" help:"Type of backend" metric:"backend_type,type:label"`
	Object        string    `bun:"object" help:"Target object of an I/O operation" metric:"object,type:label"`
	Context       string    `bun:"context" help:"The context of an I/O operation" metric:"context,type:label"`
	Reads         int64     `bun:"reads" help:"Number of read operations, each of the size specified in op_bytes" metric:"reads_total"`
	ReadTime      float64   `bun:"read_time" help:"Time spent in read operations in milliseconds" metric:"read_time_total"`
	Writes        int64     `bun:"writes" help:"Number of write operations, each of the size specified in op_bytes" metric:"writes_total"`
	WriteTime     float64   `bun:"write_time" help:"Time spent in write operations in milliseconds" metric:"write_time_total"`
	Writebacks    int64     `bun:"writebacks" help:"Number of units of size op_bytes which the process requested the kernel write out to permanent storage" metric:"writebacks_total"`
	WritebackTime float64   `bun:"writeback_time" help:"Time spent in writeback operations in milliseconds" metric:"writeback_time_total"`
	Extends       int64     `bun:"extends" help:"Number of relation extend operations, each of the size specified in op_bytes" metric:"extends_total"`
	ExtendTime    float64   `bun:"extend_time" help:"Time spent in extend operations in milliseconds" metric:"extend_time_total"`
	OPBytes       int64     `bun:"op_bytes" help:"The number of bytes per unit of I/O read, written, or extended" metric:"op_bytes,label"`
	Hits          int64     `bun:"hits" help:"The number of times a desired block was found in a shared buffer" metric:"hits_total"`
	Evictions     int64     `bun:"evictions" help:"Number of times a block has been written out from a shared or local buffer in order to make it available for another use" metric:"evictions_total"`
	Reuses        int64     `bun:"reuses" help:"The number of times an existing buffer in a size-limited ring buffer outside of shared buffers was reused as part of an I/O operation in the bulkread, bulkwrite, or vacuum contexts" metric:"reuses_total"`
	Fsyncs        int64     `bun:"fsyncs" help:"Number of fsync calls. These are only tracked in context normal" metric:"fsyncs_total"`
	FsyncTime     float64   `bun:"fsync_time" help:"Time spent in fsync operations calls in milliseconds" metric:"fsync_time_total"`
	StatsReset    time.Time `bun:"stats_reset" help:"Time at which these statistics were last reset" metric:"stats_reset_seconds"`
}
