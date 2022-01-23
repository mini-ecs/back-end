package virtlib

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-libvirt/socket/dialers"
	"github.com/mini-ecs/back-end/pkg/config"
	"github.com/mini-ecs/back-end/pkg/log"
	"math"
	"net"
	"strconv"
	"time"
)

/*
	该包提供的功能：
	1. Domain管理。
		- 获取Domain的IP √
		- 获取Domain的VNC端口 √
		- 制作快照（内存快照、磁盘快照）
		- 恢复快照
		- 删除快照
		- 列表快照
		- 获取Domain的信息，如CPU、内存资源状态
		- 创建Domain √
		- 删除Domain √
		- 重启Domain √
		- 关机Domain √
		- 启动Domain √
	2. 镜像文件管理-可替换（先实现本机存储,对用户透明）
		- iso等镜像
		- qcow2等磁盘镜像
	3. 节点管理（宿主机状态）
		- 查询宿主机的CPU、内存资源及其他状态
*/

type Lib struct {
	ip   net.IP
	port string
	con  *libvirt.Libvirt
}

func New(ip net.IP, port string) (*Lib, error) {
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.New("port input error")
	}
	return &Lib{
		ip:   ip,
		port: port,
	}, nil
}

func GetConnectedLib() *Lib {
	ip := net.ParseIP(config.GetConfig().NodeInfo.Ip)
	l, err := New(ip, strconv.Itoa(int(config.GetConfig().NodeInfo.Port)))
	if err != nil {
		panic("generate env error: " + err.Error())
	}
	err = l.Connect()
	if err != nil {
		panic(err)
	}
	return l
}
func (l *Lib) Connect() error {
	// todo: ip, port check
	c, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", l.ip.String(), l.port), 2*time.Second)
	if err != nil {
		return err
	}

	l.con = libvirt.NewWithDialer(dialers.NewAlreadyConnected(c))

	if err := l.con.ConnectToURI(libvirt.QEMUSystem); err != nil {
		return err
	}
	return nil
}

func (l *Lib) DisConnect() error {
	if l.con != nil {
		return l.con.Disconnect()
	}
	return errors.New("libvirt pointer is nil")
}

func (l *Lib) GetDomainByName(name string) (libvirt.Domain, error) {
	domain, err := l.con.DomainLookupByName(name)
	if err != nil {
		return libvirt.Domain{}, err
	}
	return domain, nil
}

var intToStateMap = map[libvirt.DomainState]string{
	0: "DomainNoState",
	1: "DomainRunning",
	2: "DomainBlocked",
	3: "DomainPaused",
	4: "DomainShutdown",
	5: "DomainShutoff",
	6: "DomainCrashed",
	7: "DomainPmSuspended",
}

func (l *Lib) GetDomainState(name string) (DomainState, error) {
	domain, err := l.GetDomainByName(name)
	if err != nil {
		return DomainState{}, err
	}
	state, maxMemory, mem, vcpu, cpu, err := l.con.DomainGetInfo(domain)
	if err != nil {
		return DomainState{}, err
	}
	ds := libvirt.DomainState(state)
	d := DomainState{
		State:     intToStateMap[ds],
		MaxMemory: uint32(maxMemory),
		Memory:    uint32(mem),
		VirCPU:    uint32(vcpu),
		CPUTime:   uint32(cpu),
	}
	return d, nil
}

func (l *Lib) GetDomainXML(name string) Domain {
	domain, err := l.GetDomainByName(name)
	if err != nil {
		return Domain{}
	}
	desc, err := l.con.DomainGetXMLDesc(domain, 0)
	if err != nil {
		panic(err)
	}
	var descDomain Domain
	err = xml.Unmarshal([]byte(desc), &descDomain)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", descDomain)
	return descDomain
}

func (l *Lib) getDomainMac(name string) string {
	d := l.GetDomainXML(name)
	return d.Devices.Interface.Mac.Address
}
func (l *Lib) getDomainBridgeName(name string) string {
	d := l.GetDomainXML(name)
	return d.Devices.Interface.Source.Bridge
}

