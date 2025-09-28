package discovery

import (
    "log"
    "strconv"
    "time"
    "github.com/hudl/fargo"
    "customer-service/pkg/config"
)

var eurekaConn fargo.EurekaConnection

func InitEureka() {
    eurekaURL := config.GetString("eureka.client.serviceUrl.defaultZone")
    servicePort := config.GetInt("server.port")
    if servicePort == 0 {
        servicePort = 8081
    }

    eurekaConn = fargo.NewConn(eurekaURL)

    instance := &fargo.Instance{
        InstanceId:       "customer-service",
        HostName:         "localhost",
        App:              "CUSTOMER-SERVICE",
        IPAddr:           "127.0.0.1",
        VipAddress:       "customer-service",
        SecureVipAddress: "customer-service",
        Status:           fargo.UP,
        Port:             servicePort,
        SecurePort:       servicePort,
        HomePageUrl:      "http://localhost:" + strconv.Itoa(servicePort),
        StatusPageUrl:    "http://localhost:" + strconv.Itoa(servicePort) + "/info",
        HealthCheckUrl:   "http://localhost:" + strconv.Itoa(servicePort) + "/health",
        DataCenterInfo: fargo.DataCenterInfo{
            Name:  "MyOwn",
            Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
        },
        LeaseInfo: fargo.LeaseInfo{
            RenewalIntervalInSecs: 30,
            DurationInSecs:       90,
        },
    }

    if err := eurekaConn.RegisterInstance(instance); err != nil {
        log.Fatal("❌ Failed to register with Eureka:", err)
    }
    log.Println("✅ Service registered with Eureka")

    go startHeartbeat(instance)
}

func startHeartbeat(instance *fargo.Instance) {
    ticker := time.NewTicker(30 * time.Second)
    for {
        <-ticker.C
        if err := eurekaConn.HeartBeatInstance(instance); err != nil {
            log.Printf("⚠️ Failed to send heartbeat: %v", err)
        }
    }
}

func DeregisterService() {
    instance := &fargo.Instance{
        App:        "CUSTOMER-SERVICE",
        InstanceId: "customer-service",
    }
    if err := eurekaConn.DeregisterInstance(instance); err != nil {
        log.Printf("⚠️ Failed to deregister from Eureka: %v", err)
    } else {
        log.Println("✅ Service deregistered from Eureka")
    }
}
