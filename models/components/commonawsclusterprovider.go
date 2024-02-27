// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type CommonAwsClusterProvider struct {
	Cloud              *string `json:"cloud,omitempty"`
	Region             *string `json:"region,omitempty"`
	MasterInstanceType *string `json:"masterInstanceType,omitempty"`
	WorkerInstanceType *string `json:"workerInstanceType,omitempty"`
}

func (o *CommonAwsClusterProvider) GetCloud() *string {
	if o == nil {
		return nil
	}
	return o.Cloud
}

func (o *CommonAwsClusterProvider) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *CommonAwsClusterProvider) GetMasterInstanceType() *string {
	if o == nil {
		return nil
	}
	return o.MasterInstanceType
}

func (o *CommonAwsClusterProvider) GetWorkerInstanceType() *string {
	if o == nil {
		return nil
	}
	return o.WorkerInstanceType
}