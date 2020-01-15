package config

// import (
// 	"gonum.org/v1/gonum/mat"
// 	log "github.com/Sirupsen/logrus"

// 	"github.com/spf13/viper"
// 	"github.com/wiless/vlib"
// )

// /* func SetDir(in, out string) {
// 	CurrDIR, _ = os.Getwd()
// 	InDIR = CurrDIR + "/" + in
// 	OutDIR = CurrDIR + "/" + out
// 	os.MkdirAll(OutDIR, os.ModePerm)

// }*/

// func (ds *Spread) FindSpread(){
// 		ds.mu
// }
// ChannelParams.Spread.mu
// ChannelParams.DS.mu
// // TestEnvironment ....
// type ChannelParams struct {
// 	Env string
// 	CorrMatrix mat.Dense //
// 	CorrDistance map[string][]float64
// 	DS Spreads
// 	 Spreads

// 	Shadowfading struct{}
// 	Clusterparameters struct{}
// 	Clusterspreads struct{}
// 	Fname string
// 	}

// 	type Spread struct {
// 		mu    []float64
// 		sigma []float64
// 	}
// 	// ASD struct {
// 	// 	mu_ASD    []float64 `json:"mu_ASD"`
// 	// 	sigma_ASD []float64 `json:"sigma_ASD"`
// 	// }
// 	ASA struct {
// 		mu_ASA    []float64 `json:"mu_ASA"`
// 		sigma_ASA []float64 `json:"sigma_ASA"`
// 	}
// 	ZSA struct {
// 		mu_ZSA    []float64 `json:"mu_ZSA"`
// 		sigma_ZSA []float64 `json:"sigma_ASA"`
// 	}
// 	ShadowFading []int `json:"ShadowFading"`
// 	K_dB         struct {
// 		mu_K_dB    []int `json:"mu_K_dB"`
// 		sigma_K_dB []int `json:"TxPowerDbm"`
// 	}
// 	var DF Xcorr

// 	CorrelationDistance struct {
// 		corr_dist_DS  []int `json:"corr_dist_DS"`
// 		corr_dist_ASD []int `json:"corr_dist_ASD"`
// 		corr_dist_ASA []int `json:"corr_dist_ASA"`
// 		corr_dist_SF  []int `json:"corr_dist_SF"`
// 		corr_dist_K   []int `json:"corr_dist_K"`
// 		corr_dist_ZSA []int `json:"corr_dist_ZSA"`
// 		corr_dist_ZSD []int `json:"corr_dist_ZSD"`
// 	}
// 	Delay_scaling_Parameter []float64 `json:"Delay_scaling_Parameter"`
// 	XPR                     struct {
// 		mu_XPR    []int `json:"mu_XPR"`
// 		sigma_XPR []int `json:"sigma_XPR"`
// 	}
// 	number_of_clusters    []int `json:"number_of_clusters"`
// 	rays_within_a_cluster []int `json:"rays_within_a_cluster"`
// 	clusterspreads        struct {
// 		cDS  []int `json:"cDS"`
// 		cASD []int `json:"cASD"`
// 		cASA []int `json:"cASA"`
// 		cZSA []int `json:"cZSA"`
// 	}
// 	Shadowstd []int     `json:"Shadowstd"`
// 	mu_ZSD    []float64 `json:"mu_ZSD"`
// 	mu_offset []float64 `json:"mu_offset"`
// 	fname     string
// }

// Xcorr{
// 	SF   []float64
// 	DF   []float64
// 	K   []float64
// 	ASA   []float64
// 	ASD   []float64
// 	ZSA   []float64
// 	crosscorr_ZSD   []float64
// }

// //SetDefaults loads the default values for the simulation
// func (i *TestEnvironment) SetDefaults() {