func (l *Lib) GetDomainVNCPort(name string) string {
	d := l.GetDomainXML(name)
	return d.Devices.Graphics.Port
}

func (l *Lib) CreateDomain(opt DomainCreateOpt) error {
	xmls, err := xml.Marshal(opt)
	if err != nil {
		return err
	}
	domain, err := l.con.DomainDefineXML(string(xmls))
	if err != nil {
		return err
	}
	return l.StartDomain(domain.Name)
}

func (l *Lib) StartDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	err = l.con.DomainCreate(d)
	if err != nil {
		return err
	}
	return nil
}

// ShutdownDomain 关机
func (l *Lib) ShutdownDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	return l.con.DomainDestroy(d)
}

// DestroyDomain 删除
func (l *Lib) DestroyDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	return l.con.DomainDestroy(d)
}

// RebootDomain 重启
func (l *Lib) RebootDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	// libvirt.DomainRebootFlagValues
	return l.con.DomainReset(d, 0)
}

// SuspendDomain 挂起
func (l *Lib) SuspendDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	err = l.con.DomainSuspend(d)
	if err != nil {
		return err
	}
	return nil
}

// ResumeDomain 恢复（挂起的反动作）
func (l *Lib) ResumeDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	err = l.con.DomainResume(d)
	if err != nil {
		return err
	}
	return nil
}

func (l *Lib) UnDefineDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	return l.con.DomainUndefine(d)
}

