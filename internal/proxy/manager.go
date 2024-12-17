package proxy

import (
	"dip/internal/errors"
	"strings"
	"sync"
)

var (
	ProxyManager DipProxyManager
	once         sync.Once
)

type DipProxyManager struct {
	extractProxyCache map[string]DipProxy
	prefixProxyCache  map[string]DipProxy
}

func InitDipProxyManager() {
	once.Do(func() {
		ProxyManager = DipProxyManager{}
	})
}

func (dipProxyManager DipProxyManager) MatchPath(path string) (DipProxy, error) {
	// Extract Match
	dipProxy, exists := dipProxyManager.extractProxyCache[path]
	if exists {
		return dipProxy, nil
	}
	// Prefix Match
	if len(dipProxyManager.prefixProxyCache) == 0 {
		return DipProxy{}, errors.NotMatchRouterError
	}
	for _, value := range dipProxyManager.prefixProxyCache {
		if strings.HasPrefix(path, value.Path) {
			return value, nil
		}
	}
	return DipProxy{}, errors.NotMatchRouterError
}
