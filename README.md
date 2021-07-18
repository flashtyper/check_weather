# check_weather
Short Icinga Plugin which checks the current weather in a user defined city

## Usage: 

```bash
./weather -c [CITY] -t [API TOKEN]
```

## Example
```
lukas@ubuntu:~/wetter$ ./wetter -c Regensburg -t foobar
Current weather in Regensburg | temp=25.500000 pressure=1017 humidity=60 wind=3.050000 clouds=1 rain=0.000000 snow=0.000000l
```

## Installation
Create a account at https://openweathermap.org/api and subscribe to the "Current Weather Data" API. 
Set up your icinga: 
```bash
object CheckCommand "weather" {
  command = [ CustomPluginDir + "/check_weather" ]
  arguments = {
    "-c" = "$city$"
    "-t" = "YOUR TOKEN"
  }
}
apply Service "Regensburg" {
  import "generic-service"
  check_command = "weather"
  vars.city = "Regensburg"
  assign where host.name == "Wetter"
}
object Host "Wetter" {
  import "generic-host"
  address = "127.0.0.1"
}
```
I recommend using a seperate host for this. 

After that you can create a grafana dashboard if you want to: 
![grafik](https://user-images.githubusercontent.com/83031404/126076177-078ce8cc-3252-4f20-b183-ed89d4038115.png)
