package virtlib

import (
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"time"
)

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

func (l *Lib) GetDomDiskUsage(domName string) (float64, error) {
	path := l.GetDomainDiskPath(domName)
	dom, err := l.GetDomainByName(domName)
	if err != nil {
		return 0, err
	}
	alloc, capacity, _, err := l.con.DomainGetBlockInfo(dom, path, 0)
	if err != nil {
		return 0, err
	}
	return float64(alloc) / float64(capacity), nil
}

func (l *Lib) GetDomCPUTime(domName string) ([]libvirt.TypedParam, error) {
	dom, err := l.GetDomainByName(domName)
	if err != nil {
		return nil, err
	}
	stats, _, err := l.con.DomainGetCPUStats(dom, 3, -1, 1, 0)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func (l *Lib) GetDomCPUUsage(domName string, interval int) (float64, error) {
	t1, err := l.GetDomCPUTime(domName)
	if err != nil {
		return 0.0, err
	}
	var cpu1, cpu2 uint64
	for _, v := range t1 {
		if v.Field == "cpu_time" {
			cpu1 = v.Value.I.(uint64)
		}
	}
	time.Sleep(time.Duration(interval) * time.Second)
	t2, err := l.GetDomCPUTime(domName)
	if err != nil {
		return 0.0, err
	}
	for _, v := range t2 {
		if v.Field == "cpu_time" {
			cpu2 = v.Value.I.(uint64)
		}
	}
	fmt.Println(cpu1, "  ", cpu2)
	ret := float64(cpu2-cpu1) / float64(interval)
	return ret / 10e9, nil
}
