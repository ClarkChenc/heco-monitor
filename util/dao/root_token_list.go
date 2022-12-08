package dao

import "github.com/x-project/heco/util/types"

var (
	RootAccountMap map[string][]RootAccountToken
	RootTokenMap   map[string]RootToken
)

func initAccountList() {
	RootAccountMap = make(map[string][]RootAccountToken)

	RootAccountMap[types.HXTZ] = []RootAccountToken{
		{
			TokenName: types.HXTZ,
			Account:   "tz1gMsMPNBnq4AqnjrAgG536DP12uVLFgRbt",
		},
	}
	RootAccountMap[types.HLTC] = []RootAccountToken{
		{
			TokenName: types.HLTC,
			Account:   "MVFs5Ez9WypyjTaFre1FzPPt9EUBWiEyVF",
		},
	}

	RootAccountMap[types.HFIL] = []RootAccountToken{
		{
			TokenName: types.HFIL,
			Account:   "f1rmrlyjzym3jgv6tt77kbbbwbz6huxbimlt5va7y",
		},
	}

	RootAccountMap[types.HDOT] = []RootAccountToken{
		{
			TokenName: types.HDOT,
			Account:   "145pfvzsdw4zcZHv8v97rY1C8MfbKN5LKDV7JaNgm6tA65bE",
		},
	}
	RootAccountMap[types.HBTC] = []RootAccountToken{
		{
			TokenName: types.HBTC,
			Account:   "3P3imMaBZryYvxJMkm1vAk9UpXsjX5W2UM",
		},
	}
	RootAccountMap[types.HBSV] = []RootAccountToken{
		{
			TokenName: types.HBSV,
			Account:   "17ve2EPbtvUaQykXvTBhvKHX9e9uS2kFi5",
		},
		{
			TokenName: types.HBSV,
			Account:   "1BhrfpkREaMYN59LGYEwcdhXeJGYT9YUiz",
		},
	}
	RootAccountMap[types.HBCH] = []RootAccountToken{
		{
			TokenName: types.HBCH,
			Account:   "3P3imMaBZryYvxJMkm1vAk9UpXsjX5W2UM",
		},
	}

	RootAccountMap[types.ADA] = []RootAccountToken{
		{
			TokenName: types.ADA,
			Account:   "DdzFFzCqrht8unUUSLfCgAAmpQvWgr5ehLo5goTnrUUjUNN32ePhVH7tqzfVAJSaewQxxXBUV5s1VNFxLU7QZPPtS3CoGQTa4zBmGUMC",
		},
	}

	RootAccountMap[types.XRP] = []RootAccountToken{
		{
			TokenName: types.XRP,
			Account:   "r9bdH4ons6eKjhVFf5XsF68gaggyLjfzXw",
		},
	}

	RootAccountMap[types.DOGE] = []RootAccountToken{
		{
			TokenName: types.DOGE,
			Account:   "DFCT8ZPq6MytZU3HSvbTJevxt4MLq2Xasu",
		},
	}

	RootAccountMap[types.EOS] = []RootAccountToken{
		{
			TokenName: types.EOS,
			Account:   "vcefvctpvyog",
		},
	}

	RootAccountMap[types.NFT] = []RootAccountToken{
		{
			TokenName: types.NFT,
			Account:   "TB652Wm5L1iFrqE9D9Sddq343KUH2vLewo",
		},
	}

	RootAccountMap[types.XCH] = []RootAccountToken{
		{
			TokenName: types.XCH,
			Account:   "xch13rtaac9ktrt3q376ll5cj2us3lrgnjgxxu83h4p5sj3al36r0nkssl99hw",
		},
	}

	RootAccountMap[types.KAVA] = []RootAccountToken{
		{
			TokenName: types.KAVA,
			Account:   "kava1yzhx84adz9rzl9xtmmmdkwgfgfv0r6wjqyzvdd",
		},
	}

	RootAccountMap[types.ZEC] = []RootAccountToken{
		{
			TokenName: types.ZEC,
			Account:   "t1Ka3zQ1sYjrRuJxpNbVLKLc34Db4Nrj3p2",
		},
	}

	RootAccountMap[types.ETC] = []RootAccountToken{
		{
			TokenName: types.ETC,
			Account:   "0x265f2039e211cfacdc259c524788edaf2b6f22d8",
		},
	}

	RootAccountMap[types.TT] = []RootAccountToken{
		{
			TokenName: types.TT,
			Account:   "0x867475e552fb0cbd31f30a1634f55771a85897a7",
		},
	}

	RootAccountMap[types.THETA] = []RootAccountToken{
		{
			TokenName: types.THETA,
			Account:   "9c33fb730f6997536d2653c99acaa24e9f5f41a9",
		},
	}

	RootAccountMap[types.TRX] = []RootAccountToken{
		{
			TokenName: types.TRX,
			Account:   "TB652Wm5L1iFrqE9D9Sddq343KUH2vLewo",
		},
	}

	RootAccountMap[types.KSM] = []RootAccountToken{
		{
			TokenName: types.KSM,
			Account:   "CoXjP51UQ2Mmnux2zassu6umdxNZyE5P8Yna7BMditTX9JY",
		},
	}

	RootAccountMap[types.LUNA] = []RootAccountToken{
		{
			TokenName: types.LUNA,
			Account:   "terra10etqj2n84hhuan87wlxshtzm96mrmc2cxd5yxp",
		},
	}

	RootAccountMap[types.VET] = []RootAccountToken{
		{
			TokenName: types.VET,
			Account:   "0x9c33fb730f6997536d2653c99acaa24e9f5f41a9",
		},
	}

	RootAccountMap[types.SOL] = []RootAccountToken{
		{
			TokenName: types.SOL,
			Account:   "45PkYiW2jAwJhQpfHbrs6h4KDW2nzGTaLiy3PAigoYCx",
		},
	}

	RootAccountMap[types.ELA] = []RootAccountToken{
		{
			TokenName: types.ELA,
			Account:   "EXoZzV2qMvCG2je8Yt7Rs3ruDMF9FAMUsp",
		},
	}

	RootAccountMap[types.CTXC] = []RootAccountToken{
		{
			TokenName: types.CTXC,
			Account:   "324b1772e8ed7db84cd9fd75e99913f0a0a94139",
		},
	}

}

