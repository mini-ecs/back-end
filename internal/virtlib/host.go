package virtlib

import "github.com/digitalocean/go-libvirt"

//func (l *Lib) GetCPUStat() {
//	l.con.NodeGetInfo()
//	l.con.ConnectGetCapabilities()
//	l.con.ConnectGetDomainCapabilities()
//}

func (l *Lib) GetCapabilities() (string, error) {
	capabilities, err := l.con.ConnectGetCapabilities()
	return capabilities, err
}

func (l *Lib) GetNodeCPUStats() ([]libvirt.NodeGetCPUStats, error) {
	stats, _, err := l.con.NodeGetCPUStats(int32(libvirt.NodeCPUStatsAllCpus), 0, 0)
	return stats, err
}

func (l *Lib) GetNodeMemStats() ([]libvirt.NodeGetMemoryStats, error) {
	stats, _, err := l.con.NodeGetMemoryStats(int32(libvirt.NodeMemoryStatsAllCells), 0, 0)
	return stats, err
}
