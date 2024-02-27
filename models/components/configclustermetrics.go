// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type ConfigClusterMetrics struct {
	CPU    *ConfigResourceUsageMetrics   `json:"cpu,omitempty"`
	Memory *ConfigResourceUsageMetrics   `json:"memory,omitempty"`
	Nodes  *ConfigNodeStatus             `json:"nodes,omitempty"`
	Pods   *ConfigPodStatus              `json:"pods,omitempty"`
	Apps   *ConfigClusterAppInstanceInfo `json:"apps,omitempty"`
}

func (o *ConfigClusterMetrics) GetCPU() *ConfigResourceUsageMetrics {
	if o == nil {
		return nil
	}
	return o.CPU
}

func (o *ConfigClusterMetrics) GetMemory() *ConfigResourceUsageMetrics {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *ConfigClusterMetrics) GetNodes() *ConfigNodeStatus {
	if o == nil {
		return nil
	}
	return o.Nodes
}

func (o *ConfigClusterMetrics) GetPods() *ConfigPodStatus {
	if o == nil {
		return nil
	}
	return o.Pods
}

func (o *ConfigClusterMetrics) GetApps() *ConfigClusterAppInstanceInfo {
	if o == nil {
		return nil
	}
	return o.Apps
}
