package config

// import (
// 	"os"
// 	"path/filepath"

// 	log "github.com/Sirupsen/logrus"
// 	"github.com/spf13/viper"
// 	"github.com/wiless/vlib"
// )

// NRconfig structure
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
