package cocktail

import "sync"

type Config struct {
	IsDebug bool
}

var lock sync.Once
var instance *Config

/**
 * Create single config instance.
 */
func ConfigInstance() *Config {
	lock.Do(func() {
		instance = new(Config)
		instance.IsDebug = true
	})
	return instance
}
