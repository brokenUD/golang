package main

import (
	"fmt"
	"golang.org/x/net/context" //"context"标准包
	"google.golang.org/grpc"
	pb "brokenUd/golang/golang_biji/proto/article"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9005", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
	client := pb.NewArticleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()
	response, err := client.ArticleList(ctx, &pb.ArticleListRequest{Sort: pb.ARTICLE_SORT_ORDER_BY_TIME})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

}
