package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"t-blog-back/pkg/setting"
)

var Tlog = logrus.New()

func init() {
	appCgf, err := setting.Cfg.GetSection("app")
	if err != nil {
		fmt.Println("error")
	} else {
		logFile, err := os.Open(appCgf.Key("RUNTIME_PATH").MustString("./runtime/log"))
		if err == nil {
			Tlog.SetOutput(logFile)
		}
	}
}



