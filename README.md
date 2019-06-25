# config
All the configuration files for ITU, 3GPP and simulator will be here

Must be able to validate all the JSON files [here](https://jsonlint.com/) 


## 3GPP
```
{
  "LAYOUTTYPE": 0, (0 OR 1) ->nETWORK lAYOUT
  "ISD": 0, (1 means 6km o/w not mentioned)
}
```


## ITU
```
{
  "TrafficModel": 0, (0 means full buffer and 1 means lookup)
  "NumUEperCell": 10, (UE/TRxP)
  "INDOORRatio": 0.8, (device deployment)
  "BuildingTypeLoss": 0.8 (check variable name)
  "BandwidthMHz": 20, (same parameters as 3GPP)
  "UEHeightout": 1.5, (check variable)
  "UEHeightin": 7.5,  (seed value)
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

### Credits
Used https://mholt.github.io/json-to-go/ . to create `golang` struct
