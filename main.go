package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.wdf.sap.corp/cs-v2/gbaas-service/pkg/app"
)

func main() {
	fmt.Println(errors.New("Import random public pkg"))
	app.Crash(fmt.Errorf("crash method from gbaas-service"))
}
