// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"openapi/internal/utils"
)

type CommonClusterCKPProvider struct {
	Provider *CommonClusterCKPProviderType `default:"BringYourOwnHost" json:"provider"`
	Byoh     *CommonByohProvider           `json:"byoh,omitempty"`
	Aws      *CommonAwsClusterProvider     `json:"aws,omitempty"`
}

func (c CommonClusterCKPProvider) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CommonClusterCKPProvider) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CommonClusterCKPProvider) GetProvider() *CommonClusterCKPProviderType {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *CommonClusterCKPProvider) GetByoh() *CommonByohProvider {
	if o == nil {
		return nil
	}
	return o.Byoh
}

func (o *CommonClusterCKPProvider) GetAws() *CommonAwsClusterProvider {
	if o == nil {
		return nil
	}
	return o.Aws
}
