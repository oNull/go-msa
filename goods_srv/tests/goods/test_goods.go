package main

import (
	"context"
	"fmt"
	"goods_srv/proto"
	"google.golang.org/grpc"
)

var goodsClient proto.GoodsClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	goodsClient = proto.NewGoodsClient(conn)
}

//func TestGetGoodsList() {
//	for i := 0; i < 10; i++ {
//		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
//			Password: "pengyuan123",
//			Mobile:   fmt.Sprintf("1760755170%d", i),
//			NickName: fmt.Sprintf("pengyuan%d", i),
//		})
//
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(rsp.Id)
//	}
//}

func TestGetGoodsList() {

	rsp, err := goodsClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		//PriceMin: 10,
		//KeyWords: "烟台红富士",
		TopCategory: 130358,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(rsp.Total)

	for _, good := range rsp.Data {
		fmt.Println(good.Name, good.ShopPrice)
	}
}

func main() {
	Init()
	TestGetGoodsList()
	conn.Close()
}
