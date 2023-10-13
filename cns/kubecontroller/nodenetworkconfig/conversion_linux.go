package nodenetworkconfig

import (
	"net/netip"
	"strconv"
	"strings"

	"github.com/Azure/azure-container-networking/cns"
	"github.com/Azure/azure-container-networking/crd/nodenetworkconfig/api/v1alpha"
	"github.com/pkg/errors"
)

// createNCRequestFromStaticNCHelper generates a CreateNetworkContainerRequest from a static NetworkContainer
// by adding all IPs in the the block to the secondary IP configs list. It does not skip any IPs.
//
//nolint:gocritic //ignore hugeparam
func createNCRequestFromStaticNCHelper(nc v1alpha.NetworkContainer, primaryIPPrefix netip.Prefix, subnet cns.IPSubnet) (*cns.CreateNetworkContainerRequest, error) {
	secondaryIPConfigs := map[string]cns.SecondaryIPConfig{}

	// iterate through all IP addresses in the subnet described by primaryPrefix and
	// add them to the request as secondary IPConfigs.
	for addr := primaryIPPrefix.Masked().Addr(); primaryIPPrefix.Contains(addr); addr = addr.Next() {
		secondaryIPConfigs[addr.String()] = cns.SecondaryIPConfig{
			IPAddress: addr.String(),
			NCVersion: int(nc.Version),
		}
	}

	// IP Configurations for Primary IP and Gateway Mask
	ipConfig := cns.IPConfiguration{
		IPSubnet:         subnet,
		GatewayIPAddress: nc.DefaultGateway,
	}

	// Add IPs from CIDR block to the secondary IPConfigs
	if nc.Type == v1alpha.VNETBlock {
		for _, ipAssignment := range nc.IPAssignments {
			cidrPrefix, err := netip.ParsePrefix(ipAssignment.IP)
			if err != nil {
				return nil, errors.Wrapf(err, "invalid CIDR block: %s", ipAssignment.IP)
			}

			// iterate through all IP addresses in the CIDR block described by cidrPrefix and
			// add them to the request as secondary IPConfigs.
			for addr := cidrPrefix.Masked().Addr(); cidrPrefix.Contains(addr); addr = addr.Next() {
				secondaryIPConfigs[addr.String()] = cns.SecondaryIPConfig{
					IPAddress: addr.String(),
					NCVersion: int(nc.Version),
				}
			}
		}
		// if the NodeIP is not a CIDR, append a /32
		if !strings.Contains(nc.NodeIP, "/") {
			nc.NodeIP += "/32"
		}
		nodeIPPrefix, err := netip.ParsePrefix(nc.NodeIP)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid Node IP %s", nc.NodeIP)
		}
		ipConfig.IPSubnet = cns.IPSubnet{
			IPAddress:    nodeIPPrefix.Addr().String(),
			PrefixLength: uint8(nodeIPPrefix.Bits()),
		}
	}

	return &cns.CreateNetworkContainerRequest{
		SecondaryIPConfigs:   secondaryIPConfigs,
		NetworkContainerid:   nc.ID,
		NetworkContainerType: cns.Docker,
		Version:              strconv.FormatInt(nc.Version, 10), //nolint:gomnd // it's decimal
		IPConfiguration:      ipConfig,
		NCStatus:             nc.Status,
	}, nil
}
