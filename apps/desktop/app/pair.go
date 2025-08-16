package app

import (
	"alat/pkg/core/pbuf"
	"fmt"
)

func (app *App) HandlePairRequest(req *pbuf.PairRequest) {
	fmt.Println("Pair request received", req)
}

func (app *App) HandlePairResponse(res *pbuf.PairResponse) {
	fmt.Println("Pair response received", res)
}
