package device

import (
	"github.com/gluster/glusterd2/pkg/api"
)

// Info represents structure in which devices are to be store in Peer Metadata
type Info struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

// AddDeviceResp is the success response sent to a AddDeviceReq request
type AddDeviceResp api.Peer

// ListDeviceResp is the success response sent to a ListDevice request
type ListDeviceResp []Info

// ListAllDeviceResp is the success response sent to a ListAllDevice request
type ListAllDeviceResp struct {
	Name   string `json:"name"`
	State  string `json:"state"`
	Peerid string `json:"peerid,omitempty"`
}
