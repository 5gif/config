# Config

This package consists of configurations files in json format and it's handlers. These files are required for the 5G NR Simulators. The handlers consists of function and method to read various configurations to golang structs.
All the configuration files for ITU, 3GPP and simulator are present here.

All the JSON files can be validated [here](https://jsonlint.com/) 


## Getting Started

- Clone the repository to the right directory.
```
>>$ cd $GOPATH
>>$GOPATH mkdir 'Github-Username'
>>$GOPATH cd 'Github-Username'
>>$GOPATH/'Github-Username' git clone "Clone Link"
```

### Pre-Requisites

- Install all go dependencies using the following command.
```
>>$GOPATH/'Github-Username' go get .
```

## Understand the JSON parameters

- The description of certain JSON parameters are given below.

### 3GPP
```
{
  "LAYOUTTYPE": 0, (0 OR 1) ->Gives the Type of Network Layout. In other words, it suggests if the layout is Indoor hotspot type or not.
}
```

### ITU
```
{
  "TrafficModel": 0, (0 means full buffer and 1 means lookup)
  "NumUEperCell": 30, (30 means that there are 10 UEs within a cell. In other words 10 UEs per TRxP)
  "INDOORRatio": 0.8, (0.8 implies 80% devices are deployed indoor and 20% are deployed outdoor)
  "BuildingTypeLoss": 0.8 (check variable name)
  "BandwidthMHz": 20, (It is the simulation Bandwidth of the system)
  "UEHeightout": 1.5, (check variable)
  "UEHeightin": 7.5,  (seed value)
  "INCARRatio":0, (Devices deployed outdoors are either in car or pedestrians. 0 means there is no user deployed within a car)
}
```

### How to use

``` go
import "github.com/5gif/config"
```

``` go
var nrcfg config.NRconfig  // loads the 3GPP related parameters for the RURAL Evaluation Config. A (of ITU-R/WP5D)
var path2file = "..path-to-3GPP_RMa_configA.json" // Add absolute path to the json file
vlib.LoadStructure(path2file, &nrcfg)
fmt.Println(nrcfg.FcGHz)

var ituconfg config.ITUconfig  // loads the 3GPP related parameters for the RURAL Evaluation Config. A (of ITU-R/WP5D)
var path2file = "..path-to-ITU_RMa_configA.json.json" // Add absolute path to the json file
vlib.LoadStructure(path2file, &ituconfg)
fmt.Println(ituconfig.FcGHz)

```

## Simplest way
``` go
config.SetDir("../json", "results")

var icfg config.ITUconfig
icfg.Read("ITU_RMa_configB.json")
fmt.Printf("Config = %#v", icfg)
icfg.Save()

icfg2 := config.ReadITUConfig("ITU_RMa_configA.json", "../json")
fmt.Printf("\n\nITU-R Method 2 Config = %#v\n", icfg2)
icfg2.Save()
```

## Built With

* [Golang](https://golang.org/) - JSON objects are handled using goalng structs and their corresponing handler methods.
* [JSON-to-Go](https://mholt.github.io/json-to-go/) struct convertor - It's an easy way to convert JSON to Go struct.

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* [PurpleBooth](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2) for this README template.
* [mholt](https://github.com/mholt) for a JSON-to-Go struct convertor 