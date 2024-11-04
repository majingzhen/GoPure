package config

// Configuration 总配置结构体
type Configuration struct {
	Server     ServerConfig     `mapstructure:"server"`
	Logger     LoggerConfig     `mapstructure:"logger"`
	Session    SessionConfig    `mapstructure:"session"`
	Datasource DatasourceConfig `mapstructure:"datasource"`
	Upload     UploadConfig     `mapstructure:"upload"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Model     string `mapstructure:"model"`
	Port      int    `mapstructure:"port"`
	ImagePath string `mapstructure:"image_path"`
}

// UploadConfig 上传配置
type UploadConfig struct {
	ImagePath    string   `mapstructure:"image_path"`
	AllowedTypes []string `mapstructure:"allowed_types"`
	MaxSize      int64    `mapstructure:"max_size"`
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level    int    `mapstructure:"level"`
	FilePath string `mapstructure:"file_path"`
}

// SessionConfig 会话配置
type SessionConfig struct {
	Expire int    `mapstructure:"expire"`
	Secret string `mapstructure:"secret"`
}

// DatasourceConfig 数据库配置
type DatasourceConfig struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	DBName      string `mapstructure:"db_name"`
	TablePrefix string `mapstructure:"table_prefix"`
	LogMode     bool   `mapstructure:"log_mode"`
}
