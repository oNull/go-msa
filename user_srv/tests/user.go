package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"user_srv/proto"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	userClient = proto.NewUserClient(conn)
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			Password: "pengyuan123",
			Mobile:   fmt.Sprintf("1760755170%d", i),
			NickName: fmt.Sprintf("pengyuan%d", i),
		})

		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 10,
	})

	if err != nil {
		panic(err)
	}

	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.Password)
		checkRsp, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          "pengyuan123",
			EncryptedPassword: user.Password,
		})

		if err != nil {
			panic(err)
		}

		fmt.Println(checkRsp.Success)
	}
}

func main() {
	Init()
	TestCreateUser()
	TestGetUserList()
	conn.Close()
}
