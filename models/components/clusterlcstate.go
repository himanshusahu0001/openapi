// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type ClusterLCstate string

const (
	ClusterLCstateUnknown       ClusterLCstate = "Unknown"
	ClusterLCstateCreating      ClusterLCstate = "Creating"
	ClusterLCstateCreateFailed  ClusterLCstate = "CreateFailed"
	ClusterLCstateUpdating      ClusterLCstate = "Updating"
	ClusterLCstateUpdateFailed  ClusterLCstate = "UpdateFailed"
	ClusterLCstateDeleting      ClusterLCstate = "Deleting"
	ClusterLCstateDeleteFailed  ClusterLCstate = "DeleteFailed"
	ClusterLCstateDeletePending ClusterLCstate = "DeletePending"
	ClusterLCstateReady         ClusterLCstate = "Ready"
)

func (e ClusterLCstate) ToPointer() *ClusterLCstate {
	return &e
}

func (e *ClusterLCstate) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Creating":
		fallthrough
	case "CreateFailed":
		fallthrough
	case "Updating":
		fallthrough
	case "UpdateFailed":
		fallthrough
	case "Deleting":
		fallthrough
	case "DeleteFailed":
		fallthrough
	case "DeletePending":
		fallthrough
	case "Ready":
		*e = ClusterLCstate(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClusterLCstate: %v", v)
	}
}
