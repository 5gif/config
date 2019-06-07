# config
All the configuration files for ITU, 3GPP and simulator will be here

Must be able to valide all the JSON files [here](https://jsonlint.com/) 


## 3GPP
```
{
"LAYOUTTYPE": 0, (0 OR 1) ->nETWORK lAYOUT
"ISD": 0, (1 means 6km o/w not mentioned)
}
```


## ITU-R
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

### Credits
Used https://mholt.github.io/json-to-go/ . to create `golang` struct
