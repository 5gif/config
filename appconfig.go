package config

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

// App - contains nested structures of other config structures
type App struct {
	AppSetting
	WorkingDir string

	ITUcfg      ITUconfig
	NRcfg       NRconfig
	SIMcfg      SIMconfig
	loadsuccess bool
}



// LoadSetting
func (app App) Ready() bool {
	return app.loadsuccess
}

// SwitchBack ...
func (app App) SwitchBack() {

	os.Chdir(app.WorkingDir)
}

// SwitchInput ...
func (app App) SwitchInput() {
	// GOPATH := os.Getenv("GOPATH")
	// InDIR = GOPATH + "/src/github.com/5gif/" + indir
	log.Println("Switching to Input DIR : ", app.INdir)
	os.Chdir(app.INdir)
}

// SwitchOutput ...
func (app App) SwitchOutput() {
	// GOPATH := os.Getenv("GOPATH")
	log.Println("Switching to Output DIR : ", app.OUTdir)
	os.Chdir(app.OUTdir)
}

//SetCfgs Set the ITUcfg,NRcfg,SIMcfg from the filenames (nr,itu,sim) given App.Appsettings
func (app *App) SetCfgs(itu ITUconfig,
	nrc NRconfig,
	sim SIMconfig) error {
	app.ITUcfg = itu
	app.NRcfg = nrc
	app.SIMcfg = sim
	app.SIMcfg.Init(itu, nrc)
	app.loadsuccess = true
	return nil
}

//LoadCfgs loads the ITUcfg,NRcfg,SIMcfg from the filenames (nr,itu,sim) given App.Appsettings
func (app *App) LoadCfgs() error {

	// SwitchInput()
	app.loadsuccess = false
	var err1, err2, err3, err error
	var ITUcfg ITUconfig
	var NRcfg NRconfig
	var SIMcfg SIMconfig

	// var Channelcfg cirConfig.TestEnvironment
	if app.ITUfname == "" || app.NRfname == "" || app.SIMfname == "" {
		return errors.New("No Config Filenames given the ini file")
	}
	log.Println("I am here ", app.ITUfname)
	log.Info("Loading ITU Config ", app.ITUfname)

	if ITUcfg, err1 = ReadITUConfig(app.INdir + "/" + app.ITUfname); err1 != nil {
		log.Info("Loading ITU Config ..failed ", err1)
		return err1
	}

	log.Info("Loading NR Config")
	if NRcfg, err2 = ReadNRConfig(app.INdir + "/" + app.NRfname); err2 != nil {
		log.Errorln("Loading NR Config ..failed ", err2)
		return err2
	}

	log.Info("Loading SIM Config")
	if SIMcfg, err3 = ReadSIMConfig(app.INdir + "/" + app.SIMfname); err3 != nil {
		log.Info("Loading SIM Config ..failed ", err3)
		return err3
	}
	SIMcfg.SetSIMconfig(ITUcfg, NRcfg)
	/// ----- DONE loading all the ITU, 3GPP ,  SimConfig and related Antenna config

	if err1 != nil || err2 != nil || err3 != nil {
		log.Println(err1, err2, err3)

		err = errors.New("Setup Error: Files not read correctly. Using default values")
	} else {
		app.loadsuccess = true
		err = nil
	}

	if err == nil {
		app.ITUcfg = ITUcfg
		app.NRcfg = NRcfg
		app.SIMcfg = SIMcfg

		// app.WorkingDir, _ = os.Getwd()
		log.Info("Loaded App ITU,NR,SIM cfg from : ", app.INdir)
		log.Info("Working Dir : ", app.WorkingDir)
		log.Info("Creating Outdir : ", app.OUTdir)
		os.MkdirAll(app.OUTdir, os.ModePerm)

	}

	return err
}
