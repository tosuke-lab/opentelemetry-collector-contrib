// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestMetricsBuilderConfig(t *testing.T) {
	tests := []struct {
		name string
		want MetricsBuilderConfig
	}{
		{
			name: "default",
			want: DefaultMetricsBuilderConfig(),
		},
		{
			name: "all_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					ContainerBlockioIoServiceBytesRecursiveRead:  MetricConfig{Enabled: true},
					ContainerBlockioIoServiceBytesRecursiveWrite: MetricConfig{Enabled: true},
					ContainerCPUPercent:                          MetricConfig{Enabled: true},
					ContainerCPUTime:                             MetricConfig{Enabled: true},
					ContainerCPUUsagePercpu:                      MetricConfig{Enabled: true},
					ContainerCPUUsageSystem:                      MetricConfig{Enabled: true},
					ContainerCPUUsageTotal:                       MetricConfig{Enabled: true},
					ContainerMemoryPercent:                       MetricConfig{Enabled: true},
					ContainerMemoryUsageLimit:                    MetricConfig{Enabled: true},
					ContainerMemoryUsageTotal:                    MetricConfig{Enabled: true},
					ContainerNetworkIoUsageRxBytes:               MetricConfig{Enabled: true},
					ContainerNetworkIoUsageTxBytes:               MetricConfig{Enabled: true},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ContainerID:        ResourceAttributeConfig{Enabled: true},
					ContainerImageName: ResourceAttributeConfig{Enabled: true},
					ContainerName:      ResourceAttributeConfig{Enabled: true},
					ContainerRuntime:   ResourceAttributeConfig{Enabled: true},
				},
			},
		},
		{
			name: "none_set",
			want: MetricsBuilderConfig{
				Metrics: MetricsConfig{
					ContainerBlockioIoServiceBytesRecursiveRead:  MetricConfig{Enabled: false},
					ContainerBlockioIoServiceBytesRecursiveWrite: MetricConfig{Enabled: false},
					ContainerCPUPercent:                          MetricConfig{Enabled: false},
					ContainerCPUTime:                             MetricConfig{Enabled: false},
					ContainerCPUUsagePercpu:                      MetricConfig{Enabled: false},
					ContainerCPUUsageSystem:                      MetricConfig{Enabled: false},
					ContainerCPUUsageTotal:                       MetricConfig{Enabled: false},
					ContainerMemoryPercent:                       MetricConfig{Enabled: false},
					ContainerMemoryUsageLimit:                    MetricConfig{Enabled: false},
					ContainerMemoryUsageTotal:                    MetricConfig{Enabled: false},
					ContainerNetworkIoUsageRxBytes:               MetricConfig{Enabled: false},
					ContainerNetworkIoUsageTxBytes:               MetricConfig{Enabled: false},
				},
				ResourceAttributes: ResourceAttributesConfig{
					ContainerID:        ResourceAttributeConfig{Enabled: false},
					ContainerImageName: ResourceAttributeConfig{Enabled: false},
					ContainerName:      ResourceAttributeConfig{Enabled: false},
					ContainerRuntime:   ResourceAttributeConfig{Enabled: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadMetricsBuilderConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(MetricConfig{}, ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadMetricsBuilderConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}

func TestResourceAttributesConfig(t *testing.T) {
	tests := []struct {
		name string
		want ResourceAttributesConfig
	}{
		{
			name: "default",
			want: DefaultResourceAttributesConfig(),
		},
		{
			name: "all_set",
			want: ResourceAttributesConfig{
				ContainerID:        ResourceAttributeConfig{Enabled: true},
				ContainerImageName: ResourceAttributeConfig{Enabled: true},
				ContainerName:      ResourceAttributeConfig{Enabled: true},
				ContainerRuntime:   ResourceAttributeConfig{Enabled: true},
			},
		},
		{
			name: "none_set",
			want: ResourceAttributesConfig{
				ContainerID:        ResourceAttributeConfig{Enabled: false},
				ContainerImageName: ResourceAttributeConfig{Enabled: false},
				ContainerName:      ResourceAttributeConfig{Enabled: false},
				ContainerRuntime:   ResourceAttributeConfig{Enabled: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, tt.name)
			if diff := cmp.Diff(tt.want, cfg, cmpopts.IgnoreUnexported(ResourceAttributeConfig{})); diff != "" {
				t.Errorf("Config mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func loadResourceAttributesConfig(t *testing.T, name string) ResourceAttributesConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	sub, err = sub.Sub("resource_attributes")
	require.NoError(t, err)
	cfg := DefaultResourceAttributesConfig()
	require.NoError(t, sub.Unmarshal(&cfg))
	return cfg
}
