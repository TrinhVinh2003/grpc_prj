package main

import (
	"context"
	"grpc-prj/api/articlepb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.NewClient("localhost:50069", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer cc.Close()
	client := articlepb.NewArticleServiceClient(cc)
	// insertArticle(client, "Caa thu vinh that dzai", "https//sfsfsfs", "Address 3", "sdfsdfsdfsdfsdfsd")
	// readArticle(client, 3)
	updateArticle(client, 2, "Caa thu vinh that dzai", "https//sfsfsfs", "Address 3", "sdfsdfsdfsdfsdfsd")

	// log.Println("service client &f", client)
}

func insertArticle(cli articlepb.ArticleServiceClient, title string, link string, image string, description string) {
	req := &articlepb.InsertRequest{
		Article: &articlepb.Article{
			Title:       title,
			Link:        link,
			Image:       image,
			Description: description,
		},
	}
	resp, err := cli.Insert(context.Background(), req)

	if err != nil {
		log.Printf("call insert err %v\n", err)
		return
	}

	log.Printf("insert response %+v\n", resp)
}
func readArticle(cli articlepb.ArticleServiceClient, id uint) {
	req := &articlepb.ReadRequest{Id: int32(id)}
	resp, err := cli.Read(context.Background(), req)
	if err != nil {
		log.Printf("call read err %v\n", err)
		return
	}

	log.Printf("read response %+v\n", resp.GetArticle())
}
func updateArticle(cli articlepb.ArticleServiceClient, id uint, title string, link string, image string, description string) {
	rep := &articlepb.UpdateRequest{
		Newarticle: &articlepb.Article{
			Id:          int32(id),
			Title:       title,
			Link:        link,
			Image:       image,
			Description: description,
		},
	}
	resp, err := cli.Update(context.Background(), rep)
	if err != nil {
		log.Printf("call update err %v\n", err)
		return
	}

	log.Printf("update response %+v\n", resp.GetArticle())
}
