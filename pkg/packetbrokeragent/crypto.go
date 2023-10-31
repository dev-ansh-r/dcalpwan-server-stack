package packetbrokeragent

import "fmt"

func forwarderAdditionalData(netID uint32, tenantID string, clusterID string) []byte {
	return []byte(fmt.Sprintf("%06x:%s:%s", netID, tenantID, clusterID))
}
