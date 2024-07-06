// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for podman_stats metrics.
type MetricsConfig struct {
	ContainerBlockioIoServiceBytesRecursiveRead  MetricConfig `mapstructure:"container.blockio.io_service_bytes_recursive.read"`
	ContainerBlockioIoServiceBytesRecursiveWrite MetricConfig `mapstructure:"container.blockio.io_service_bytes_recursive.write"`
	ContainerCPUPercent                          MetricConfig `mapstructure:"container.cpu.percent"`
	ContainerCPUTime                             MetricConfig `mapstructure:"container.cpu.time"`
	ContainerCPUUsagePercpu                      MetricConfig `mapstructure:"container.cpu.usage.percpu"`
	ContainerCPUUsageSystem                      MetricConfig `mapstructure:"container.cpu.usage.system"`
	ContainerCPUUsageTotal                       MetricConfig `mapstructure:"container.cpu.usage.total"`
	ContainerDiskIo                              MetricConfig `mapstructure:"container.disk.io"`
	ContainerMemoryPercent                       MetricConfig `mapstructure:"container.memory.percent"`
	ContainerMemoryUsage                         MetricConfig `mapstructure:"container.memory.usage"`
	ContainerMemoryUsageLimit                    MetricConfig `mapstructure:"container.memory.usage.limit"`
	ContainerMemoryUsageTotal                    MetricConfig `mapstructure:"container.memory.usage.total"`
	ContainerNetworkIo                           MetricConfig `mapstructure:"container.network.io"`
	ContainerNetworkIoUsageRxBytes               MetricConfig `mapstructure:"container.network.io.usage.rx_bytes"`
	ContainerNetworkIoUsageTxBytes               MetricConfig `mapstructure:"container.network.io.usage.tx_bytes"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		ContainerBlockioIoServiceBytesRecursiveRead: MetricConfig{
			Enabled: true,
		},
		ContainerBlockioIoServiceBytesRecursiveWrite: MetricConfig{
			Enabled: true,
		},
		ContainerCPUPercent: MetricConfig{
			Enabled: true,
		},
		ContainerCPUTime: MetricConfig{
			Enabled: true,
		},
		ContainerCPUUsagePercpu: MetricConfig{
			Enabled: true,
		},
		ContainerCPUUsageSystem: MetricConfig{
			Enabled: true,
		},
		ContainerCPUUsageTotal: MetricConfig{
			Enabled: true,
		},
		ContainerDiskIo: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryPercent: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryUsage: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryUsageLimit: MetricConfig{
			Enabled: true,
		},
		ContainerMemoryUsageTotal: MetricConfig{
			Enabled: true,
		},
		ContainerNetworkIo: MetricConfig{
			Enabled: true,
		},
		ContainerNetworkIoUsageRxBytes: MetricConfig{
			Enabled: true,
		},
		ContainerNetworkIoUsageTxBytes: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Experimental: MetricsInclude defines a list of filters for attribute values.
	// If the list is not empty, only metrics with matching resource attribute values will be emitted.
	MetricsInclude []filter.Config `mapstructure:"metrics_include"`
	// Experimental: MetricsExclude defines a list of filters for attribute values.
	// If the list is not empty, metrics with matching resource attribute values will not be emitted.
	// MetricsInclude has higher priority than MetricsExclude.
	MetricsExclude []filter.Config `mapstructure:"metrics_exclude"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for podman_stats resource attributes.
type ResourceAttributesConfig struct {
	ContainerID        ResourceAttributeConfig `mapstructure:"container.id"`
	ContainerImageName ResourceAttributeConfig `mapstructure:"container.image.name"`
	ContainerName      ResourceAttributeConfig `mapstructure:"container.name"`
	ContainerRuntime   ResourceAttributeConfig `mapstructure:"container.runtime"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		ContainerID: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerImageName: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerName: ResourceAttributeConfig{
			Enabled: true,
		},
		ContainerRuntime: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for podman_stats metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
