package config

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "github.com/spf13/viper"
)

// Structures pour analyser JSON Config Server
type PropertySource struct {
    Name   string                 `json:"name"`
    Source map[string]interface{} `json:"source"`
}

type ConfigResponse struct {
    Name            string           `json:"name"`
    PropertySources []PropertySource `json:"propertySources"`
}

// Charger la config depuis Config Server
func LoadConfigFromServer(configServerURL, appName, profile string) {
    url := fmt.Sprintf("%s/%s/%s", configServerURL, appName, profile)
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("❌ Impossible de se connecter au Config Server: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("❌ Impossible de lire la réponse du Config Server: %v", err)
    }

    var configResp ConfigResponse
    if err := json.Unmarshal(body, &configResp); err != nil {
        log.Fatalf("❌ Impossible de parser JSON Config Server: %v", err)
    }

    // Charger toutes les propriétés dans Viper
    viper.AutomaticEnv()
    for _, ps := range configResp.PropertySources {
        for k, v := range ps.Source {
            viper.Set(k, v)
        }
    }

    log.Println("✅ Configuration chargée depuis Config Server :", configServerURL)
}

// Utilitaires pour récupérer les valeurs
func GetString(key string) string {
    return viper.GetString(key)
}

func GetInt(key string) int {
    return viper.GetInt(key)
}
