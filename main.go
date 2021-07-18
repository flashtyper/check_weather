package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "encoding/json"
)



func main() {
        arg1 := os.Args[1]
        city := os.Args[2]
        arg2 := os.Args[3]
        token := os.Args[4]

        if (arg1 == "-c" && len(city) > 0 && arg2 == "-t" && len(token) > 0 ){
                get(city, token)
        } else {
                help()
        }
}


func get(city string, token string) {
  http_response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric&APPID=" + token)

  if err != nil {
    fmt.Print(err.Error())
    os.Exit(2)
  }
  if (http_response.StatusCode == 200) {
    responseData, err := ioutil.ReadAll(http_response.Body)

    if err != nil {
      fmt.Print(err.Error())
      os.Exit(2)
    }

    //convert structs
    var weatherJSON WeatherJSON
    var weather Weather

    json.Unmarshal(responseData, &weatherJSON)
    weather.Temperature = weatherJSON.Main.Temp
    weather.Pressure = weatherJSON.Main.Pressure
    weather.Humidity = weatherJSON.Main.Humidity
    weather.Wind = weatherJSON.Wind.Speed
    weather.Cloudiness = weatherJSON.Clouds.All
    weather.Rain = weatherJSON.Rain.Hour
    weather.Snow = weatherJSON.Snow.Hour
    fmt.Printf("Current weather in " + city + " | temp=%f pressure=%d humidity=%d wind=%f clouds=%d rain=%f snow=%f", weather.Temperature, weather.Pressure, weather.Humidity, weather.Wind, weather.Cloudiness, weather.Rain, weather.Snow)
    os.Exit(0)
  } else {
    fmt.Println("ERROR:", http_response.StatusCode, http.StatusText(http_response.StatusCode))
    os.Exit(1)
  }
}

func help() {
  fmt.Println("--------------------------------")
  fmt.Println("Usage:")
  fmt.Println("./check_weather -c [City] -t [Token]")
  fmt.Println("--------------------------------\n")
}
