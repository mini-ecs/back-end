package virtlib

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestConnectEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.249.46.250:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Println(err)
	}
	defer cli.Close()
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//resp, err := cli.Get(ctx, "test")
	//resp, err := cli.Put(context.TODO(), "/a/A/B", "123123312", clientv3.WithPrevKV())
	resp, err := cli.Get(context.TODO(), "/a/A/B", clientv3.WithPrefix())
	//cancel()
	if err != nil {
		fmt.Println(err)
	}
	// use the response
	fmt.Printf("%+v", resp)

}
