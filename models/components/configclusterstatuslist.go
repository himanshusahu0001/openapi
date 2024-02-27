// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type ConfigClusterStatusList struct {
	Count         *int                  `json:"count,omitempty"`
	Total         *int                  `json:"total,omitempty"`
	Connected     *int                  `json:"connected,omitempty"`
	NotConnected  *int                  `json:"notConnected,omitempty"`
	NotRegistered *int                  `json:"notRegistered,omitempty"`
	Items         []ConfigClusterStatus `json:"items,omitempty"`
}

func (o *ConfigClusterStatusList) GetCount() *int {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ConfigClusterStatusList) GetTotal() *int {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *ConfigClusterStatusList) GetConnected() *int {
	if o == nil {
		return nil
	}
	return o.Connected
}

func (o *ConfigClusterStatusList) GetNotConnected() *int {
	if o == nil {
		return nil
	}
	return o.NotConnected
}

func (o *ConfigClusterStatusList) GetNotRegistered() *int {
	if o == nil {
		return nil
	}
	return o.NotRegistered
}

func (o *ConfigClusterStatusList) GetItems() []ConfigClusterStatus {
	if o == nil {
		return nil
	}
	return o.Items
}