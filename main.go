package main
import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)

type Config struct {
	Url             string `mapstructure:"url"`
	Apikey string `mapstructure:"apikey"`
}
var AppConfig *Config
func LoadAppConfig(){
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}

type Response struct {
    Name    string    `json:"name"`
	Main struct {
	Temp float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
     } `json:"main"`

}

func main(){
	LoadAppConfig()
	var city string = "London"
	response, err := http.Get(fmt.Sprintf("%s?q=%s&appid=%s&units=metric",AppConfig.Url,city,AppConfig.Apikey))
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }
	responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Name)
	fmt.Println("Pressure: ",responseObject.Main.Pressure)
	fmt.Println("Temperature: ",responseObject.Main.Temp)
	fmt.Println("Humidity: ",responseObject.Main.Humidity)

	
	
	
}
