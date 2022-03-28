/*
 * @Author: FANG
 * @Date: 2022-03-28 14:01:30
 */
package main

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
)

const (
	consulAgentAddress = "agent ip:8500"
)

// 从consul中发现服务
func ConsulFindServer() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAgentAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 获取指定service
	service, _, err := client.Agent().Service("337", nil)
	if err == nil {
		fmt.Println(service.Address)
		fmt.Println(service.Port)
	}

	// 只获取健康的service
	// serviceHealthy, _, err := client.Health().Service("service337", "", true, nil)
	// if err == nil {
	// 	fmt.Println(serviceHealthy[0].Service.Address)
	// }

}

func main() {
	ConsulFindServer()
}
