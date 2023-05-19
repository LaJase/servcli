package models

// type ServerGroup struct {
// 	Description string   `mapstructure:"description"`
// 	Servers     []Server `mapstructure:"servers"`
// }

type Entity struct {
	Description string              `mapstructure:"description"`
	Elements    map[string][]Server `mapstructure:"elements"`
}

type Server struct {
	Name  string `mapstructure:"name"`
	IsAWS bool   `mapstructure:"isaws"`
}

type ServerConfig struct {
	ServerList map[string]Entity `mapstructure:"server_list"`
}
