package config

import (
	"context"

	conf "github.com/carlware/go-common/config"
	"github.com/carlware/go-common/log"
	"github.com/carlware/go-common/sentry"
	"github.com/prometheus/common/version"
	"go.uber.org/zap"
)

var Conf = &Configuration{}

func InitConfig(cfgFile string) {
	if err := conf.Load(Conf, "CONDO", cfgFile); err != nil {
		log.Bg().Fatal("Unable load config", zap.Error(err))
	}
}

func InitLogger(ctx context.Context, cfg *Configuration) {
	log.Setup(ctx, log.Options{
		Debug:       cfg.Debug.Enable,
		LogLevel:    cfg.Logger.LogLevel,
		Encoding:    cfg.Logger.Encoding,
		AppName:     "condo-admin",
		Environment: cfg.Environment,
		Version:     version.Version,
		Revision:    version.Revision,
		SentryDSN:   cfg.Logger.Sentry,
	})
	sentry.Setup(cfg.Logger.Sentry, cfg.Environment, cfg.Debug.Enable)
}
