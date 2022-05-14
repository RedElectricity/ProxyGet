package GetSub

import (
	"ProxyGet/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"reflect"
)

func GetProxyFromWeb() {
	Slice := make([]model.ProxyRaw, 5)
	// GetFromAPI
	Slice = append(Slice, S5ListMTProto()...)
	Slice = append(Slice, S5ListMTProto()...)
	// GetFromClashSubConfig
	type ClashConfDB struct {
		URL string
	}
	var ConfigUrls []ClashConfDB
	err := g.DB("main").Model("ClashSub").Fields("URL").Scan(&ConfigUrls)
	if err != nil {
		panic(err)
	}
	for _, l := range ConfigUrls {
		Slice = append(Slice, GetClashConfig(l.URL)...)
	}

	SortResult := make([]model.ProxyRaw, 5)

	for i := 0; i < len(Slice); i++ {
		state := false
		for j := i + 1; j < len(Slice); j++ {
			if j > 0 && reflect.DeepEqual(Slice[i], Slice[j]) {
				state = true
				break
			}
		}
		if !state {
			SortResult = append(SortResult, Slice[i])
		}
	}

	_, err1 := g.DB("main").Model("RawProxy").Insert(SortResult)
	if err1 != nil {
		panic(err1)
	}
}
