package main

import (
	"github.com/hashicorp/consul/api"
)

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "124.222.97.37:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成对应的检测对象
	check := &api.AgentServiceCheck{
		//GRPC:                           "124.222.97.37:8999/health",
		HTTP:                           "http://124.222.97.37:8555/",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Address: address,
		Port:    port,
		Check:   check,
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func AllServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "124.222.97.37:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	services, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}

	for key, value := range services {
		println(key, value)
	}
}

func main() {
	err := Register("124.222.97.37", 8500, "user-web", []string{"mxshop"}, "user-web")
	if err != nil {
		return
	}
	AllServices()
}
