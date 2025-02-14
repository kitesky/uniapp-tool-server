package config

type Config struct {
	App       App       `mapstructure:"app" yaml:"app"`
	Redis     Redis     `mapstructure:"redis" yaml:"redis"`
	Database  Database  `mapstructure:"database" yaml:"database"`
	WechatPay WechatPay `mapstructure:"wechat-pay" yaml:"wechat-pay"`
	Alipay    Alipay    `mapstructure:"alipay" yaml:"alipay"`
}

// APP配置
type App struct {
	Name     string `mapstructure:"name" yaml:"name"`
	Host     string `mapstructure:"host" yaml:"host"`
	Port     string `mapstructure:"port" yaml:"port"`
	Mode     string `mapstructure:"mode" yaml:"mode"`
	Url      string `mapstructure:"url" yaml:"url"`
	AssetUrl string `mapstructure:"asset-url" yaml:"asset-url"`
}

// Redis配置
type Redis struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" yaml:"port"`
	Password string `mapstructure:"password" yaml:"password"`
	DB       int    `mapstructure:"db" yaml:"db"`
}

// 数据库配置
type Database struct {
	Host         string `mapstructure:"host" yaml:"host"`
	Port         string `mapstructure:"port" yaml:"port"`
	User         string `mapstructure:"user" yaml:"user"`
	Password     string `mapstructure:"password" yaml:"password"`
	Database     string `mapstructure:"database" yaml:"database"`
	Prefix       string `mapstructure:"prefix" yaml:"prefix"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" yaml:"max-open-conns"`
}

type WechatPay struct {
	Appid               string `mapstructure:"appid" yaml:"appid"`
	AppSecret           string `mapstructure:"app_secret" yaml:"app_secret"`
	MchId               string `mapstructure:"mch_id" yaml:"mch_id"`
	SerialNo            string `mapstructure:"serial_no" yaml:"serial_no"`
	ApiV3Key            string `mapstructure:"api_v3_key" yaml:"api_v3_key"`
	ApiclientKeyContent string `mapstructure:"apiclient_key_content" yaml:"apiclient_key_content"`
}

type Alipay struct {
	Appid                   string `mapstructure:"appid" yaml:"appid"`
	PrivateKey              string `mapstructure:"private_key" yaml:"private_key"`
	AppPublicCertContent    string `mapstructure:"app_public_cert_content" yaml:"app_public_cert_content"`
	AlipayRootCertContent   string `mapstructure:"alipay_root_cert_content" yaml:"alipay_root_cert_content"`
	AlipayPublicCertContent string `mapstructure:"appid" yaml:"alipay_public_cert_content"`
}
