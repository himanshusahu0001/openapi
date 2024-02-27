// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type CommonClusterAKSInfo struct {
	RgName     *string                   `json:"rgName,omitempty"`
	Region     *string                   `json:"region,omitempty"`
	K8sVersion *string                   `json:"k8sVersion,omitempty"`
	NodeInfo   *CommonClusterAKSNodeInfo `json:"nodeInfo,omitempty"`
	Provider   *string                   `json:"provider,omitempty"`
}

func (o *CommonClusterAKSInfo) GetRgName() *string {
	if o == nil {
		return nil
	}
	return o.RgName
}

func (o *CommonClusterAKSInfo) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *CommonClusterAKSInfo) GetK8sVersion() *string {
	if o == nil {
		return nil
	}
	return o.K8sVersion
}

func (o *CommonClusterAKSInfo) GetNodeInfo() *CommonClusterAKSNodeInfo {
	if o == nil {
		return nil
	}
	return o.NodeInfo
}

func (o *CommonClusterAKSInfo) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}
