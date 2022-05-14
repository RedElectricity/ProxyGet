package Proxy

import (
	"ProxyGet/internal/model"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/xiaoqidun/qqwry"
	"net"
	"time"
)

func SpeedTest() {
	var Proxys []model.ProxyRaw
	err := g.DB("main").Model("RawProxy").Scan(&Proxys)
	if err != nil {
		panic(err)
	}
	_timeout, err1 := g.Cfg().Get(gctx.New(), "Proxy.SpeedTest.Timeout")
	if err1 != nil {
		panic(err1)
	}
	timeout := _timeout.Duration() * time.Second

	var SortProxys []model.SortProxy

	for _, proxy := range Proxys {
		start := time.Now()
		if proxy.IsUdp == true {
			_, err2 := net.DialTimeout("udp", fmt.Sprintf(proxy.Server+":"+proxy.Port), timeout)
			if err2 == nil {
				SortProxys = append(SortProxys, model.SortProxy{
					Server:   proxy.Server,
					Port:     proxy.Port,
					Type:     proxy.Type,
					Cipher:   proxy.Cipher,
					Password: proxy.Password,
					IsUdp:    proxy.IsUdp,
					Country:  GetCountry(proxy.Server),
					Ping:     time.Now().Sub(start).String(),
				})
			}
		} else {
			_, err2 := net.DialTimeout("tcp", fmt.Sprintf(proxy.Server+":"+proxy.Port), timeout)
			if err2 == nil {
				SortProxys = append(SortProxys, model.SortProxy{
					Server:   proxy.Server,
					Port:     proxy.Port,
					Type:     proxy.Type,
					Cipher:   proxy.Cipher,
					Password: proxy.Password,
					IsUdp:    proxy.IsUdp,
					Country:  GetCountry(proxy.Server),
					Ping:     time.Now().Sub(start).String(),
				})
			}
		}
	}
	_, err3 := g.DB("main").Model("Proxys").Insert(SortProxys)
	if err3 != nil {
		panic(err3)
	}

	_, err4 := g.DB("main").Model("RawProxy").Delete()
	if err4 != nil {
		panic(err4)
	}
}

func GetCountry(ip string) string {
	if err := qqwry.LoadFile("resource/qqwry.dat"); err != nil {
		panic(err)
	}
	city, _, err := qqwry.QueryIP(ip)
	if err != nil {
		return "Unknown"
	} else {
		return city
	}
}
