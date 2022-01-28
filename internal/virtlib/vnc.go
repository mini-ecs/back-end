package virtlib

import (
	"errors"
	"fmt"
	"math/rand"
	"os/exec"
	"sync"
	"time"
)

var portPool map[int]int
var mutex sync.Mutex

func init() {
	portPool = make(map[int]int, 0)
}

func ProxyVNCToWebSocket(vncPort int) (int, error) {
	// 循环2000次，如果没有则报错
	var httpPort = -1
	for cnt := 0; cnt < 2000; cnt++ {
		httpPort = rand.Int()%(65535-5000) + 5000
		if _, ok := portPool[httpPort]; !ok && !PortInUse(httpPort) {
			break
		}
	}
	if httpPort == -1 {
		return -1, errors.New("port pool is full")
	}
	if err := establishVNCtoHTTP(vncPort, httpPort); err != nil {
		return -1, err
	}

	return httpPort, nil
}
func establishVNCtoHTTP(vncPort int, httpPort int) error {
	mutex.Lock()
	portPool[httpPort] = vncPort
	mutex.Unlock()
	cmd := exec.Command("sh", "-c", fmt.Sprintf(
		"novnc_proxy --listen %v --vnc localhost:%v",
		httpPort,
		vncPort,
	))

	if err := cmd.Start(); err != nil {
		mutex.Lock()
		delete(portPool, httpPort)
		mutex.Unlock()
		return err
	}
	go func() {
		// two hour
		time.Sleep(2 * 60 * 60 * time.Second)
		cmd.Process.Kill()
		mutex.Lock()
		delete(portPool, httpPort)
		mutex.Unlock()
	}()
	return nil
}
func PortInUse(port int) bool {
	checkStatement := fmt.Sprintf("lsof -i:%d ", port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if len(output) > 0 {
		return true
	}
	return false
}
