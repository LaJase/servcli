package internal

type (
	Entity struct {
		Description string              `mapstructure:"description"`
		Elements    map[string][]Server `mapstructure:"elements"`
	}

	Server struct {
		Name  string `mapstructure:"name"`
		IsAWS bool   `mapstructure:"isaws"`
	}

	ServerConfig struct {
		ServerList map[string]Entity `mapstructure:"server_list"`
		SshCommand string            `mapstructure:"ssh_command"`
	}
)
