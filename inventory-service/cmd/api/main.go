package main

import (
    "inventory-service/internal/app/router"
    "inventory-service/pkg/discovery"
    "inventory-service/pkg/utils"
    "inventory-service/pkg/config"
    "os"
    "os/signal"
    "syscall"
    "strconv"

)

func main() {
    
    // 1️⃣ Charger la config depuis Spring Cloud Config Server
    configServerURL := "http://localhost:9999" // Config Server
    appName := "inventory-service"
    profile := "default"
    config.LoadConfigFromServer(configServerURL, appName, profile)

    // 2️⃣ Initialiser DB
    utils.InitDBFromConfig()

    // 3️⃣ Initialiser Eureka
    discovery.InitEureka()

    // 4️⃣ Setup graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    // 5️⃣ Lancer serveur HTTP
    r := router.SetupRouter()
    go r.Run(":" + strconv.Itoa(config.GetInt("server.port")))

    <-sigChan

    // 6️⃣ Désinscription Eureka
    discovery.DeregisterService()
}
