package cocktail

import "sync"

type Config struct {
	IsRelease bool
}

var lock sync.Once
var instance *Config

/**
 * Create single config instance.
 */
func ConfigInstance() *Config {
	// Condition validation
	if instance != nil {
		return instance
	}

	// Create one instance only
	lock.Do(func() {
		if instance == nil {
			instance = new(Config)
		}
		instance.IsRelease = false
	})
	return instance
}