// GetDomainIP error
func (l *Lib) GetDomainIP(name string) {
	mac := l.getDomainBridgeName(name)
	inter, err := l.con.InterfaceLookupByName(mac)
	if err != nil {
		panic(err)
	}
	desc, err := l.con.InterfaceGetXMLDesc(inter, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(desc)
}

// GetDomainIPAddress 不知道为什么会返回一个数组
func (l *Lib) GetDomainIPAddress(name string) (string, error) {
	log.GetGlobalLogger().Infof("query %+v's ip", name)

	d, err := l.GetDomainByName(name)
	if err != nil {
		return "", err
	}

	addresses, err := l.con.DomainInterfaceAddresses(d, uint32(libvirt.DomainInterfaceAddressesSrcLease), 0)
	if err != nil {
		return "", err
	}
	log.GetGlobalLogger().Infof("get ip address: %+v", addresses)
	if len(addresses) > 0 && len(addresses[0].Addrs) > 0 {
		return addresses[0].Addrs[0].Addr, nil
	}
	return "", nil
}

// CreateSnapshot
// 		libvirt不支持对于external snapshot的管理操作：恢复快照、删除快照会有如下报错:
//		unsupported configuration: deletion of 1 external disk snapshots not supported yet.
//		而如果在创建快照时flag为libvirt.DomainSnapshotCreateDiskOnly（libvirt.DomainSnapshotCreateFlags）时，
//		快照会被保存到外部，即快照是external的。这就需要手动来管理这些external的快照。
//
//		创建：将 libvirt.DomainSnapshotCreateFlags 传入函数，即可在Domain对应的镜像文件所在的目录中生成一个镜像文件，
//			镜像文件的后缀是snapshot的名称。
//		回滚：只能调用带 libvirt.DomainSnapshotDeleteMetadataOnly flag的函数，删除镜像元数据，然后手动将xml中的镜像
//			文件修改为parent的快照，并删除要删除的镜像文件。
//
//		手工操作可参考： https://fabianlee.org/2021/01/10/kvm-creating-and-reverting-libvirt-external-snapshots/
func (l *Lib) CreateSnapshot(name string, opt DomainSnapshot) error {
	domain, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	xmlResult, err := xml.Marshal(opt)
	if err != nil {
		return err
	}
	//libvirt.DomainSnapshotCreateFlags
	diskOnly := libvirt.DomainSnapshotCreateDiskOnly // 只保存磁盘则使用该flag
	createXML, err := l.con.DomainSnapshotCreateXML(domain, string(xmlResult), uint32(diskOnly))
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", createXML)
	return nil
}

func (l *Lib) RevertSnapshot(name string) error {
	domain, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	snap, err := l.con.DomainSnapshotCurrent(domain, 0)
	if err != nil {
		return err
	}
	err = l.con.DomainRevertToSnapshot(snap, 0)
	if err != nil {
		return err
	}
	return nil
}

func (l *Lib) GetSnapshotByName(domain, snapshotName string) (libvirt.DomainSnapshot, error) {
	d, err := l.GetDomainByName(domain)
	if err != nil {
		return libvirt.DomainSnapshot{}, err
	}
	snapshot, err := l.con.DomainSnapshotLookupByName(d, snapshotName, 0)
	if err != nil {
		return libvirt.DomainSnapshot{}, err
	}
	return snapshot, nil
}

// DeleteSnapshot
// 		args: libvirt.DomainSnapshotDeleteChildren 此快照和任何后代快照都将被删除
// 		libvirt.DomainSnapshotDeleteChildrenOnly 删除任何后代快照，但保留此快照
// 		libvirt.DomainSnapshotDeleteMetadataOnly 任何由libvirt追踪的快照元数据都会被移除，同时保持快照内容不变；
// 		如果管理程序不需要任何libvirt元数据来追踪快照，那么这个标志会被默默地忽略。
//		=> 可以删除external快照的metadata，但不删除文件，需要手动管理文件
//
func (l *Lib) DeleteSnapshot(name, snapshotName string, args libvirt.DomainSnapshotDeleteFlags) error {
	snapshot, err := l.GetSnapshotByName(name, snapshotName)
	if err != nil {
		return err
	}
	return l.con.DomainSnapshotDelete(snapshot, args)
}

func (l *Lib) ListSnapshots(name string) ([]libvirt.DomainSnapshot, error) {
	domain, err := l.GetDomainByName(name)
	if err != nil {
		return nil, err
	}
	snapshots, _, err := l.con.DomainListAllSnapshots(domain, math.MaxInt8, 0)
	if err != nil {
		return nil, err
	}
	return snapshots, nil
}

func (l *Lib) PullAllSnapshots(domainName string) error {
	domain, err := l.GetDomainByName(domainName)
	if err != nil {
		return err
	}
	// 确认是否有快照，没有快照，则说明已经只有一份镜像了
	_, err = l.GetCurrentSnapshot(domainName)
	if err != nil {
		return nil
	}

	//d := l.GetDomainXML(domainName)
	//err := l.con.DomainBlockPull(domain, d.Devices.Disk[0].Source.File, 0, 0)
	return l.con.DomainBlockPull(domain, "hda", 0, 0)
}

func (l *Lib) GetCurrentSnapshot(domainName string) (libvirt.DomainSnapshot, error) {
	domain, err := l.GetDomainByName(domainName)
	if err != nil {
		return libvirt.DomainSnapshot{}, err
	}
	return l.con.DomainSnapshotCurrent(domain, 0)
}

func (l *Lib) GetAllInterfaces() {
	interfaces, _, _ := l.con.ConnectListAllInterfaces(10, 0)
	//l.con.DomainInterfaceAddresses()
	listInterfaces, err := l.con.ConnectListInterfaces(10)
	if err != nil {
		return
	}
	d, err := l.GetDomainByName("ubuntu")
	addresses, err := l.con.DomainInterfaceAddresses(d, uint32(libvirt.DomainInterfaceAddressesSrcLease), 0)
	if err != nil {
		panic(err)
	}
	//l.con.InterfaceLookupByMacString()
	//l.con.DomainSnapshotCurrent()
	fmt.Printf("%+v \n\n %+v\n\n %+v", interfaces, listInterfaces, addresses)
}
