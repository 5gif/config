# config
All the configuration files for ITU, 3GPP and simulator will be here

Must be able to validate all the JSON files [here](https://jsonlint.com/) 


## 3GPP
```
{
  "LAYOUTTYPE": 0, (0 OR 1) ->Gives the Type of Network Layout. In other words, it suggests if the layout is Indoor hotspot type or not.
}
```


## ITU
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

## How to use 

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

### Credits
Used https://mholt.github.io/json-to-go/ . to create `golang` struct
