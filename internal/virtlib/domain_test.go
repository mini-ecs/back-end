package virtlib

import (
	"encoding/xml"
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

var domainName = "123"

func GenerateEnv() *Lib {
	ip := net.ParseIP("10.249.46.250")
	l, err := New(ip, "16509")
	if err != nil {
		panic("generate env error: " + err.Error())
	}
	err = l.Connect()
	if err != nil {
		panic(err)
	}
	return l
}
func TestLib_GetInterface(t *testing.T) {
	l := GenerateEnv()
	l.GetAllInterfaces()
}

func TestLib_GetDomainIP(t *testing.T) {
	l := GenerateEnv()
	i1, err := l.GetDomainIPAddress(domainName)
	assert.Equal(t, err, nil)
	fmt.Printf("%+v\n", i1)
}

func TestLib_NewDomainOpt(t *testing.T) {
	const opt = `<domain type='kvm'>
  <name>ubuntu</name>
  <uuid>c7a5fdbd-cdaf-9455-926a-d65c16db1809</uuid>
  <memory unit="GiB">4</memory>
  <currentMemory unit="GiB">4</currentMemory>
  <vcpu>4</vcpu>
  <os>
    <type arch='x86_64' machine='pc'>hvm</type>
    <boot dev='hd'/>
    <boot dev='cdrom'/>
  </os>
  <devices>
    <emulator>/usr/bin/qemu-system-x86_64</emulator>
    <disk type='file' device='cdrom'>
      <source file='/home/fangaoyang/work/ubuntu-20.04.3-desktop-amd64.iso'/>
      <target dev='hdc'/>
      <readonly/>
    </disk>
    <disk type='file' device='disk'>
      <driver name="qemu" type="qcow2" />
      <source file='/home/fangaoyang/work/desktop.qcow2'/>
      <target dev='hda'/>
    </disk>
    <interface type='network'>
      <source network='default'/>
    </interface>
    <graphics type="vnc" port="-1" listen="0.0.0.0" passwd="123" keymap="en-us"/>
    <serial type="pty">
    	<target port="0" />
	</serial>
  	<console type="pty">
   		<target type="serial" port="0" />
  	</console>
  </devices>
</domain>`
	var d Domain
	err := xml.Unmarshal([]byte(opt), &d)
	if err != nil {
		t.Error(err)
	}

	res, err := xml.Marshal(DefaultCreateDomainOpt)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", string(res))

	res2, _ := xml.Marshal(d)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", string(res2))

	assert.Equal(t, res, res2)
}

// 经过测试，virsh destory $domainName 只销毁运行时数据及其元数据。相当于数据库里已经没有了$domainName这个vm的信息，但是其根本文件还存在。
// 那么"恢复到某个快照"这个操作有两种方法：
//		1. 可以先destory这个domain，再根据要恢复到的快照的文件来重新创建一个同名的domain，那么就相当于切换了快照。不过这样的话会导致元数据丢失，
// 		通过 virsh snapshot-list domainName 中无法看到当前快照的父快照了，但是使用 qemu-img info 命令可以追溯父快照。在qemu虚拟机运行时，
//		会自行寻找父快照来找信息。如下所示：表示snap-1237.qcow2的父快照是/home/fangaoyang/work/e9216248-3cbc-4a01-9a67-496dfd2cc4ec.forth
//			$ sudo qemu-img info snap-1237.qcow2
//				image: snap-1237.qcow2
//				file format: qcow2
//				virtual size: 40 GiB (42949672960 bytes)
//				disk size: 5.33 GiB
//				cluster_size: 65536
//				backing file: /home/fangaoyang/work/e9216248-3cbc-4a01-9a67-496dfd2cc4ec.forth
//				backing file format: qcow2
//				Snapshot list:
//				ID        TAG                     VM SIZE                DATE       VM CLOCK
//				1         five                   1.45 GiB 2022-01-21 14:51:40   00:22:45.908
//				Format specific information:
//				compat: 1.1
//				lazy refcounts: false
//				refcount bits: 16
//				corrupt: false
//		虽然该方法可以使全量快照和硬盘快照的行为逻辑一致，但该方法的缺点也比较明显，会造成元数据的丢失，同时需要手动管理快照文件。例如恢复到某
//		个快照、删除了某个domain时，需要有对应的正确行为逻辑。
//		2. 第二种方法就是各管各的，internal直接使用libvirt api的逻辑，external就自己管理。其中，external的可以修改其xml文件，来调整镜像，尚未测试。
//		发现重大bug：之前使用的创建Domain的函数一直是DomainCreateXML，导致机器被Destroy之后，其xml就不见了，很让人苦恼。使用了DomainDefineXML后就不存在该问题了
func TestLib_CreateDomain(t *testing.T) {
	l := GenerateEnv()
	d := DefaultCreateDomainOpt
	u, _ := uuid.NewUUID()
	d.Uuid = u.String()
	d.Name = domainName
	d.Devices.Disk[1].Source.File = "/home/fangaoyang/work/e9216248-3cbc-4a01-9a67-496dfd2cc4ec.first"
	fmt.Printf("%+v\n", d)
	err := l.CreateDomain(d)
	assert.Equal(t, err, nil)
}

func TestLib_ShutdownDomain(t *testing.T) {
	l := GenerateEnv()
	err := l.ShutdownDomain(domainName)
	assert.Equal(t, err, nil)
}

func TestLib_RebootDomain(t *testing.T) {
	l := GenerateEnv()
	err := l.RebootDomain(domainName)
	assert.Equal(t, err, nil)
}
func TestLib_SuspendDomain(t *testing.T) {
	l := GenerateEnv()
	err := l.SuspendDomain(domainName)
	assert.Equal(t, err, nil)
}
func TestLib_ResumeDomain(t *testing.T) {
	l := GenerateEnv()
	err := l.ResumeDomain(domainName)
	assert.Equal(t, err, nil)
}
func TestLib_DestroyDomain(t *testing.T) {
	l := GenerateEnv()
	err := l.DestroyDomain(domainName)
	assert.Equal(t, err, nil)
}

func TestLib_CreateSnapshot(t *testing.T) {
	l := GenerateEnv()
	opt := DomainSnapshot{
		Name: "forth",
	}
	err := l.CreateSnapshot(domainName, opt)
	assert.Equal(t, err, nil)

}
func TestLib_ListSnapshots(t *testing.T) {
	l := GenerateEnv()
	list, err := l.ListSnapshots(domainName)
	fmt.Printf("%+v", list)
	assert.Equal(t, err, nil)
}
func TestLib_DeleteSnapshot(t *testing.T) {
	l := GenerateEnv()
	list, err := l.ListSnapshots(domainName)
	fmt.Printf("%+v", list)
	assert.Equal(t, err, nil)
	assert.NotNil(t, list)
	assert.Greater(t, len(list), 0, "please make more snapshots")

	snapshot := list[0]
	err = l.DeleteSnapshot(domainName, snapshot.Name, libvirt.DomainSnapshotDeleteChildren)
	assert.Nil(t, err)
}
