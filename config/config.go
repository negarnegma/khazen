package config

import "time"

const (
	SentryTimeout time.Duration = 10

	InvalidArgsMessage string = "invalid args"
)

var Config = struct {
	Port    string `default:"3000"`
	LogFile string `default:"/var/log/khazen.log"`

	SentryDSN string

	AccessKey string `required:"true"`
	SecretKey string `required:"true"`

	MySQL struct {
		Host     string `default:"127.0.0.1"`
		User     string `default:"root"`
		Password string `required:"true"`
		Port     string `default:"3306"`
	}
}{}
