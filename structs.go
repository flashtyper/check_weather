package main

type WeatherJSON struct {
        Coord struct {
                Lon float64 `json:"lon"`
                Lat float64 `json:"lat"`
        } `json:"coord"`
        Weather []struct {
                ID          int    `json:"id"`
                Main        string `json:"main"`
                Description string `json:"description"`
                Icon        string `json:"icon"`
        } `json:"weather"`
        Base string `json:"base"`
        Main struct {
                Temp      float64 `json:"temp"`
                FeelsLike float64 `json:"feels_like"`
                TempMin   float64 `json:"temp_min"`
                TempMax   float64 `json:"temp_max"`
                Pressure  int     `json:"pressure"`
                Humidity  int     `json:"humidity"`
        } `json:"main"`
        Visibility int `json:"visibility"`
        Wind       struct {
                Speed float64 `json:"speed"`
                Deg   int     `json:"deg"`
                Gust  float64 `json:"gust"`
        } `json:"wind"`
        Rain struct {
                Hour float64 `json:"1h"`
        }
        Snow struct {
                Hour float64 `json:"1h"`
        }
        Clouds struct {
                All int `json:"all"`
        } `json:"clouds"`
        Dt  int `json:"dt"`
        Sys struct {
                Type    int    `json:"type"`
                ID      int    `json:"id"`
                Country string `json:"country"`
                Sunrise int    `json:"sunrise"`
                Sunset  int    `json:"sunset"`
        } `json:"sys"`
        Timezone int    `json:"timezone"`
        ID       int    `json:"id"`
        Name     string `json:"name"`
        Cod      int    `json:"cod"`
}

type Weather struct {
        Temperature float64
        Pressure int
        Humidity int
        Wind float64
        Cloudiness int
        Rain float64
        Snow float64
}