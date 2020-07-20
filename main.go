package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/wmsx/store_svc/handler"
	"github.com/wmsx/store_svc/models"
	"github.com/wmsx/store_svc/setting"

	proto "github.com/wmsx/store_svc/proto/store"
)

const name = "wm.sx.svc.store"

func main() {
	service := micro.NewService(
		micro.Name(name),
		micro.Flags(
			&cli.StringFlag{
				Name:    "env",
				Usage:   "指定运行环境",
				Value:   "dev",
				EnvVars: []string{"RUN_ENV"},
			},
		),
	)

	var env string
	service.Init(
		micro.Action(func(c *cli.Context) error {
			env = c.String("env")
			return nil
		}),

		micro.BeforeStart(func() (err error) {
			if err = setting.SetUp(name, env); err != nil {
				return err
			}
			if err =  models.SetUp(); err != nil {
				return err
			}
			return nil
		}),
		micro.AfterStop(func() error {
			models.CloseDB()
			return nil
		}),
	)

	_ = proto.RegisterStoreHandler(service.Server(), new(handler.StoreHandler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
