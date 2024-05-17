package commondefine

const (
	AribitrumMainnet = "AribitrumMainnet"
	ArbitrumSepolia  = "ArbitrumSepolia"
	AribitrumNova    = "AribitrumNova"
)

var ChainIdMap = map[string]string{
	"AribitrumMainnet": "42161",
	"ArbitrumSepolia":  "421614",
	"AribitrumNova":    "42170",
}

var ChainNameMap = map[string]int{
	"mumbai":    80001,
	"polygon":   137,
	"ethereum":  1,
	"bnb":       56,
	"arbitrum":  42161,
	"arbitgo":   421613,
	"optimism":  10,
	"polyzkm":   1101,
	"polyzkt":   1442,
	"avalanche": 43114,
}
