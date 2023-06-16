package gormfx

import (
	"context"
	"github.com/guoyk93/ufx"
	"gorm.io/gorm"
)

func NewConfig() *gorm.Config {
	return &gorm.Config{}
}

func NewClient(d gorm.Dialector, c *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(d, c)
}

func AddCheckerForClient(db *gorm.DB, v ufx.Probe) {
	v.AddChecker("gorm", func(ctx context.Context) error {
		return db.WithContext(ctx).Select("SELECT 1").Error
	})
}
