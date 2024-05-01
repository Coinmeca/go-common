package chain

import (
	ABI "github.com/coinmeca/go-common/abi"
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

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

var ChainIdMap = map[int]string{
	80001:  "mumbai",
	137:    "polygon",
	1101:   "polyzkm",
	1442:   "polyzkt",
	42161:  "arbitrum",
	421613: "arbitgo",
}

var HTTPSProvider = map[int]string{
	80001:  HTTPSAlchemyMumbai,
	421613: HTTPSAlchemyArbitgo,
	1442:   HTTPSAlchemyPzkT,
}

var (
	CTX      context.Context
	CHAINxID int

	OrderbookABI abi.ABI
	MarketABI    abi.ABI
	VaultABI     abi.ABI
	FarmABI      abi.ABI

	BookDecimals  map[string]int32
	TokenDecimals map[string]int32
	TokenSymbols  map[string]string
	CAxBOOKS      []common.Address
	CAxVAULT      common.Address
)

var (
	TP1 = common.HexToHash(ABI.BuyEventHash)
	TP2 = common.HexToHash(ABI.SellEventHash)
	TP4 = common.HexToHash(ABI.DepositEventHash)
	TP5 = common.HexToHash(ABI.WithdrawEventHash)
	TP6 = common.HexToHash(ABI.ListingEventHash)
)

var WSSProvider = map[int]string{
	80001:  WSSAlchemyMumbai,
	421613: WSSAlchemyArbitgo,
	1442:   WSSQuickPzkT,
}

const ( // polygon - mumbai
	ETHxDAIMumbai    = "0x6a74a7ce1c79bc658f0dbffc1acf928302b2a922" //"0x6a74A7cE1C79bC658f0dbFFc1aCf928302b2A922"
	ETHxUSDTMumbai   = "0xdcad16d69c55c070a961fd530d216f9a89615f62" //"0xDcaD16d69C55c070A961FD530d216F9A89615f62"
	ETHxUSDCMumbai   = "0x943a2589e701e21f37ac0d0e1838e3854bd3f5e4" //"0x943a2589E701E21F37aC0d0E1838e3854BD3f5e4"
	MATICxDAIMumbai  = "0x49f6350246ae30d99ee0eeba456a84f33a01e57a" //"0x49f6350246aE30D99EE0EEba456A84f33a01E57a"
	MATICxUSDTMumbai = "0x05ee6e7b19dc904fb536a68051bf8ed95b84ad78" //"0x05ee6E7B19Dc904Fb536a68051Bf8eD95B84aD78"
	MATICxUSDCMumbai = "0x9845488eb4c232b8cf993b3f5e430523fdb8507a" //"0x9845488eB4c232b8Cf993b3f5e430523fDB8507A"

	HistoryMumbai = "0xb5b85e2023105e10b8410e9b6d37aef879310918" //"0xb5B85E2023105e10b8410E9B6d37Aef879310918"
	MarketMumbai  = "0xec9727df088b6fc6d6fc663338e054e3bcca6d92" //"0xeC9727DF088b6fC6D6fc663338e054E3BCca6d92"
	ReserveMumbai = "0x5c7fa590e7106192932add5a936fcee480884246" //"0x5c7fa590E7106192932Add5a936FCee480884246"
	FarmMumbai    = "0x6aa5fc53b6c01517fbf646f1828cc51c60d2aa68" //"0x6AA5fC53b6C01517FbF646f1828CC51C60d2aA68"
	VaultMumbai   = "0xd5559f8513356f8a199f166d7c2374351a18b5c3" //"0xd5559f8513356f8a199f166d7c2374351a18B5C3"
	CreditMumbai  = "0x381f3082ca499fbd8e6f704a65e56da3eaf70f15" //"0x381f3082ca499Fbd8E6F704A65E56da3EAF70f15"
	AppMumbai     = "0xec3c5f31308c0d0d03b1d99f4cc5b797551d8806" //"0xeC3c5f31308C0D0d03B1d99F4cC5b797551d8806"

	ETHAddressMumbai   = "0x744e85a6b15c5a9b3dae1692e506e2d4f52168fb" //"0x744E85A6B15c5A9B3daE1692e506e2D4f52168fB"
	USDTAddressMumbai  = "0x85e136cc3c5e3084c2d98d2857598c450e370843" //"0x85e136CC3C5E3084c2D98d2857598c450e370843"
	USDCAddressMumbai  = "0x4c7ef3a51a227405c31dc416d3e54c373ab3a059" //"0x4c7ef3A51A227405C31dC416d3e54C373ab3a059"
	DAIAddressMumbai   = "0x326324049bef9ef815ef614ef1b2203efcc5a3da" //"0x326324049Bef9ef815Ef614EF1b2203eFCC5A3dA"
	MATICAddressMumbai = "0x4940f2faf081760d3ebe59bc9f70a59b395a27b4" //"0x4940f2fAF081760D3EBE59Bc9f70A59b395A27b4"
	WBTCAddressMumbai  = "0x47d9f77f0b8f74ecc72611057c4556f63f6b6012" //"0x47d9F77f0B8f74ECC72611057c4556f63f6b6012"
	WSOLAddressMumbai  = "0x6886fc9f253f1eeb65c006a3dcba4a8331d05c74" //"0x6886fC9f253F1eEB65C006a3DcbA4a8331d05c74"
	MECAAddressMumbai  = "0xec85450dacdd3ce0a9f50091e20002b3100786b0" //"0xEc85450dacDd3Ce0a9F50091E20002B3100786B0"
)

const (
	WSSAlchemyMumbai = "wss://polygon-mumbai.g.alchemy.com/v2/Uv9c_wrKgskcO9oGg55GoaX04SwAE8ep"
	WSSInfuraMumbai  = "wss://polygon-mumbai.infura.io/v3/0d5f631f4514442b9cf965dfb974a115"
	WSSQuickMumbai   = "wss://quiet-black-diamond.matic-testnet.discover.quiknode.pro/8a00605626a2d6a895c985bb1b2e3a14752fb6bf/"
	WSPolygonEdge    = "ws://104.199.239.170:10002/ws"

	HTTPSAlchemyMumbai2 = "https://polygon-mumbai.g.alchemy.com/v2/Uv9c_wrKgskcO9oGg55GoaX04SwAE8ep"
	HTTPSAlchemyMumbai  = "https://polygon-mumbai.g.alchemy.com/v2/2u4p3yHD_5v0Ar448eEW4ihzxqZ1zqDl" // meca.net
	HTTPSInfuraMumbai   = "https://polygon-mumbai.infura.io/v3/0d5f631f4514442b9cf965dfb974a115"
	HTTPPolygonEdge     = "http://104.199.239.170:10002"
)

const (
	WSSAlchemyArbitgo    = "wss://arb-goerli.g.alchemy.com/v2/lFkQIRi56VhrvFGm45rGrNqgHdIrowKV"
	WSSAlchemyArbitrum   = "https://arb-mainnet.g.alchemy.com/v2/wja3-V5Edze3Bwd2YyBw0XwgndD-NCIH"
	HTTPSAlchemyArbitgo  = "https://arb-goerli.g.alchemy.com/v2/lFkQIRi56VhrvFGm45rGrNqgHdIrowKV"
	HTTPSAlchemyArbitrum = "https://arb-mainnet.g.alchemy.com/v2/wja3-V5Edze3Bwd2YyBw0XwgndD-NCIH"
)

const (
	MarketArbitgo = "0xfb03c105916eed91f9b585023c4cd61279d6f45b"
	FarmArbitgo   = "0x2e60aacffc050dcf53389834ea0fcd853823fa3e"
	VaultArbitgo  = "0x340b3b844342bd67991de496a12837b2cde064e4"
)

const ( // polygon zkEVM - testnet
	ETHxDAIPzkT    = "0xc3c75ce7499e70bcac8d422dc8b025a021931589" //"0xc3c75ce7499E70bcAC8d422dc8B025a021931589"
	ETHxUSDTPzkT   = "0xa097aeb83927a186c5196ccf9fd20a050e91f278" //"0xa097Aeb83927A186c5196cCF9Fd20A050e91f278"
	ETHxUSDCPzkT   = "0x1b4e37ee45c018fb9a08312414a83a622093b2bc" //"0x1B4E37eE45C018FB9a08312414A83a622093B2BC"
	MATICxDAIPzkT  = "0xb6bcb0b8ae40dd43beaba823d1a9879c5ec94479" //"0xB6bCb0B8Ae40dD43beaBA823D1A9879C5ec94479"
	MATICxUSDTPzkT = "0xec932d34ede99e7c6f85e9784cce3388539929bb" //"0xEC932d34edE99E7c6f85E9784ccE3388539929bb"
	MATICxUSDCPzkT = "0xb579959209aeb4d1338bf70d46a63b5134fdef06" //"0xb579959209aeB4d1338bf70D46A63b5134FDeF06"

	HistoryPzkT = "0xb30b342e55f50af2660b17fd5b2df6e1623c8533" //"0xb30B342E55F50af2660B17Fd5B2Df6e1623c8533"
	MarketPzkT  = "0x36cb67330345a06b5a5d22e237eebb9cf3c64d7c" //"0x36cB67330345A06B5A5D22E237EeBb9Cf3C64d7c"
	ReservePzkT = "0xe305e49c16ac4eb9a2d70c2471c0cb06b5e437bf" //"0xe305E49c16AC4Eb9a2D70c2471C0Cb06B5e437BF"
	FarmPzkT    = "0xb8a20a315210995179e771416688a3a5e274aaeb" //"0xB8A20a315210995179e771416688a3A5e274aAeB"
	VaultPzkT   = "0xbf81fef567ea570d46348d93143f04710de5b53d" //"0xbF81fEf567eA570d46348D93143f04710De5b53d"
	CreditPzkT  = "0x1ae771c30d9dd235b27d90e7382cabcf2c0f86f2" //"0x1ae771C30d9dd235b27d90e7382cABCf2c0F86F2"
	AppPzkT     = "0x4c09b30a4ffb3a67f6096599e187e39608c3df5f" //"0x4c09B30a4ffB3a67F6096599E187e39608c3DF5F"

	ETHAddressPzkT   = "0xb5e934624ae0554f07237a155b894e841232c8df" //"0xB5E934624aE0554f07237a155B894e841232C8dF"
	USDTAddressPzkT  = "0x298cdc3874e806beb3b35fcc65ec9760225e8e0a" //"0x298CdC3874e806bEB3b35FcC65ec9760225E8E0A"
	USDCAddressPzkT  = "0xfd55a8d294d9a3aa1bf229c872a9ecb598b32455" //"0xfD55A8d294D9A3AA1bF229C872A9ECB598b32455"
	DAIAddressPzkT   = "0xe52b3b996c284c76f9f23b2564e44b6523f98944" //"0xE52b3B996C284c76f9F23b2564e44B6523f98944"
	MATICAddressPzkT = "0x20b3b764559badd5c23f7f89e64d0bf03b89e801" //"0x20b3b764559BADd5c23f7f89e64D0Bf03b89e801"
	WBTCAddressPzkT  = "0xcec5e5151c0509dd0d7850730580159ad8e6e300" //"0xCeC5E5151C0509DD0d7850730580159AD8e6e300"
	WSOLAddressPzkT  = "0x9b02b69cfc6e7fd1ec0486808da5463026613ba9" //"0x9b02B69CFC6e7fd1eC0486808dA5463026613BA9"
	MECAAddressPzkT  = "0x6523717a1e3bdce24856c7fabf57a536b753f4c2" //"0x6523717a1E3BDCe24856c7faBF57a536B753f4c2"
)

const (
	WSSQuickPzkT     = "wss://flashy-lively-card.zkevm-testnet.discover.quiknode.pro/65d931ebf9680cff5f0017d1b632f2495157deb3/"
	HTTPSQuickPzkT   = "https://flashy-lively-card.zkevm-testnet.discover.quiknode.pro/65d931ebf9680cff5f0017d1b632f2495157deb3/"
	HTTPSAlchemyPzkT = "https://polygonzkevm-testnet.g.alchemy.com/v2/Z6zfDxEi6FLFvceu47KqCO7V9ahThspW"
)

type OrderbookID uint32

const ( // polygon - mainnet
	ETHxDAIMainnet = ""
	MarketMainnet  = ""
)

const (
	WSSAlchemyMainnet = "wss://polygon-mainnet.g.alchemy.com/v2/cFseXWeJqslDzJMj31zhh65_XQDNdP-_"
	WSSInfuraMainnet  = "wss://polygon-mainnet.infura.io/v3/0d5f631f4514442b9cf965dfb974a115"

	HTTPSAlchemyMainnet = "https://polygon-mainnet.g.alchemy.com/v2/cFseXWeJqslDzJMj31zhh65_XQDNdP-_"
	HTTPSInfuraMainnet  = "https://polygon-mainnet.infura.io/v3/0d5f631f4514442b9cf965dfb974a115"
)

const ( // general
	APICurrencyRate = "https://quotation-api-cdn.dunamu.com/v1/forex/recent?codes=FRX.KRWUSD"
	CMCLatestQuote  = "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest"
	DateFormat      = "2006-01-02 15:04:05"
)

type AddressPair map[string]string

var MecaAddrMap = map[int]AddressPair{
	80001: {
		"MARKET": MarketMumbai,
		"VAULT":  VaultMumbai,
		"FARM":   FarmMumbai,
	},
	1442: {
		"MARKET": MarketPzkT,
		"VAULT":  VaultPzkT,
		"FARM":   FarmPzkT,
	},
	421613: {
		"MARKET": MarketArbitgo,
		"VAULT":  VaultArbitgo,
		"FARM":   FarmArbitgo,
	},
}
