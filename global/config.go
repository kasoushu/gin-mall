package global

type Config struct {
	Server Server `mapstructure:"server"`
	Upload Upload `mapstructure:"upload"`
	RedisServer RedisServer `mapstructure:"redis_server"`
	MysqlServer MysqlServer `mapstructure:"mysql_server"`
}
//server:
//address:
//port:
//
//upload:
//savepath: "g:/img/picture"
//
//redis_server:
//address: "localhost:6378"
//password: ""
//database: 0
//
//mysql_server:
//address: "localhost:6378"
//password: ""
//database: 0

type Server struct {
	Address string `mapstructure:"address"`
	Port string		`mapstructure:"port"`
}
type Upload struct {
	Path string			`mapstructure:"savepath"`
}
type RedisServer struct {
	Address string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	Database int `mapstructure:"databse"`
}
type MysqlServer struct {
	Address string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	Database int `mapstructure:"database"`

}