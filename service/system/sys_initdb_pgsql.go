package system

import (
	"context"
	"errors"
	"github.com/gookit/color"
	"github.com/zouchangfu/QanLong/config"
	"github.com/zouchangfu/QanLong/utils"

	uuid "github.com/satori/go.uuid"
	"github.com/zouchangfu/QanLong/global"
)

type PgsqlInitHandler struct{}

func NewPgsqlInitHandler() *PgsqlInitHandler {
	return &PgsqlInitHandler{}
}

// WriteConfig pgsql 回写配置
func (h PgsqlInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Pgsql)
	if !ok {
		return errors.New("postgresql config invalid")
	}
	global.GVA_CONFIG.System.DbType = "pgsql"
	global.GVA_CONFIG.Pgsql = c
	global.GVA_CONFIG.JWT.SigningKey = uuid.NewV4().String()
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	return global.GVA_VP.WriteConfig()
}

func (h PgsqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h PgsqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for i := 0; i < len(inits); i++ {
		if inits[i].DataInserted(next) {
			color.Info.Printf(InitDataExist, Pgsql, inits[i].InitializerName())
			continue
		}
		if n, err := inits[i].InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Pgsql, inits[i].InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Pgsql, inits[i].InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Pgsql)
	return nil
}
