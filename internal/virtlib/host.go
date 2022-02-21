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

func (l *Lib) GetDomMemStats(domName string) ([]libvirt.DomainMemoryStat, error) {
	dom, err := l.GetDomainByName(domName)

	if err != nil {
		return nil, err
	}
	stats, err := l.con.DomainMemoryStats(dom, 100, 0)
	return stats, err
}

func (l *Lib) GetDomMemUsage(domName string) (float64, error) {
	stats, err := l.GetDomMemStats(domName)
	var total, usable uint64
	for _, stat := range stats {
		if stat.Tag == 5 {
			total = stat.Val
		}
		if stat.Tag == 8 {
			usable = stat.Val
		}
	}
	return float64(total-usable) / float64(total), err
}