func initRootTokenMap() {
	RootTokenMap = make(map[string]RootToken)

	RootTokenMap[types.HBTC] = RootToken{
		TokenName:   types.HBTC,
		SideType:    types.SIDE_ETH,
		SideAddress: "0x0316eb71485b0ab14103307bf65a021042c6d380",
	}

	RootTokenMap[types.HDOT] = RootToken{
		TokenName:   types.HDOT,
		SideType:    types.SIDE_ETH,
		SideAddress: "0x9ffc3bcde7b68c46a6dc34f0718009925c1867cb",
	}

	RootTokenMap[types.HFIL] = RootToken{
		TokenName:   types.HFIL,
		SideType:    types.SIDE_ETH,
		SideAddress: "0x9afb950948c2370975fb91a441f36fdc02737cd4",
	}

	RootTokenMap[types.HBCH] = RootToken{
		TokenName:   types.HBCH,
		SideType:    types.SIDE_ETH,
		SideAddress: "0xaac679720204aaa68b6c5000aa87d789a3ca0aa5",
	}

	RootTokenMap[types.HLTC] = RootToken{
		TokenName:   types.HLTC,
		SideType:    types.SIDE_ETH,
		SideAddress: "0x2c000c0093de75a8fa2fccd3d97b314e20b431c3",
	}

	RootTokenMap[types.HXTZ] = RootToken{
		TokenName:   types.HXTZ,
		SideType:    types.SIDE_ETH,
		SideAddress: "0x4a10307e221781570e4b7e409eb315f11e8d0385",
	}

	RootTokenMap[types.HBSV] = RootToken{
		TokenName:   types.HBSV,
		SideType:    types.SIDE_ETH,
		SideAddress: "0x14007c545e6664c8370f27fa6b99dc830e6510a6",
	}

	RootTokenMap[types.ADA] = RootToken{
		TokenName:   types.ADA,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x843af718ef25708765a8e0942f89edeae1d88df0",
	}

	RootTokenMap[types.XRP] = RootToken{
		TokenName:   types.XRP,
		SideType:    types.SIDE_HECO,
		SideAddress: "0xa2f3c2446a3e20049708838a779ff8782ce6645a",
	}

	RootTokenMap[types.DOGE] = RootToken{
		TokenName:   types.DOGE,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x40280e26a572745b1152a54d1d44f365daa51618",
	}

	RootTokenMap[types.EOS] = RootToken{
		TokenName:   types.EOS,
		SideType:    types.SIDE_HECO,
		SideAddress: "0xae3a94a6dc7fce46b40d63bbf355a3ab2aa2a588",
	}

	RootTokenMap[types.NFT] = RootToken{
		TokenName:   types.NFT,
		SideType:    types.SIDE_HECO,
		SideAddress: "0xd2daf463cda501027cdf4a6d94749a72b8c7c72d",
	}

	RootTokenMap[types.XCH] = RootToken{
		TokenName:   types.XCH,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x66e8fda5abe2cd9910157010f7694de15a73a562",
	}

	RootTokenMap[types.KAVA] = RootToken{
		TokenName:   types.KAVA,
		SideType:    types.SIDE_HECO,
		SideAddress: "0xc37cf87d7525e3331d2cfe968837f42975a59de2",
	}

	RootTokenMap[types.ZEC] = RootToken{
		TokenName:   types.ZEC,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x631926e8b579fb5dfc2dbd7e849e06a4679f0b4c",
	}

	RootTokenMap[types.ETC] = RootToken{
		TokenName:   types.ETC,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x1d5e22d77df5bf3edd220b7b3637a7a917756468",
	}

	RootTokenMap[types.TT] = RootToken{
		TokenName:   types.TT,
		SideType:    types.SIDE_HECO,
		SideAddress: "0xa0befd31b118b169cc00b587a47f4c8563e80aa1",
	}

	RootTokenMap[types.THETA] = RootToken{
		TokenName:   types.THETA,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x3fb5a381a565c61325afe0660590a1784b207f57",
	}

	RootTokenMap[types.TRX] = RootToken{
		TokenName:   types.TRX,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x16f854238505fc04b4917a74c5466009b55c1c49",
	}

	RootTokenMap[types.KSM] = RootToken{
		TokenName:   types.KSM,
		SideType:    types.SIDE_HECO,
		SideAddress: "0xb6921d2eb7cef124d57c0216aedde2217c5db88a",
	}

	RootTokenMap[types.LUNA] = RootToken{
		TokenName:   types.LUNA,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x136abb69a1f078bf92106234a3a7cd2739cd55ae",
	}

	RootTokenMap[types.VET] = RootToken{
		TokenName:   types.VET,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x8650c13eeeb89610722b97696ed822ec33b93378",
	}

	RootTokenMap[types.SOL] = RootToken{
		TokenName:   types.SOL,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x2c73b1e8e02d832033ab1f7d26090331c88c1a7a",
	}

	RootTokenMap[types.ELA] = RootToken{
		TokenName:   types.ELA,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x102A56E6c2452bcee99dF8f61167E3e0f0749dbE",
	}

	RootTokenMap[types.CTXC] = RootToken{
		TokenName:   types.CTXC,
		SideType:    types.SIDE_HECO,
		SideAddress: "0x5a713f233e937b692eda1f7dd803ce12487b8bec",
	}
}

func init() {
	initAccountList()
	initRootTokenMap()
}
