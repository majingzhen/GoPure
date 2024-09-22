package global

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var (
	// GormDao db对象
	GormDao *gorm.DB
	// Viper 配置对象
	Viper *viper.Viper
	// Logger 日志对象
	Logger *zap.Logger
)

// InitViper 初始化配置
func InitViper() {
	Viper = viper.New()
	Viper.AddConfigPath(".")           // 添加配置文件搜索路径，点号为当前目录
	Viper.AddConfigPath("./config")    // 添加多个搜索目录
	Viper.SetConfigType("yml")         // 如果配置文件没有后缀，可以不用配置
	Viper.SetConfigName("application") // 文件名，没有后缀
	// 读取配置文件
	if err := Viper.ReadInConfig(); err != nil {
		panic("读取配置文件错误")
	}
}

// InitDataSource 初始化数据库
func InitDataSource() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		Viper.GetString("datasource.username"),
		Viper.GetString("datasource.password"),
		Viper.GetString("datasource.host"),
		Viper.GetString("datasource.port"),
		Viper.GetString("datasource.db_name"))
	gcf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   Viper.GetString("datasource.table_prefix"), // 控制表前缀
			SingularTable: true,
		},
		Logger: logger.Default, // 控制是否sql输出，默认是不输出
	}
	if Viper.GetBool("datasource.log_mode") {
		gcf.Logger = NewGormLogger() // 使用zap进行日志输出
	}

	if tmp, err := gorm.Open(mysql.Open(dsn), gcf); err != nil {
		Logger.Error("MySQL启动异常", zap.Error(err))
		panic(err)
	} else {
		// 设置delete_at字段类型
		tmp.Set("gorm:softDelete", "is_del")
		//Logger.Info("Connect to database success")
		//// 全局禁用表名复数
		//tmp = tmp.Set("gorm:table_options", "ENGINE=InnoDB")
		//// 全局设置表前缀
		sqlDB, _ := tmp.DB()
		// 设置最大空闲连接数
		sqlDB.SetMaxIdleConns(10)
		// 设置最大打开的连接数
		sqlDB.SetMaxOpenConns(100)
		// 设置连接的最大可复用时间
		sqlDB.SetConnMaxLifetime(60 * time.Second)
		// DbList = make(map[string]*gorm.DB)
		// DbList[Viper.GetString("datasource.db_name")] = GormDao
		GormDao = tmp
	}
}

// InitLogger 初始化日志
func InitLogger() {
	logPath := Viper.GetString("logger.file_path")
	if logPath == "" {
		logPath = "./log/manager.log" // 如果未配置日志路径，则默认在项目根目录下创建log目录
	}
	// 设置日志文件的位置、文件名、最大大小、最大备份数量和压缩
	hook := lumberjack.Logger{
		Filename:   logPath, // 日志路径
		MaxSize:    128,     // MB
		MaxBackups: 30,
		MaxAge:     7, // days
		Compress:   true,
	}
	// 配置日志级别
	atomicLevel := zap.NewAtomicLevel()
	logLevel := Viper.GetInt32("logger.level")
	atomicLevel.SetLevel(zapcore.Level(logLevel))
	// 创建编码器
	// 设置日志格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// 创建core
	writer := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	).With([]zap.Field{})
	// 初始化logger
	Logger = zap.New(writer)
}