// 	i.ENV = "RMa"
// 	i.mu_delay_spread = []float64{-7.49, 7.43, -7.47}
// 	i.sigma_delay_spread = []float64{0.55, 0.48, 0.24}
// 	i.mu_ASD = []float64{0.90, 0.95, 0.67}
// 	i.sigma_ASD = []float64{0.38, 0.45, 0.18}
// 	i.mu_ASA = []float64{1.52, 1.52, 1.66}
// 	i.sigma_ASA = []float64{0.24, 0.13, 0.21}
// 	i.mu_ZSA = []float64{0.47, 0.58, 0.93}
// 	i.sigma_ZSA = []float64{0.40, 0.37, 0.22}
// 	i.ShadowFading = []int{0, 0, 8}
// 	i.mu_K_dB = []int{7, 0, 0}
// 	i.sigma_K_dB = []int{4, 0, 0}
// 	i.crosscorr_SFvsK = []float64{0, 0, 0}
// 	i.crosscorr_DSvsSF = []float64{-0.5, -0.5, 0}
// 	i.crosscorr_ASDvsSF = []float64{0, 0.6, 0}
// 	i.crosscorr_ASAvsSF = []float64{0, 0, 0}
// 	i.crosscorr_ZSDvsSF = []float64{0.01, -0.04, 0}
// 	i.crosscorr_ZSAvsSF = []float64{-0.17, -0.25, 0}
// 	i.crosscorr_DSvsK = []float64{0, 0, 0}
// 	i.crosscorr_ASDvsK = []float64{0, 0, 0}
// 	i.crosscorr_ASAvsK = []float64{0, 0, 0}
// 	i.crosscorr_ZSDvsK = []float64{0, 0, 0}
// 	i.crosscorr_ZSAvsK = []float64{0, 0, 0}
// 	i.crosscorr_ASDvsDS = []float64{0, -0.4, 0}
// 	i.crosscorr_ASAvsDS = []float64{0, 0, 0}
// 	i.crosscorr_ZSDvsDS = []float64{-0.05, -0.1, 0}
// 	i.crosscorr_ZSAvsDS = []float64{0.27, -0.4, 0}
// 	i.crosscorr_ASDvsASA = []float64{0, 0, -0.5}
// 	i.crosscorr_ZSDvsASD = []float64{0.73, 0.42, 0.66}
// 	i.crosscorr_ZSAvsASD = []float64{-0.14, -0.27, 0.47}
// 	i.crosscorr_ZSDvsASA = []float64{-0.20, -0.18, -0.55}
// 	i.crosscorr_ZSAvsASA = []float64{0.24, 0.26, -0.22}
// 	i.crosscorr_ZSDvsZSA = []float64{-0.07, -0.27, 0}
// 	i.corr_dist_DS = []int{50, 36, 36}
// 	i.corr_dist_ASD = []int{25, 25, 30}
// 	i.corr_dist_ASA = []int{35, 35, 40}
// 	i.corr_dist_SF = []int{37, 120, 120}
// 	i.corr_dist_DS = []int{40, 0, 0}
// 	i.corr_dist_ZSA = []int{15, 50, 50}
// 	i.corr_dist_ZSD = []int{15, 50, 50}
// 	i.Delay_scaling_Parameter = []float64{3.8, 1.7, 1.7}
// 	i.mu_XPR = []int{12, 7, 7}
// 	i.sigma_XPR = []int{4, 3, 3}
// 	i.number_of_clusters = []int{11, 10, 10}
// 	i.rays_within_a_cluster = []int{20, 20, 20}
// 	i.cDS = []int{0, 0, 0}
// 	i.cASD = []int{2, 2, 2}
// 	i.cASA = []int{3, 3, 3}
// 	i.cZSA = []int{3, 3, 3}
// 	i.Shadowstd = []int{3, 3, 3}
// 	i.mu_ZSD = []float64{0, 0}
// 	i.mu_offset = []float64{0, 0}
// 	// Save config
// }
// func (i *TestEnvironment) Save() {

// 	SwitchOutput()
// 	// vlib.SaveStructure(s, "OutputSetting.json", true)

// 	vlib.SaveStructure(i, i.fname, true)
// 	SwitchBack()

// 	// os.Chdir(CurrDIR)
// }

// func (i *TestEnvironment) Read(f string) error {
// 	i.SetDefaults()
// 	i.fname = f
// 	viper.AddConfigPath(InDIR)
// 	// viper.SetConfigName(f)
// 	viper.SetConfigFile(InDIR + "/" + f)
// 	viper.SetConfigType("json")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Print("ReadInConfig Error: ", err)
// 	}
// 	err = viper.Unmarshal(i)
// 	if err != nil {
// 		log.Print("Error unmarshalling ", err)
// 	}
// 	return err
// }

// // ReadITUConfig reads all the configuration for the app
// func ReadEnvironmentConfig(configname string) (TestEnvironment, error) {
// 	var cfg TestEnvironment
// 	// fmt.Println(InDIR)
// 	err := cfg.Read(configname)
// 	return cfg, err
// }
