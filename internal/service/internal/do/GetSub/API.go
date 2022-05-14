package GetSub

import (
	"ProxyGet/internal/model"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func S5ListMTProto() []model.ProxyRaw {
	type JsonStruct struct {
		Host    string `json:"host"`
		Port    string `json:"port"`
		Secret  string `json:"secret"`
		Country string `json:"country"`
		Up      string `json:"up"`
		Down    string `json:"down"`
		Uptime  string `json:"uptime"`
		Unix    string `json:"unix"`
		Ping    string `json:"ping"`
	}
	Slice := make([]model.ProxyRaw, 5)
	r, err := g.Client().Get(
		gctx.New(),
		"https://github.com/hookzof/socks5_list/raw/master/tg/mtproto.json",
	)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	j, _ := gjson.LoadJson(r.ReadAllString())
	Proxys := make([]JsonStruct, 5)
	if err := j.Scan(Proxys); err != nil {
		panic(err)
	}

	for _, s := range Proxys {
		Slice = append(Slice, model.ProxyRaw{
			Server:   s.Host,
			Port:     s.Port,
			Type:     "MTProto",
			Cipher:   "",
			Password: s.Secret,
			IsUdp:    false,
		})
	}

	return Slice
}

func S5ListSock() []model.ProxyRaw {
	type JsonStruct struct {
		Unix    string `json:"unix"`
		IP      string `json:"ip"`
		Port    string `json:"port"`
		Country string `json:"country"`
		Ping    string `json:"ping"`
	}
	Slice := make([]model.ProxyRaw, 5)
	r, err := g.Client().Get(
		gctx.New(),
		"https://github.com/hookzof/socks5_list/raw/master/tg/mtproto.json",
	)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	j, _ := gjson.LoadJson(r.ReadAllString())
	Proxys := make([]JsonStruct, 5)
	if err := j.Scan(Proxys); err != nil {
		panic(err)
	}

	for _, s := range Proxys {
		Slice = append(Slice, model.ProxyRaw{
			Server:   s.IP,
			Port:     s.Port,
			Type:     "Sock5",
			Cipher:   "",
			Password: "",
			IsUdp:    false,
		})
	}

	return Slice
}
