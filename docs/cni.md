# Microsoft Azure Container Networking

## Azure VNET CNI Plugins
`azure-vnet` CNI plugin implements the [CNI network plugin interface](https://github.com/containernetworking/cni/blob/master/SPEC.md).

`azure-vnet-ipam` CNI plugin implements the [CNI IPAM plugin interface](https://github.com/containernetworking/cni/blob/master/SPEC.md#ip-address-management-ipam-interface).

The plugins are available on both Linux and Windows platforms.

The network and IPAM plugins are designed to work together. The IPAM plugin can also be used by 3rd party software to manage IP addresses from Azure VNET space.

This page describes how to setup the CNI plugins manually on Azure IaaS VMs. If you are planning to deploy an ACS cluster, see [ACS](acs.md) instead.

## Install
Copy the plugin package from the [release](https://github.com/Azure/azure-container-networking/releases) share to your Azure VM and extract the contents to the CNI directories.

You can also install by running the `install-cni-plugin.sh` (Linux) or `install-cni-plugin.ps1` (Windows) scripts provided in the scripts directory of this repository.

```bash
$ scripts/install-cni-plugin.sh [version]
```

```PowerShell
PS> scripts\install-cni-plugin.ps1 [version]
```

The plugin package comes with a simple network configuration file that works out of the box. See the [network configuration](https://github.com/Azure/azure-container-networking/blob/master/docs/cni.md#network-configuration) section below for customization options.

## Build
Plugins can also be built directly from the source code in this repository.

```bash
make azure-vnet
make azure-vnet-ipam
make azure-cni-plugins
```

The first two commands build an individual plugin, whereas the third one builds both and generates a tar archive. The binaries are placed in the `output` directory.

## Network Configuration
Network configuration for CNI plugins is described in JSON format. The default location for configuration files is `/etc/cni/net.d` for Linux and `c:\cni` for Windows.

```json
{
  "cniVersion": "0.2.0",
  "name": "azure",
  "type": "azure-vnet",
  "master": "eth0",
  "bridge": "azure0",
  "logLevel": "info",
  "ipam": {
    "type": "azure-vnet-ipam",
    "environment": "azure"
  }
}
```

The following fields are well-known and have the following meaning:

Network plugin
* `cniVersion`: Azure plugins currently support versions 0.1.0 and 0.2.0 of the [CNI spec](https://github.com/containernetworking/cni/blob/master/SPEC.md). Support for new spec versions will be added shortly after each CNI release.
* `name`: Name of the network. This property can be set to any unique value.
* `type`: Name of the network plugin. This property should always be set to `azure-vnet`.
* `mode`: Operational mode. This field is optional. See the [operational modes](https://github.com/Azure/azure-container-networking/blob/master/docs/network.md) for more details.
* `master`: Name of the host network interface that will be used to connect containers to a VNET. This field is optional. If omitted, the plugin will automatically pick a suitable host network interface. Typically, the primary host interface name is `"Ethernet"` on Windows and `"eth0"` on Linux.
* `bridge`: Name of the bridge that will be used to connect containers to a VNET. This field is optional. If omitted, the plugin will automatically pick a unique name based on the master interface index.
* `logLevel`: Log verbosity. Valid values are `info` and `debug`. This field is optional. If omitted, the plugin will log at `info` level.

IPAM plugin
* `type`: Name of the IPAM plugin. This property should always be set to `azure-vnet-ipam`.
* `environment`: Name of the environment. Valid values are `azure` for [Azure](https://azure.microsoft.com) and `mas` for [Microsoft Azure Stack](https://azure.microsoft.com/en-us/overview/azure-stack/). This field is optional. The default value is `azure`.

You can create multiple network configuration files to connect containers to multiple networks.

Network configuration files are processed in lexical order during container creation, and in the reverse-lexical order during container deletion.

## Logs
Logs generated by `azure-vnet` plugin are available in `/var/log/azure-vnet.log` on Linux and `c:\cni\azure-vnet.log` on Windows.

Logs generated by `azure-vnet-ipam` plugin are available in `/var/log/azure-vnet.log` on Linux and `c:\cni\azure-vnet-ipam.log` on Windows.