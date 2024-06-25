package appinfoHandlers

import (
	"github.com/MarkTBSS/076_Appinfo_Module/config"
	"github.com/MarkTBSS/076_Appinfo_Module/modules/appinfo/appinfoUsecases"
)

type IAppinfoHandler interface {
}

type appinfoHandler struct {
	cfg            config.IConfig
	appinfoUsecase appinfoUsecases.IAppinfoUsecase
}

func AppinfoHandler(cfg config.IConfig, appinfoUsecase appinfoUsecases.IAppinfoUsecase) IAppinfoHandler {
	return &appinfoHandler{
		cfg:            cfg,
		appinfoUsecase: appinfoUsecase,
	}
}
