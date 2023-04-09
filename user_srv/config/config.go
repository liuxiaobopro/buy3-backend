package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}

type SrvConfig struct {
	Ip   string `mapstructure:"ip" json:"ip"`
	Port string `mapstructure:"port" json:"port"`
}

type Conf struct {
	Name       string       `mapstructure:"name" json:"name"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	Srv        SrvConfig    `mapstructure:"user_srv" json:"user_srv"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
}
