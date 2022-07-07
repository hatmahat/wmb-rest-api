package delivery

import (
	"go-api-with-gin2/config"
	"go-api-with-gin2/delivery/controller"
	"go-api-with-gin2/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(&appConfig)
	repoManager := manager.NewRepoManager(infra)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := appConfig.Url
	return &appServer{
		useCaseManager: useCaseManager,
		engine:         r,
		host:           host,
	}
}

func (a *appServer) initControllers() {
	controller.NewMenuController(a.engine, a.useCaseManager.CreateMenuUseCase(), a.useCaseManager.ListMenuUseCase(), a.useCaseManager.DeleteMenuUseCase(), a.useCaseManager.UpdateMenuUseCase())
	controller.NewMenuPriceController(a.engine, a.useCaseManager.CreateMenuPriceUseCase(), a.useCaseManager.ListMenuPriceUseCase(), a.useCaseManager.DeleteMenuPriceUseCase(), a.useCaseManager.UpdateMenuPriceUseCase())
	controller.NewBillDetailController(a.engine, a.useCaseManager.CreateBillDetailUseCase(), a.useCaseManager.ListBillDetailUseCase())
	controller.NewBillController(a.engine, a.useCaseManager.CreateBillUseCase(), a.useCaseManager.ListBillUseCase())
	controller.NewTableController(a.engine, a.useCaseManager.CreateTableUseCase(), a.useCaseManager.ListTableUseCase(), a.useCaseManager.DeleteTableUseCase(), a.useCaseManager.UpdateTableUseCase())
	controller.NewTransTypeController(a.engine, a.useCaseManager.CreateTransTypeUseCase(), a.useCaseManager.ListTransTypeUseCase(), a.useCaseManager.DeleteTransTypeUseCase(), a.useCaseManager.UpdateTransTypeUseCase())
	controller.NewCustomerController(a.engine, a.useCaseManager.CreateCustomerUseCase(), a.useCaseManager.ListCustomerUseCase(), a.useCaseManager.UpdateCustomerUseCase())
	controller.NewDiscountController(a.engine, a.useCaseManager.CreateDiscountUseCase(), a.useCaseManager.ListDiscountUseCase(), a.useCaseManager.DeleteDiscountUseCase(), a.useCaseManager.UpdateDiscountUseCase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
