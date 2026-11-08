package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-kratos/kratos/v2/config"
	fileKratos "github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
)

const remoteConfigSourceConfigFile = "config.yaml"

// NewFileConfigSource 创建一个本地文件配置源
func NewFileConfigSource(filePath string) config.Source {
	return fileKratos.NewSource(filePath)
}

// NewConfigProvider 创建一个配置
func NewConfigProvider(configPath string, sources ...string) (config.Config, error) {
	source := "local"
	if len(sources) > 1 {
		return nil, fmt.Errorf("config: expected one configuration source")
	}
	if len(sources) == 1 {
		source = sources[0]
	}
	switch source {
	case "local":
		// 本地模式只读本地，不创建远程客户端。 / Local mode never connects remotely.
		return config.New(config.WithSource(NewFileConfigSource(configPath))), nil
	case "nacos":
	default:
		return nil, fmt.Errorf("config: unsupported source %q (expected local or nacos)", source)
	}
	err, rc := LoadRemoteConfigSourceConfigs(configPath)
	if err != nil {
		log.Error("LoadRemoteConfigSourceConfigs: ", err.Error())
		return nil, err
	}

	if rc != nil && rc.GetType() == "nacos" && rc.GetNacos() != nil {
		rcs, err := NewProvider(rc)
		if err != nil {
			log.Error("NewProvider: ", err.Error())
			return nil, err
		}
		return config.New(
			config.WithSource(
				rcs,
			),
		), nil
	}

	return nil, fmt.Errorf("config: nacos mode requires config.type=nacos and config.nacos in %s", filepath.Join(configPath, remoteConfigSourceConfigFile))
}

// LoadBootstrapConfig 加载程序引导配置
func LoadBootstrapConfig(configPath string, sources ...string) error {
	cfg, err := NewConfigProvider(configPath, sources...)
	if err != nil {
		return err
	}

	if err = cfg.Load(); err != nil {
		return err
	}

	initBootstrapConfig()

	if err = scanConfigs(cfg); err != nil {
		return err
	}

	return nil
}

func scanConfigs(cfg config.Config) error {
	initBootstrapConfig()

	for _, c := range configList {
		if err := cfg.Scan(c); err != nil {
			return err
		}
	}
	return nil
}

// LoadRemoteConfigSourceConfigs 加载远程配置源的本地配置
func LoadRemoteConfigSourceConfigs(configPath string) (error, *conf.RemoteConfig) {
	configPath = filepath.Join(configPath, remoteConfigSourceConfigFile)
	if !pathExists(configPath) {
		return nil, nil
	}

	cfg := config.New(
		config.WithSource(
			NewFileConfigSource(configPath),
		),
	)
	defer func(cfg config.Config) {
		if err := cfg.Close(); err != nil {
			panic(err)
		}
	}(cfg)

	var err error

	if err = cfg.Load(); err != nil {
		return err, nil
	}

	// 引导配置独立解析，避免污染业务配置或复用上次加载的全局值。
	// Decode bootstrap separately to avoid stale global or local business values.
	var bootstrap conf.Bootstrap
	if err = cfg.Scan(&bootstrap); err != nil {
		return err, nil
	}
	return nil, bootstrap.Config
}

// pathExists 判断路径是否存在
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
