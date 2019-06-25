package config

type NRconfig struct {
	ENV                     string  `json:"ENV"`
	SCENARIO                string  `json:"SCENARIO"`
	CONFIG                  string  `json:"CONFIG"`
	LAYOUTTYPE              int     `json:"LAYOUTTYPE"`
	ISD                     int     `json:"ISD"`
	FcGHz                   float64 `json:"FcGHz"`
	BandwidthMHz            float64 `json:"BandwidthMHz"`
	SCSKHz                  int     `json:"SCSKHz"`
	Duplexity               string  `json:"Duplexity"`
	AntennaScheme           string  `json:"AntennaScheme"`
	ReliabiltyAntennaScheme string  `json:"ReliabiltyAntennaScheme"`
	BSAntennaConfig         []int   `json:"BSAntennaConfig"`
	UEAntennaConfig         string  `json:"UEAntennaConfig"`
	FrameStructure          string  `json:"FrameStructure"`
	ChannelModel            string  `json:"ChannelModel"`
	NumTRxP                 string  `json:"NumTRxP"`
	MobilityClass           int     `json:"MobilityClass"`
}

type ITUconfig struct {
	ENV                 string  `json:"ENV"`
	SCENARIO            string  `json:"SCENARIO"`
	CONFIG              string  `json:"CONFIG"`
	LAYOUTTYPE          int     `json:"LAYOUTTYPE"`
	BSantennaType       string  `json:"BSantennaType"`
	UEantennaType       string  `json:"UEantennaType"`
	FcGHz               int     `json:"FcGHz"`
	Duplexity           string  `json:"Duplexity"`
	BSHeight            int     `json:"BSHeight"`
	UEHeightout         float64 `json:"UEHeightout"`
	UEHeightin          float64 `json:"UEHeightin"`
	BSTxPowerDbm        int     `json:"BSTxPowerDbm"`
	BandwidthMHz        int     `json:"BandwidthMHz"`
	UETxDbm             int     `json:"UETxDbm"`
	BuildingTypeLoss    float64 `json:"BuildingTypeLoss"`
	ISD                 int     `json:"ISD"`
	NumBSelements       int     `json:"NumBSelements"`
	NumUEelements       int     `json:"NumUEelements"`
	INDOORRatio         float64 `json:"INDOORRatio"`
	IndoorSpeed         int     `json:"IndoorSpeed"`
	Outdoorspeed        int     `json:"Outdoorspeed"`
	BSNoiseFigureDb     int     `json:"BSNoiseFigureDb"`
	UENoiseFigureDb     int     `json:"UENoiseFigureDb"`
	BSAntennaEleGainDbi int     `json:"BSAntennaEleGainDbi"`
	UEAntennaEleGainDbi int     `json:"UEAntennaEleGainDbi"`
	N0                  int     `json:"N0"`
	TrafficModel        int     `json:"TrafficModel"`
	NumUEperCell        int     `json:"NumUEperCell"`
}
