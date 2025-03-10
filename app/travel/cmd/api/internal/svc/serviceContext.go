package svc

import (
	"looklook/app/order/cmd/rpc/order"
	"looklook/app/travel/cmd/api/internal/config"
	"looklook/app/travel/cmd/rpc/travel"
	"looklook/app/travel/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	//local
	Config config.Config

	//rpc
	UsercenterRpc usercenter.Usercenter
	TravelRpc     travel.Travel
	OrderRpc      order.Order

	//model
	HomestayModel         model.HomestayModel
	HomestayActivityModel model.HomestayActivityModel
	HomestayBusinessModel model.HomestayBusinessModel
	HomestayCommentModel  model.HomestayCommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		TravelRpc:     travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),

		HomestayModel:         model.NewHomestayModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		HomestayActivityModel: model.NewHomestayActivityModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		HomestayBusinessModel: model.NewHomestayBusinessModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		HomestayCommentModel:  model.NewHomestayCommentModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
