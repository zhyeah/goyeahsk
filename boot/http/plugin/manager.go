package plugin

import "sync"

var pluginCollectorOnce sync.Once
var pluginCollector *PluginCollector

// PluginCollector 插件集合
type PluginCollector struct {
	Plugins []Plugin
}

// AddPlugins add plugins
func (collector *PluginCollector) AddPlugins(plugin Plugin) {
	collector.Plugins = append(collector.Plugins, plugin)
}

// GetPluginCollector 获取插件集合单例
func GetPluginCollector() *PluginCollector {
	if pluginCollector == nil {
		pluginCollectorOnce.Do(func() {
			pluginCollector = &PluginCollector{
				Plugins: make([]Plugin, 0),
			}
		})
	}
	return pluginCollector
}
