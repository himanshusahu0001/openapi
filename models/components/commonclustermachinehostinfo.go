// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type CommonClusterMachineHostInfo struct {
	MasterHostGroup *string `json:"masterHostGroup,omitempty"`
	WorkerHostGroup *string `json:"workerHostGroup,omitempty"`
}

func (o *CommonClusterMachineHostInfo) GetMasterHostGroup() *string {
	if o == nil {
		return nil
	}
	return o.MasterHostGroup
}

func (o *CommonClusterMachineHostInfo) GetWorkerHostGroup() *string {
	if o == nil {
		return nil
	}
	return o.WorkerHostGroup
}
