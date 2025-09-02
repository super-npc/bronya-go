package register

import (
	"fmt"
	"time"

	"github.com/super-npc/bronya-go/src/framework/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var EtcdClient *clientv3.Client

func InitEtcd() {
	cnf := conf.Settings.Etcd
	// 1. 建立连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%d", cnf.Host, cnf.Port)},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	EtcdClient = cli
	//defer func(cli *clientv3.Client) {
	//	err := cli.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(cli)
}
