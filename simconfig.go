package config

import (
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

// SIMconfig ...
type SIMconfig struct {
	SCENARIO       string `json:"SCENARIO"`
	SimulationTime int    `json:"SimulationTime"`
	FrameType      int    `json:"FrameType"`
	SlotType       int    `json:"SlotType"`
	UEperSlot      int    `json:"UEperSlot"`
	Fname          string `json:"fname"`
	ActiveBSCells  int    `json:"ActiveBSCells"`
	ActiveUECells  int    `json:"ActiveUECells"`
	Extended       bool   `json:"Extended"`
	ForceAllLOS    bool   `json:"ForceAllLOS"`
	ShadowLoss     bool   `json:"ShadowLoss"`
	LogInfo        bool   `json:"LogInfo"`
	UEcells        []int  `json:"UEcells"`
	BScells        []int  `json:"BScells"`
	TrueCells      int    `json:"TrueCells"`
	SchedulerType  int    `json:"SchedulerType"`
	IsDownLink     bool   `json:"IsDownLink"`
}

//SetDefaults loads the default values for the simulation
func (i *SIMconfig) SetDefaults() {

	i.ActiveBSCells = -1 // Default all the cells are active
	i.ActiveUECells = -1 // UEs are dropped in all the cells
	i.Extended = false
	i.ForceAllLOS = false
	i.ShadowLoss = true
	i.LogInfo = false
	i.UEcells = []int{0, 10}
	i.BScells = []int{0, 1, 2}
	i.TrueCells = -1 // Default to all the cells
	i.IsDownLink = true

}

// Save ...
func (i *SIMconfig) Save() {
	CurrDIR, _ := os.Getwd()
	t := time.Now()
	year, month, day := t.Date()
	root, _ := os.Getwd()
	OutDIR := root + "/" + OutDIR + "/" + strconv.Itoa(day) + "_" + strconv.Itoa(int(month)) + "_" + strconv.Itoa(year)
	// fmt.Println(OutDIR)
	_, err := os.Stat(OutDIR)
	if err != nil {
		os.MkdirAll(OutDIR, 0700)
	}
	os.Chdir(OutDIR)
	log.Println("Saving SIM config to OUTPUT DIR: ", OutDIR)
	vlib.SaveStructure(i, i.fname, true)
	//SwitchBack()
	os.Chdir(CurrDIR)
}

func (i *SIMconfig) Read(f string) error {
	i.SetDefaults()
	i.fname = f
	viper.AddConfigPath(InDIR)
	// viper.SetConfigName(f)
	viper.SetConfigFile(InDIR + "/" + f)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("ReadInConfig Error: ", err)
		return err
	}
	err = viper.Unmarshal(i)
	if err != nil {
		log.Print("Error unmarshalling ", err)
		return err
	}
	return nil
}

// ReadSIMConfig reads all the configuration for the app
func ReadSIMConfig(configname string) (SIMconfig, error) {
	var cfg SIMconfig
	// fmt.Println(InDIR)
	err := cfg.Read(configname)
	return cfg, err
}
