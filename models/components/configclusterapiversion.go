// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type ConfigClusterAPIVersion struct {
	MajorVersion *string `json:"majorVersion,omitempty"`
	MinorVersion *string `json:"minorVersion,omitempty"`
	GitVersion   *string `json:"gitVersion,omitempty"`
	Platform     *string `json:"platform,omitempty"`
}

func (o *ConfigClusterAPIVersion) GetMajorVersion() *string {
	if o == nil {
		return nil
	}
	return o.MajorVersion
}

func (o *ConfigClusterAPIVersion) GetMinorVersion() *string {
	if o == nil {
		return nil
	}
	return o.MinorVersion
}

func (o *ConfigClusterAPIVersion) GetGitVersion() *string {
	if o == nil {
		return nil
	}
	return o.GitVersion
}

func (o *ConfigClusterAPIVersion) GetPlatform() *string {
	if o == nil {
		return nil
	}
	return o.Platform
}
