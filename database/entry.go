package database

type DatabaseEntry struct {
	DatabaseConfig DatabaseConfig
}

func ProvideDatabaseEntry(p DatabaseConfig) DatabaseEntry {
	return DatabaseEntry{DatabaseConfig: p}
}

func (p *DatabaseEntry) Connect() *DatabaseConfig {
	return p.DatabaseConfig.Connect()
}
