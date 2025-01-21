package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	// pb "messagebox/article/article"
	pb "brokenUd/golang/golang_biji/proto/article"
	"net"
	"runtime"
)

type server struct {
}

type ArticleData struct {
	Id      uint32 `db : "id" `
	Title   string `db : "title" `
	Content []byte `db : "content" `
	Viewnum uint32 `db : "viewnum" `
	Time    uint32 `db : "time" `
}

// func (s *server) mustEmbedUnimplementedArticleServer() {}

func (s *server) ArticleList(ctx context.Context, request *pb.ArticleListRequest) (response *pb.ArticleListResponse, err error) {
	orderby := "time"
	if request.Sort == pb.ARTICLE_SORT_ORDER_BY_TIME {
		orderby = "time"
	}

	if request.Sort == pb.ARTICLE_SORT_ORDER_BY_CLICK {
		orderby = "viewnum"
	}
	host := "localhost:33060"
	dbname := "messagebox"
	username := "root"
	password := "47810058k2"

	db, err := sqlx.Open("mysql", username+":"+password+"@tcp("+host+")/"+dbname+"?charset=utf8")
	SimplePanic(err)
	defer db.Close()

	//定义一个结构重新赋值
	var list1 []*ArticleData
	var list2 []*pb.ArticleStruct
	sql := fmt.Sprintf("select id, title, content, viewnum, time from msg_article order by %s desc limit 10", orderby)
	err = db.Select(&list1, sql)
	SimplePanic(err)

	for _, v := range list1 {
		var list *pb.ArticleStruct = &pb.ArticleStruct{
			Id:      v.Id,
			Title:   v.Title,
			Content: v.Content,
			Viewnum: v.Viewnum,
			Time:    v.Time,
		}
		list2 = append(list2, list)
	}
	return &pb.ArticleListResponse{Code: pb.RPC_RESULT_RESULT_SUCCESS, Msg: "", List: list2}, nil
}

func main() {
	//绑定和监听一个端口
	listen, err := net.Listen("tcp", "127.0.0.1:9005")
	if err != nil {
		fmt.Println(err)
	}

	//创建 gRPC 服务器的一个实例(初始化配置)
	s := grpc.NewServer()
	//绑定一个实现该gRPC的服务
	// a := new(pb.ArticleServer)
	pb.RegisterArticleServer(s, &server{})
	//在gRPC上注册服务
	reflection.Register(s)

	//启动服务(阻塞等待客户端连接)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func SimplePanic(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line, err)
		runtime.Goexit()
	}
}
