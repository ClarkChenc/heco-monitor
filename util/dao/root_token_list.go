package dao

import "github.com/x-project/heco/util/types"

var (
	RootAccountMap map[string][]RootAccountToken
	RootTokenMap   map[string]RootToken
)

func initAccountList() {
	RootAccountMap = make(map[string][]RootAccountToken)
	RootAccountMap[types.HBTC] = []RootAccountToken{
		{
			TokenName: types.HBTC,
			Account:   "15BAjnnitb4UPrPG8iGpe3p21hwdP6wwU7",
		},
		{
			TokenName: types.HBTC,
			Account:   "31rggRJ3iB8GE8yxTNa28pXhd24GbCQK3T",
		},
		{
			TokenName: types.HBTC,
			Account:   "36KAwNUR8VeLpUfGwdk7LEN6F4yvoRWMjn",
		},
	}

	RootAccountMap[types.HDOT] = []RootAccountToken{
		{
			TokenName: types.HDOT,
			Account:   "145pfvzsdw4zcZHv8v97rY1C8MfbKN5LKDV7JaNgm6tA65bE",
		},
		{
			TokenName: types.HDOT,
			Account:   "1ZDBfr37ZoU7iPeDNd5dE66xrrhUHURB3hDmE2HeUVBFyrc",
		},
	}

	RootAccountMap[types.HFIL] = []RootAccountToken{
		{
			TokenName: types.HFIL,
			Account:   "f1rmrlyjzym3jgv6tt77kbbbwbz6huxbimlt5va7y",
		},
		{
			TokenName: types.HFIL,
			Account:   "f1xsq7i5dm53l7xq5jqrw7exwciz6vdqro2w5kaey",
		},
	}

	RootAccountMap[types.HBCH] = []RootAccountToken{
		{
			TokenName: types.HBCH,
			Account:   "15BAjnnitb4UPrPG8iGpe3p21hwdP6wwU7",
		},
		{
			TokenName: types.HBCH,
			Account:   "3CDNA2DXsWPLFEUy1m2egXCxRSHwFugWGj",
		},
		{
			TokenName: types.HBCH,
			Account:   "3PfwaZNVCdgVwy5i667aB4bjypthLhazxB",
		},
	}

	RootAccountMap[types.HLTC] = []RootAccountToken{
		{
			TokenName: types.HLTC,
			Account:   "LiYNtSPX6azFE42PZShCfarUaveQg1Nw4M",
		},
		{
			TokenName: types.HLTC,
			Account:   "MStsioAFxVSBAEMEcackrC9eR3NmYDvwJV",
		},
	}

	RootAccountMap[types.HXTZ] = []RootAccountToken{
		{
			TokenName: types.HXTZ,
			Account:   "tz1gMsMPNBnq4AqnjrAgG536DP12uVLFgRbt",
		},
		{
			TokenName: types.HXTZ,
			Account:   "tz1V2dubScs1zrs7rE77yKsj44oFygxUMN5F",
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
}

func initRootTokenMap() {
	RootTokenMap = make(map[string]RootToken)

	RootTokenMap[types.HBTC] = RootToken{
		TokenName:  types.HBTC,
		EthAddress: "0x0316eb71485b0ab14103307bf65a021042c6d380",
	}

	RootTokenMap[types.HDOT] = RootToken{
		TokenName:  types.HDOT,
		EthAddress: "0x9ffc3bcde7b68c46a6dc34f0718009925c1867cb",
	}

	RootTokenMap[types.HFIL] = RootToken{
		TokenName:  types.HFIL,
		EthAddress: "0x9afb950948c2370975fb91a441f36fdc02737cd4",
	}

	RootTokenMap[types.HBCH] = RootToken{
		TokenName:  types.HBCH,
		EthAddress: "0xaac679720204aaa68b6c5000aa87d789a3ca0aa5",
	}

	RootTokenMap[types.HLTC] = RootToken{
		TokenName:  types.HLTC,
		EthAddress: "0x2c000c0093de75a8fa2fccd3d97b314e20b431c3",
	}

	RootTokenMap[types.HXTZ] = RootToken{
		TokenName:  types.HXTZ,
		EthAddress: "0x4a10307e221781570e4b7e409eb315f11e8d0385",
	}

	RootTokenMap[types.HBSV] = RootToken{
		TokenName:  types.HBSV,
		EthAddress: "0x14007c545e6664c8370f27fa6b99dc830e6510a6",
	}
}

func init() {
	initAccountList()
	initRootTokenMap()
}
