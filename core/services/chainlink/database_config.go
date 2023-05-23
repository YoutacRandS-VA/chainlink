package chainlink

import (
	"net/url"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/config"
	v2 "github.com/smartcontractkit/chainlink/v2/core/config/v2"
	"github.com/smartcontractkit/chainlink/v2/core/store/dialects"
)

type backupConfig struct {
	c v2.DatabaseBackup
	s  v2.DatabaseSecrets
}

func (b *backupConfig) Dir() string {
	return *b.c.Dir
}

func (b *backupConfig) Frequency() time.Duration {
	return b.c.Frequency.Duration()
}

func (b *backupConfig) Mode() config.DatabaseBackupMode {
	return *b.c.Mode
}

func (b *backupConfig) OnVersionUpgrade() bool {
	return *b.c.OnVersionUpgrade
}

func (b *backupConfig) URL() *url.URL {
	return b.s.BackupURL.URL()
}

var _ config.Database = (*databaseConfig)(nil)

type databaseConfig struct {
	c  v2.Database
	s  v2.DatabaseSecrets
	mu *sync.RWMutex
}

func (d *databaseConfig) Backup() config.Backup {
	return &backupConfig{
		c: d.c.Backup,
		s: d.s,
	}
}

func (d *databaseConfig) DefaultIdleInTxSessionTimeout() time.Duration {
	return d.c.DefaultIdleInTxSessionTimeout.Duration()
}

func (d *databaseConfig) DefaultLockTimeout() time.Duration {
	return d.c.DefaultLockTimeout.Duration()
}

func (d *databaseConfig) DatabaseDefaultQueryTimeout() time.Duration {
	return d.c.DefaultQueryTimeout.Duration()
}

func (d *databaseConfig) DatabaseListenerMaxReconnectDuration() time.Duration {
	return d.c.Listener.MaxReconnectDuration.Duration()
}

func (d *databaseConfig) DatabaseListenerMinReconnectInterval() time.Duration {
	return d.c.Listener.MinReconnectInterval.Duration()
}

func (d *databaseConfig) DatabaseLockingMode() string {
	return d.c.LockingMode()
}

func (d *databaseConfig) DatabaseURL() url.URL {
	return *d.s.URL.URL()
}

func (d *databaseConfig) GetDatabaseDialectConfiguredOrDefault() dialects.DialectName {
	return d.c.Dialect
}

func (d *databaseConfig) LeaseLockDuration() time.Duration {
	return d.c.Lock.LeaseDuration.Duration()
}

func (d *databaseConfig) LeaseLockRefreshInterval() time.Duration {
	return d.c.Lock.LeaseRefreshInterval.Duration()
}

func (d *databaseConfig) MigrateDatabase() bool {
	return *d.c.MigrateOnStartup
}

func (d *databaseConfig) ORMMaxIdleConns() int {
	return int(*d.c.MaxIdleConns)
}

func (d *databaseConfig) ORMMaxOpenConns() int {
	return int(*d.c.MaxOpenConns)
}

func (d *databaseConfig) TriggerFallbackDBPollInterval() time.Duration {
	return d.c.Listener.FallbackPollInterval.Duration()
}

func (d *databaseConfig) LogSQL() (sql bool) {
	d.mu.RLock()
	sql = *d.c.LogQueries
	d.mu.RUnlock()
	return
}

func (d *databaseConfig) SetLogSQL(logSQL bool) {
	d.mu.Lock()
	d.c.LogQueries = &logSQL
	d.mu.Unlock()
}
