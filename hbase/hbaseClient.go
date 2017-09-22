package main

import (
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"context"
	"log"
)

func main() {
	client := gohbase.NewClient("192.168.1.40:2181")
	//getRequest, err := hrpc.NewGetStr(context.Background(), "common_data_history", "100000596_656005413_02")
	//if(err!=nil){
	//	log.Println("getRequest err")
	//}
	//getRsp, err := client.Get(getRequest)
	//if(err!=nil){
	//	log.Println("getRsp err")
	//}else{
	//	for _,cell := range getRsp.Cells {
	//		log.Println(string(cell.Family))
	//		log.Println(string(cell.Qualifier))
	//		log.Println(string(cell.Value))
	//	}
	//}


	//sacn
	scanRequest, _ := hrpc.NewScanRangeStr(context.Background(), "common_data_history","100000596_656005413_02","100000596_656005842_02")
	//log.Println(string(scanRequest.StartRow()))
	//log.Println(string(scanRequest.StopRow()))
	scanRsp := client.Scan(scanRequest)
	for  {
		result,err :=scanRsp.Next()
		if(err!=nil){
			break
		}
		for _,cell := range result.Cells {
			log.Println(string(cell.Family))
			log.Println(string(cell.Qualifier))
			log.Println(string(cell.Value))
		}

	}
}
