package internal

type (
	Entity struct {
		Elements    map[string][]Server `mapstructure:"elements"`
		Description string              `mapstructure:"description"`
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
