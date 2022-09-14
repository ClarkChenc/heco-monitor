package dao

var (
	TokenMap map[string]BridgeToken
)

func init() {
	TokenMap = make(map[string]BridgeToken)

	TokenMap["ETH"] = BridgeToken{
		TokenName:   "ETH",
		EthAddress:  "0x0000000000000000000000000000000000000001",
		HecoAddress: "0x64FF637fB478863B7468bc97D30a5bF3A428a1fD",
	}
	TokenMap["UDST"] = BridgeToken{
		TokenName: "UDST",

		EthAddress:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
		HecoAddress: "0xa71edc38d189767582c38a3145b5873052c3e47a",
	}
	TokenMap["HT"] = BridgeToken{
		TokenName: "HT",

		EthAddress:  "0x6f259637dcd74c767781e37bc6133cd6a68aa161",
		HecoAddress: "0x5545153ccfca01fbd7dd11c0b23ba694d9509a6f",
	}
	TokenMap["HBTC"] = BridgeToken{
		TokenName: "HBTC",

		EthAddress:  "0x0316eb71485b0ab14103307bf65a021042c6d380",
		HecoAddress: "0x66a79D23E58475D2738179Ca52cd0b41d73f0BEa",
	}
	TokenMap["MX"] = BridgeToken{
		TokenName: "MX",

		EthAddress:  "0x11eef04c884e24d9b7b4760e7476d06ddf797f36",
		HecoAddress: "0x8d854e603dc777337134286f5b3408261736a88F",
	}
	TokenMap["SHIB"] = BridgeToken{
		TokenName: "SHIB",

		EthAddress:  "0x95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce",
		HecoAddress: "0xdd86dd2dc0aca2a8f41a680fc1f88ec1b7fc9b09",
	}
	TokenMap["USDC"] = BridgeToken{
		TokenName: "USDC",

		EthAddress:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		HecoAddress: "0x9362Bbef4B8313A8Aa9f0c9808B80577Aa26B73B",
	}
	TokenMap["UNI"] = BridgeToken{
		TokenName: "UNI",

		EthAddress:  "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
		HecoAddress: "0x22C54cE8321A4015740eE1109D9cBc25815C46E6",
	}

	TokenMap["HPT"] = BridgeToken{
		TokenName: "HPT",

		EthAddress:  "0xa66daa57432024023db65477ba87d4e7f5f95213",
		HecoAddress: "0xe499ef4616993730ced0f31fa2703b92b50bb536",
	}
	TokenMap["Seele"] = BridgeToken{
		TokenName: "Seele",

		EthAddress:  "0xb1e93236ab6073fdac58ada5564897177d4bcc43",
		HecoAddress: "0xA9634C25DcEA58dBA4Eb06caE307724b41cda241",
	}
	TokenMap["DAI"] = BridgeToken{
		TokenName: "DAI",

		EthAddress:  "0x6b175474e89094c44da98b954eedeac495271d0f",
		HecoAddress: "0x3D760a45D0887DFD89A2F5385a236B29Cb46ED2a",
	}
	TokenMap["LINK"] = BridgeToken{
		TokenName: "LINK",

		EthAddress:  "0x514910771af9ca656af840dff83e8264ecf986ca",
		HecoAddress: "0x9e004545c59d359f6b7bfb06a26390b087717b42",
	}
	TokenMap["WOO"] = BridgeToken{
		TokenName: "WOO",

		EthAddress:  "0x4691937a7508860f876c9c0a2a617e7d9e945d4b",
		HecoAddress: "0x3befb2308bce92da97264077faf37dcd6c8a75e6",
	}

	TokenMap["NEST"] = BridgeToken{
		TokenName: "NEST",

		EthAddress:  "0x04abeda201850ac0124161f037efd70c74ddc74c",
		HecoAddress: "0x2e3443e2ded0c769dd7229dac9ba37155f94fdf1",
	}
	TokenMap["AAVE"] = BridgeToken{
		TokenName: "AAVE",

		EthAddress:  "0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9",
		HecoAddress: "0x202b4936fE1a82A4965220860aE46d7d3939Bb25",
	}
	TokenMap["BIX"] = BridgeToken{
		TokenName: "BIX",

		EthAddress:  "0x009c43b42aefac590c719e971020575974122803",
		HecoAddress: "0x1229e81d6f79b0b49ec686749573307f919f2c71",
	}

	TokenMap["TUSD"] = BridgeToken{
		TokenName: "TUSD",

		EthAddress:  "0x0000000000085d4780b73119b644ae5ecd22b376",
		HecoAddress: "0x5ee41ab6edd38cdfb9f6b4e6cf7f75c87e170d98",
	}
	TokenMap["LAMB"] = BridgeToken{
		TokenName: "LAMB",

		EthAddress:  "0x8971f9fd7196e5cEE2C1032B50F656855af7Dd26",
		HecoAddress: "0xE131F048D85f0391A24435eEFB7763199B587d0e",
	}

	TokenMap["GOF"] = BridgeToken{
		TokenName: "GOF",

		EthAddress:  "0x488e0369f9bc5c40c002ea7c1fe4fd01a198801c",
		HecoAddress: "0x2AAFe3c9118DB36A20dd4A942b6ff3e78981dce1",
	}

	TokenMap["MATIC"] = BridgeToken{
		TokenName: "MATIC",

		EthAddress:  "0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0",
		HecoAddress: "0xdB11743fe8B129b49b11236E8a715004BDabe7e5",
	}

	TokenMap["SWFTC"] = BridgeToken{
		TokenName: "SWFTC",

		EthAddress:  "0x0bb217E40F8a5Cb79Adf04E1aAb60E5abd0dfC1e",
		HecoAddress: "0x329dda64Cbc4DFD5FA5072b447B3941CE054ebb3",
	}

	TokenMap["SUSHI"] = BridgeToken{
		TokenName: "SUSHI",

		EthAddress:  "0x6b3595068778dd592e39a122f4f5a5cf09c90fe2",
		HecoAddress: "0x52e00b2da5bd7940ffe26b609a42f957f31118d5",
	}

	TokenMap["MANA"] = BridgeToken{
		TokenName: "MANA",

		EthAddress:  "0x0f5d2fb29fb7d3cfee444a200298f468908cc942",
		HecoAddress: "0x09006b66d89e5213Fc173403AACBA30620A91F4e",
	}

	TokenMap["ACH"] = BridgeToken{
		TokenName: "ACH",

		EthAddress:  "0xed04915c23f00a313a544955524eb7dbd823143d",
		HecoAddress: "0x4a31D1Ad7430586752A1888fE947E3E7D52aFfB8",
	}

	TokenMap["CHZ"] = BridgeToken{
		TokenName: "CHZ",

		EthAddress:  "0x3506424f91fd33084466f402d5d97f05f8e3b4af",
		HecoAddress: "0xB37AD4461bB12ba3ac110ed20c3FefE173fa66D3",
	}

	TokenMap["HDOT"] = BridgeToken{
		TokenName: "HDOT",

		EthAddress:  "0x9ffc3bcde7b68c46a6dc34f0718009925c1867cb",
		HecoAddress: "0xA2c49cEe16a5E5bDEFDe931107dc1fae9f7773E3",
	}

	TokenMap["XTM"] = BridgeToken{
		TokenName: "XTM",

		EthAddress:  "0xcd1faff6e578fa5cac469d2418c95671ba1a62fe",
		HecoAddress: "0xA7BDA574f59CDA932e4Ba8B4c4225fEc134fF05E",
	}

	TokenMap["IMX"] = BridgeToken{
		TokenName: "IMX",

		EthAddress:  "0xf57e7e7c23978c3caec3c3548e3d615c346e79ff",
		HecoAddress: "0x813824DB5a0AE172E4691b7a35f8f4C3e79c051C",
	}

	TokenMap["YGG"] = BridgeToken{
		TokenName: "YGG",

		EthAddress:  "0x25f8087EAD173b73D6e8B84329989A8eEA16CF73",
		HecoAddress: "0x39359CdA2349BfB32d5A936c7E3BCaf4badbc437",
	}

	TokenMap["EPIK"] = BridgeToken{
		TokenName: "EPIK",

		EthAddress:  "0x4da0c48376c277cdbd7fc6fdc6936dee3e4adf75",
		HecoAddress: "0x96cb35b73A75C60aB4988069DC163e06b0bEaA9C",
	}

	TokenMap["SRM"] = BridgeToken{
		TokenName: "SRM",

		EthAddress:  "0x476c5e26a75bd202a9683ffd34359c0cc15be0ff",
		HecoAddress: "0xcE85f9A1d7056057fCbbf372f6717348366Dad3d",
	}

	TokenMap["SKU"] = BridgeToken{
		TokenName: "SKU",

		EthAddress:  "0xab7d26316f5c0adD393C2dDA2E6dA43D750DD399",
		HecoAddress: "0x420D50715532A68299311A998F6eFf1Dd0039206",
	}

	TokenMap["DFA"] = BridgeToken{
		TokenName: "DFA",

		EthAddress:  "0x62959c699a52ec647622c91e79ce73344e4099f5",
		HecoAddress: "0x8C746bc62b0C59fb65A89D7372Ac2C577989A252",
	}

	TokenMap["RING"] = BridgeToken{
		TokenName: "RING",

		EthAddress:  "0x9469d013805bffb7d3debe5e7839237e535ec483",
		HecoAddress: "0x15e65456310ecb216B51EfBd8a1dBf753353DcF9",
	}

	TokenMap["1INCH"] = BridgeToken{
		TokenName: "1INCH",

		EthAddress:  "0x111111111117dc0aa78b770fa6a738034120c302",
		HecoAddress: "0xD192f8e3224Ff0f48B08DB4791576B6878B426A0",
	}

	TokenMap["AXS"] = BridgeToken{
		TokenName: "AXS",

		EthAddress:  "0xbb0e17ef65f82ab018d8edd776e8dd940327b28b",
		HecoAddress: "0x0bde532913915cb6e2f0467b8e46e286cfcfb2bb",
	}

	TokenMap["CRU"] = BridgeToken{
		TokenName: "CRU",

		EthAddress:  "0x32a7c02e79c4ea1008dd6564b35f131428673c41",
		HecoAddress: "0x926875854917c871ebbd44b75ac33ac8386d4ed5",
	}

	TokenMap["CVP"] = BridgeToken{
		TokenName: "CVP",

		EthAddress:  "0x38e4adb44ef08f22f5b5b76a8f0c2d0dcbe7dca1",
		HecoAddress: "0xcb2210ba327d4919e740b427465373b408652d89",
	}

	TokenMap["DHT"] = BridgeToken{
		TokenName: "DHT",

		EthAddress:  "0xca1207647ff814039530d7d35df0e1dd2e91fa84",
		HecoAddress: "0xfc06c94bdb23ea9684053b47aedb70613184821f",
	}

	TokenMap["GRT"] = BridgeToken{
		TokenName: "GRT",

		EthAddress:  "0xc944e90c64b2c07662a292be6244bdf05cda44a7",
		HecoAddress: "0xfAdD0c7762c59cEBC5248019dBAC652319CEBDbd",
	}

	TokenMap["INSUR"] = BridgeToken{
		TokenName: "INSUR",

		EthAddress:  "0x544c42fBB96B39B21DF61cf322b5EDC285EE7429",
		HecoAddress: "0x6e2beb4a1375562acce354b00b192196515d108f",
	}

	TokenMap["LRC"] = BridgeToken{
		TokenName: "LRC",

		EthAddress:  "0xBBbbCA6A901c926F240b89EacB641d8Aec7AEafD",
		HecoAddress: "0xbf22F76657601A522Cf9Ac832718A3404302D6bC",
	}

	TokenMap["MASK"] = BridgeToken{
		TokenName: "MASK",

		EthAddress:  "0x69af81e73a73b40adf4f3d4223cd9b1ece623074",
		HecoAddress: "0xd651503b0Ca005920780810733f8a7F3b61c530B",
	}

	TokenMap["NHBTC"] = BridgeToken{
		TokenName: "NHBTC",

		EthAddress:  "0x1f832091faf289ed4f50fe7418cfbd2611225d46",
		HecoAddress: "0xD81D5ccC891329a3b97A6752dD0F3c6A5FC0F6Ea",
	}

	TokenMap["PHA"] = BridgeToken{
		TokenName: "PHA",

		EthAddress:  "0x6c5ba91642f10282b576d91922ae6448c9d52f4e",
		HecoAddress: "0x51aa144E0C2fEC20C5f19e25352838ff7fE1ed74",
	}

	TokenMap["UMA"] = BridgeToken{
		TokenName: "UMA",

		EthAddress:  "0x04fa0d235c4abf4bcf4787af4cf447de572ef828",
		HecoAddress: "0xe2563F0a1787ac8a9fA67f3124a3d0AE92d574d2",
	}

	TokenMap["WBTCHECO"] = BridgeToken{
		TokenName: "WBTCHECO",

		EthAddress:  "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
		HecoAddress: "0x70D171d269D964d14aF9617858540061e7bE9EF1",
	}

	TokenMap["ZKS"] = BridgeToken{
		TokenName: "ZKS",

		EthAddress:  "0xe4815ae53b124e7263f08dcdbbb757d41ed658c6",
		HecoAddress: "0xe25A7D0B6bcDa2a832d255edef1387E72d20883b",
	}

	TokenMap["MTA"] = BridgeToken{
		TokenName: "MTA",

		EthAddress:  "0xa3bed4e1c75d00fa6f4e5e6922db7261b5e9acd2",
		HecoAddress: "0x551d2be01bab5dd9b2f228ea32450c9437817366",
	}

	TokenMap["HFIL"] = BridgeToken{
		TokenName: "HFIL",

		EthAddress:  "0x9afb950948c2370975fb91a441f36fdc02737cd4",
		HecoAddress: "0xae3a768f9aB104c69A7CD6041fE16fFa235d1810",
	}

	TokenMap["SWRV"] = BridgeToken{
		TokenName: "SWRV",

		EthAddress:  "0xb8baa0e4287890a5f79863ab62b7f175cecbd433",
		HecoAddress: "0x0b7745aef68520ab5bb192d219e8007d276e7b95",
	}

	TokenMap["MLN"] = BridgeToken{
		TokenName: "MLN",

		EthAddress:  "0xec67005c4e498ec7f55e092bd1d35cbc47c91892",
		HecoAddress: "0x3194cfcceee1e364f9e80ac52e779651279a67ea",
	}

	TokenMap["DF"] = BridgeToken{
		TokenName: "DF",

		EthAddress:  "0x431ad2ff6a9c365805ebad47ee021148d6f7dbe0",
		HecoAddress: "0xd633ac35d763a0481c9d4bfb20c9a10785d1a8c1",
	}

	TokenMap["BNT"] = BridgeToken{
		TokenName: "BNT",

		EthAddress:  "0x1f573d6fb3f13d689ff844b4ce37794d79a7ff1c",
		HecoAddress: "0x756a3dCF7171Ad147A801318a7caB3284cFf6a85",
	}

	TokenMap["STPT"] = BridgeToken{
		TokenName: "STPT",

		EthAddress:  "0xde7d85157d9714eadf595045cc12ca4a5f3e2adb",
		HecoAddress: "0xd137d3009499c09387eb00eb3712b0984f28a60d",
	}

	TokenMap["SNX"] = BridgeToken{
		TokenName: "SNX",

		EthAddress:  "0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F",
		HecoAddress: "0x777850281719d5a96C29812ab72f822E0e09F3Da",
	}

	TokenMap["COMP"] = BridgeToken{
		TokenName: "COMP",

		EthAddress:  "0xc00e94cb662c3520282e6f5717214004a7f26888",
		HecoAddress: "0xCe0A5CA134fb59402B723412994B30E02f083842",
	}

	TokenMap["CRV"] = BridgeToken{
		TokenName: "CRV",

		EthAddress:  "0xd533a949740bb3306d119cc777fa900ba034cd52",
		HecoAddress: "0x6BCE534a02f8347f747124082Ef3e35Dd696748D",
	}

	TokenMap["DAIHECO"] = BridgeToken{
		TokenName: "DAIHECO",

		EthAddress:  "0x6b175474e89094c44da98b954eedeac495271d0f",
		HecoAddress: "0x3D760a45D0887DFD89A2F5385a236B29Cb46ED2a",
	}

	TokenMap["MKR"] = BridgeToken{
		TokenName: "MKR",

		EthAddress:  "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2",
		HecoAddress: "0x34D75515090902a513F009f4505A750efaaD63b0",
	}

	TokenMap["BAL"] = BridgeToken{
		TokenName: "BAL",

		EthAddress:  "0xba100000625a3754423978a60c9317c58a424e3D",
		HecoAddress: "0x045de15ca76e76426e8fc7cba8392a3138078d0f",
	}

	TokenMap["CHR"] = BridgeToken{
		TokenName: "CHR",

		EthAddress:  "0x8a2279d4a90b6fe1c4b30fa660cc9f926797baa2",
		HecoAddress: "0xf41A72C63D853217574a7f792a51b70F79ec8935",
	}

	TokenMap["HDOT"] = BridgeToken{
		TokenName: "HDOT",

		EthAddress:  "0x9ffc3bcde7b68c46a6dc34f0718009925c1867cb",
		HecoAddress: "0xA2c49cEe16a5E5bDEFDe931107dc1fae9f7773E3",
	}

	TokenMap["NEST"] = BridgeToken{
		TokenName: "NEST",

		EthAddress:  "0x04abeda201850ac0124161f037efd70c74ddc74c",
		HecoAddress: "0x2e3443e2ded0c769dd7229dac9ba37155f94fdf1",
	}

	TokenMap["SKM"] = BridgeToken{
		TokenName: "SKM",

		EthAddress:  "0x048fe49be32adfc9ed68c37d32b5ec9df17b3603",
		HecoAddress: "0x96674f8da3F9c6ACb4A56b393AF9A490D70D16d0",
	}

	TokenMap["FOR"] = BridgeToken{
		TokenName: "FOR",

		EthAddress:  "0x1fcdce58959f536621d76f5b7ffb955baa5a672f",
		HecoAddress: "0xcb3ca50f2b627ae3f907575a66834ec4434e341a",
	}

	TokenMap["ARPA"] = BridgeToken{
		TokenName: "ARPA",

		EthAddress:  "0xba50933c268f567bdc86e1ac131be072c6b0b71a",
		HecoAddress: "0x5A6B72Dd6209A770aE1C02a7A2E1900636072d0b",
	}

	TokenMap["FTT"] = BridgeToken{
		TokenName: "FTT",

		EthAddress:  "0x50d1c9771902476076ecfc8b2a83ad6b9355a4c9",
		HecoAddress: "0xC7f7a54892B78b5c812c58d9Df8035FcE9F4D445",
	}

	TokenMap["CNNS"] = BridgeToken{
		TokenName: "CNNS",

		EthAddress:  "0x6c3BE406174349cfa4501654313d97e6a31072e1",
		HecoAddress: "0x4BF06f76C68D81BDE1F87535fdCb60Adadb01CF5",
	}

	TokenMap["ARKO"] = BridgeToken{
		TokenName: "ARKO",

		EthAddress:  "0x8ab7404063ec4dbcfd4598215992dc3f8ec853d7",
		HecoAddress: "0x7047fa80c93214e5bf47f1d51a685511ddbe4041",
	}

	TokenMap["CRO"] = BridgeToken{
		TokenName: "CRO",

		EthAddress:  "0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b",
		HecoAddress: "0xF7e1E39e239C5A920849f435F66097D2e412859e",
	}

	TokenMap["RSR"] = BridgeToken{
		TokenName: "RSR",

		EthAddress:  "0x320623b8e4ff03373931769a31fc52a4e78b5d70",
		HecoAddress: "0x8f6303C8a21398c91abB5EfF3b3C9315CA59CbE8",
	}

	TokenMap["NEXO"] = BridgeToken{
		TokenName:   "NEXO",
		EthAddress:  "0xb62132e35a6c13ee1ee0f84dc5d40bad8d815206",
		HecoAddress: "0x3a6B8B642aa154bAdA89eD3730C9E820Da79AD24",
	}

	TokenMap["HXTZ"] = BridgeToken{
		TokenName:   "HXTZ",
		EthAddress:  "0x4a10307e221781570e4b7e409eb315f11e8d0385",
		HecoAddress: "0x45e97daD828AD735af1dF0473fc2735F0Fd5330c",
	}

	TokenMap["HBSV"] = BridgeToken{
		TokenName:   "HBSV",
		EthAddress:  "0x14007c545e6664c8370f27fa6b99dc830e6510a6",
		HecoAddress: "0xc2CB6B5357CcCE1B99Cd22232942D9A225Ea4eb1",
	}

	TokenMap["KCASH"] = BridgeToken{
		TokenName:   "KCASH",
		EthAddress:  "0x32d74896f05204d1b6ae7b0a3cebd7fc0cd8f9c7",
		HecoAddress: "0x8b4e6c32d7e401f2951e56155236ee4fe717f445",
	}

	TokenMap["MUSK"] = BridgeToken{
		TokenName:   "MUSK",
		EthAddress:  "0x719e7f0dadfdea25b78595da944f44d15d7e6795",
		HecoAddress: "0x245A3bb0fB1385522d883D8d8b4a91da17548C07",
	}

	TokenMap["IDT"] = BridgeToken{
		TokenName:   "IDT",
		EthAddress:  "0x02c4c78c462e32cca4a90bc499bf411fb7bc6afb",
		HecoAddress: "0x950bfdda329d120f6763d3c7bedb35d2880bf7bf",
	}

	TokenMap["YCC"] = BridgeToken{
		TokenName:   "YCC",
		EthAddress:  "0x37e1160184f7dd29f00b78c050bf13224780b0b0",
		HecoAddress: "0x80898823FdD0B961C2E9809eF75B673E7f4aa02D",
	}

	TokenMap["REN"] = BridgeToken{
		TokenName:   "REN",
		EthAddress:  "0x408e41876cccdc0f92210600ef50372656052a38",
		HecoAddress: "0x212208bcc81F3a2D0188afF76A2d39351eb53b96",
	}

	TokenMap["GET"] = BridgeToken{
		TokenName:   "GET",
		EthAddress:  "0x60c68a87be1e8a84144b543aacfa77199cd3d024",
		HecoAddress: "0xb4a1148d99876ce3d48f827136fe3f50967f2686",
	}

	TokenMap["MAN"] = BridgeToken{
		TokenName:   "MAN",
		EthAddress:  "0xe25bcec5d3801ce3a794079bf94adf1b8ccd802d",
		HecoAddress: "0x7a45d615f26da940ffff43eca1274d8ea116ce0c",
	}

	TokenMap["HOT"] = BridgeToken{
		TokenName:   "HOT",
		EthAddress:  "0x9af839687f6c94542ac5ece2e317daae355493a1",
		HecoAddress: "0x26dB8742DA87d2E74911BFA4A349D4f6F7fc6037",
	}

	TokenMap["GTC"] = BridgeToken{
		TokenName:   "GTC",
		EthAddress:  "0xb70835d7822ebb9426b56543e391846c107bd32c",
		HecoAddress: "0xb73a553dae0e7fa63cf7158a3aadd77418f5e458",
	}

	TokenMap["BUT"] = BridgeToken{
		TokenName:   "BUT",
		EthAddress:  "0xb2e260f12406c401874ecc960893c0f74cd6afcd",
		HecoAddress: "0x311bcb634a4111e6516d3899f9fbfbfe984f021a",
	}

	TokenMap["BOX"] = BridgeToken{
		TokenName:   "BOX",
		EthAddress:  "0x63f584fa56e60e4d0fe8802b27c7e6e3b33e007f",
		HecoAddress: "0x0aff28e1627daff9ae4dcea022a169325982b723",
	}

	TokenMap["LET"] = BridgeToken{
		TokenName:   "LET",
		EthAddress:  "0xfa3118b34522580c35ae27f6cf52da1dbb756288",
		HecoAddress: "0xce5f99937bf17751100c9d2326f1a2f27c395f8e",
	}

	TokenMap["YEE"] = BridgeToken{
		TokenName:   "YEE",
		EthAddress:  "0x922105fad8153f516bcfb829f56dc097a0e1d705",
		HecoAddress: "0xde4C8A05f1b665b48B5f5D232F9cb4eADaB54d21",
	}

	TokenMap["AST"] = BridgeToken{
		TokenName:   "AST",
		EthAddress:  "0x27054b13b1b798b345b591a4d22e6562d47ea75a",
		HecoAddress: "0xee3e56beef0606424e0ff34f93e967ba32c8da72",
	}

	TokenMap["STORJ"] = BridgeToken{
		TokenName:   "STORJ",
		EthAddress:  "0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac",
		HecoAddress: "0x98c5751010fd738d88ef177f77A4988E47133fB1",
	}

	TokenMap["ZRX"] = BridgeToken{
		TokenName:   "ZRX",
		EthAddress:  "0xe41d2489571d322189246dafa5ebde1f4699f498",
		HecoAddress: "0x0212dA773704cbc4F476bA827406363c87E8D3Bd",
	}

	TokenMap["BAT"] = BridgeToken{
		TokenName:   "BAT",
		EthAddress:  "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
		HecoAddress: "0xb04Ee982E6329FeBe4c70A53d1725469A1F6963A",
	}

	TokenMap["LTC"] = BridgeToken{
		TokenName:   "LTC",
		EthAddress:  "0x2c000c0093de75a8fa2fccd3d97b314e20b431c3",
		HecoAddress: "0xecb56cf772B5c9A6907FB7d32387Da2fCbfB63b4",
	}

	TokenMap["HT"] = BridgeToken{
		TokenName:   "HT",
		EthAddress:  "0x6f259637dcd74c767781e37bc6133cd6a68aa161",
		HecoAddress: "0x5545153ccfca01fbd7dd11c0b23ba694d9509a6f",
	}

	TokenMap["PAX"] = BridgeToken{
		TokenName:   "PAX",
		EthAddress:  "0x8e870d67f660d95d5be530380d0ec0bd388289e1",
		HecoAddress: "0x62057691DD4970dD109231b2658Ff201E15E767a",
	}

	TokenMap["HUSD"] = BridgeToken{
		TokenName:   "HUSD",
		EthAddress:  "0xdf574c24545e5ffecb9a659c229253d4111d87e1",
		HecoAddress: "0x0298c2b32eae4da002a15f36fdf7615bea3da047",
	}

	TokenMap["BCC"] = BridgeToken{
		TokenName:   "BCC",
		EthAddress:  "0xaac679720204aaa68b6c5000aa87d789a3ca0aa5",
		HecoAddress: "0xeF3CEBD77E0C52cb6f60875d9306397B5Caca375",
	}

	TokenMap["YFI"] = BridgeToken{
		TokenName:   "YFI",
		EthAddress:  "0x0bc529c00c6401aef6d220be8c6ea1667f6ad93e",
		HecoAddress: "0xB4F019bEAc758AbBEe2F906033AAa2f0F6Dacb35",
	}

	TokenMap["YFII"] = BridgeToken{
		TokenName:   "YFII",
		EthAddress:  "0xa1d0e215a23d7030842fc67ce582a6afa3ccab83",
		HecoAddress: "0x1e584d0ee01291ceac59a0ff75f3f4e1011bb18c",
	}

	TokenMap["WNXM"] = BridgeToken{
		TokenName:   "WNXM",
		EthAddress:  "0x0d438f3b5175bebc262bf23753c1e53d03432bde",
		HecoAddress: "0x817c9fa9d6c463b256f3aca9a0e37772421cb34c",
	}
}
