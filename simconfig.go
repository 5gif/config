package config

import (
	"fmt"
	"os"
	"path/filepath"

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
	fname          string
	ActiveBSCells  int
	ActiveUECells  int
	Extended       bool
	ForceAllLOS    bool
	ShadowLoss     bool
	LogInfo        bool
	UEcells        []int
	BScells        []int
	TrueCells      int
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

}

// Save ...
func (i *SIMconfig) Save() {
	//Switch Input
	pwd, _ := os.Getwd()
	currentdir := pwd
	rel, _ := filepath.Rel(currentdir, OutDIR)
	_ = rel
	os.Mkdir(OutDIR, 0700)
	os.Chdir(OutDIR)
	log.Println("Switching to OUTPUT DIR ", OutDIR)
	vlib.SaveStructure(i, i.fname, true)
	//SwitchBack()
	os.Chdir(currentdir)
}

func (i *SIMconfig) Read(f string) {
	i.SetDefaults()
	i.fname = f
	viper.AddConfigPath(InDIR)
	// viper.SetConfigName(f)
	viper.SetConfigFile(InDIR + "/" + f)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("ReadInConfig ", err)
	}
	err = viper.Unmarshal(i)
	if err != nil {
		log.Print("Error unmarshalling ", err)
	}

}

// ReadSIMConfig reads all the configuration for the app
func ReadSIMConfig(configname string, indir string) SIMconfig {
	var cfg SIMconfig
	fmt.Println(InDIR)
	cfg.Read(configname)
	return cfg
}
