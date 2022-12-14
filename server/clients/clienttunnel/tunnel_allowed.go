package clienttunnel

import (
	"strings"

	"golang.org/x/crypto/ssh"

	"github.com/cloudradar-monitoring/rport/share/comm"
)

func IsAllowed(remote string, conn ssh.Conn) (bool, error) {
	req := &comm.CheckTunnelAllowedRequest{
		Remote: remote,
	}
	resp := &comm.CheckTunnelAllowedResponse{}
	err := comm.SendRequestAndGetResponse(conn, comm.RequestTypeCheckTunnelAllowed, req, resp)
	if err != nil {
		if strings.Contains(err.Error(), "unknown request") {
			return true, nil
		}
		return false, err
	}

	return resp.IsAllowed, nil
}
