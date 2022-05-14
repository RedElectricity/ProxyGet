package model

type ProxyRaw struct {
	Server   string
	Port     string
	Type     string
	Cipher   string
	Password string
	IsUdp    bool
}

type SortProxy struct {
	Server   string
	Port     string
	Type     string
	Cipher   string
	Password string
	IsUdp    bool
	Country  string
	Ping     string
}
