// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"openapi/internal/utils"
)

type ConfigCluster struct {
	Key              *ConfigClusterKey             `json:"key,omitempty"`
	Desc             *string                       `json:"desc,omitempty"`
	Identifier       *string                       `json:"identifier,omitempty"`
	ClusterType      *string                       `json:"clusterType,omitempty"`
	Registry         *string                       `json:"registry,omitempty"`
	Version          *ConfigClusterAPIVersion      `json:"version,omitempty"`
	Platform         *ConfigClusterPlatformInfo    `json:"platform,omitempty"`
	CreateTime       *string                       `json:"createTime,omitempty"`
	CreatedBy        *string                       `json:"createdBy,omitempty"`
	LifecycleManaged *bool                         `json:"lifecycleManaged,omitempty"`
	K8sInfo          *CommonClusterKubeInfo        `json:"k8sInfo,omitempty"`
	HostInfo         *CommonClusterMachineHostInfo `json:"hostInfo,omitempty"`
	Location         *ConfigClusterLocationResp    `json:"location,omitempty"`
	Timezone         *ConfigTimeZone               `json:"timezone,omitempty"`
	EksInfo          *CommonClusterEKSInfo         `json:"eksInfo,omitempty"`
	AksInfo          *CommonClusterAKSInfo         `json:"aksInfo,omitempty"`
	GkeInfo          *CommonClusterGKEInfo         `json:"gkeInfo,omitempty"`
	IsDeleted        *bool                         `json:"isDeleted,omitempty"`
	LifeCycleState   *ClusterLCstate               `default:"Unknown" json:"lifeCycleState"`
	FailedStateError *string                       `json:"failedStateError,omitempty"`
	ProxyEnabled     *bool                         `json:"proxyEnabled,omitempty"`
	ProxyInfo        *CommonProxyInfo              `json:"proxyInfo,omitempty"`
	IsProvider       *bool                         `json:"isProvider,omitempty"`
	CkpProvider      *CommonClusterCKPProvider     `json:"ckpProvider,omitempty"`
}

func (c ConfigCluster) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConfigCluster) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConfigCluster) GetKey() *ConfigClusterKey {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *ConfigCluster) GetDesc() *string {
	if o == nil {
		return nil
	}
	return o.Desc
}

func (o *ConfigCluster) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *ConfigCluster) GetClusterType() *string {
	if o == nil {
		return nil
	}
	return o.ClusterType
}

func (o *ConfigCluster) GetRegistry() *string {
	if o == nil {
		return nil
	}
	return o.Registry
}

func (o *ConfigCluster) GetVersion() *ConfigClusterAPIVersion {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *ConfigCluster) GetPlatform() *ConfigClusterPlatformInfo {
	if o == nil {
		return nil
	}
	return o.Platform
}

func (o *ConfigCluster) GetCreateTime() *string {
	if o == nil {
		return nil
	}
	return o.CreateTime
}

func (o *ConfigCluster) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *ConfigCluster) GetLifecycleManaged() *bool {
	if o == nil {
		return nil
	}
	return o.LifecycleManaged
}

func (o *ConfigCluster) GetK8sInfo() *CommonClusterKubeInfo {
	if o == nil {
		return nil
	}
	return o.K8sInfo
}

func (o *ConfigCluster) GetHostInfo() *CommonClusterMachineHostInfo {
	if o == nil {
		return nil
	}
	return o.HostInfo
}

func (o *ConfigCluster) GetLocation() *ConfigClusterLocationResp {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *ConfigCluster) GetTimezone() *ConfigTimeZone {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *ConfigCluster) GetEksInfo() *CommonClusterEKSInfo {
	if o == nil {
		return nil
	}
	return o.EksInfo
}

func (o *ConfigCluster) GetAksInfo() *CommonClusterAKSInfo {
	if o == nil {
		return nil
	}
	return o.AksInfo
}

func (o *ConfigCluster) GetGkeInfo() *CommonClusterGKEInfo {
	if o == nil {
		return nil
	}
	return o.GkeInfo
}

func (o *ConfigCluster) GetIsDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.IsDeleted
}

func (o *ConfigCluster) GetLifeCycleState() *ClusterLCstate {
	if o == nil {
		return nil
	}
	return o.LifeCycleState
}

func (o *ConfigCluster) GetFailedStateError() *string {
	if o == nil {
		return nil
	}
	return o.FailedStateError
}

func (o *ConfigCluster) GetProxyEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.ProxyEnabled
}

func (o *ConfigCluster) GetProxyInfo() *CommonProxyInfo {
	if o == nil {
		return nil
	}
	return o.ProxyInfo
}

func (o *ConfigCluster) GetIsProvider() *bool {
	if o == nil {
		return nil
	}
	return o.IsProvider
}

func (o *ConfigCluster) GetCkpProvider() *CommonClusterCKPProvider {
	if o == nil {
		return nil
	}
	return o.CkpProvider
}