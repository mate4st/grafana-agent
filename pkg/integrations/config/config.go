// Package config provides common configuration structs shared among
// implementations of integrations.Integration.
package config

import (
	"flag"
	"time"
)

// Common is a set of common options shared by all integrations. It should be
// utilised by an integration's config by inlining the common options:
//
// type IntegrationConfig struct {
//   Common config.Common `yaml:",inline"`
// }
type Common struct {
	ScrapeInterval time.Duration `yaml:"scrape_interval"`
	ScrapeTimeout  time.Duration `yaml:"scrape_timeout"`
}

func (c *Common) RegisterFlagsWithPrefix(prefix string, f *flag.FlagSet) {
	f.DurationVar(&c.ScrapeInterval, prefix+"scrape-interval", 0, "how frequently should the integration be scraped. 0 = use global default")
	f.DurationVar(&c.ScrapeTimeout, prefix+"scrape-timeout", 0, "timeout for scraping metrics. 0 = use global default")
}

// ScrapeConfig is a subset of options used by integrations to inform how samples
// should be scraped. It is utilized by the integrations.Manager to define a full
// Prometheus-compatible ScrapeConfig.
type ScrapeConfig struct {
	// JobName should be a unique name indicating the collection of samples to be
	// scraped. It will be prepended by "integrations/" when used by the integrations
	// manager.
	JobName string

	// MetricsPath is the path relative to the integration where metrics are exposed.
	// It should match a route added to the router provided in Integration.RegisterRoutes.
	// The path will be prepended by "/integrations/<integration name>" when read by
	// the integrations manager.
	MetricsPath string
}
