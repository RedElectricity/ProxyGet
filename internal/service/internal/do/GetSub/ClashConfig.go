package GetSub

import (
	"ProxyGet/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"gopkg.in/yaml.v3"
)

func GetClashConfig(url string) []model.ProxyRaw {
	Slice := make([]model.ProxyRaw, 5)
	r, err := g.Client().Get(
		gctx.New(),
		url,
	)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	ProxysRaw := make(map[interface{}]interface{})

	if err1 := yaml.Unmarshal(r.ReadAll(), &ProxysRaw); err1 != nil {
		panic(err1)
	}

	if ProxysRaw["proxies"] != nil {
		for _, i := range ProxysRaw["proxies"].([]map[interface{}]interface{}) {
			Slice = append(Slice, model.ProxyRaw{
				Server:   i["server"].(string),
				Port:     i["port"].(string),
				Type:     i["type"].(string),
				Cipher:   i["cipher"].(string),
				Password: i["password"].(string),
				IsUdp:    i["udp"].(bool),
			})
		}
	}

	return Slice
}
