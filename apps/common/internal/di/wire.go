// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"app/apps/common/internal/dao"
	"app/apps/common/internal/service"
	"app/apps/common/internal/server/grpc"
	"app/apps/common/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
