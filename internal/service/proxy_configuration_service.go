package service

type ProxyConfiguration interface {
	applyProxyConfiguration() error
	deleteProxyConfiguration() error
}

type ProxyConfigurationService struct {
}

var _proxyConfiguration = (*ProxyConfigurationService)(nil)

func (proxyConfigurationService *ProxyConfigurationService) applyProxyConfiguration() error {

	return nil
}

func (proxyConfigurationService *ProxyConfigurationService) deleteProxyConfiguration() error {
	return nil
}
