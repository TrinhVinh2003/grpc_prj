package service

import (
	"context"
	"fmt"
	"grpc-prj/api/articlepb"
	"grpc-prj/api/models"
	"grpc-prj/api/repository"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	Repo repository.ArticleRepository
	articlepb.ArticleServiceServer
}

func Newserver(repo repository.ArticleRepository) *server {
	return &server{Repo: repo}
}

func ConvertPbArticle2ArticleInfor(pbArticle *articlepb.Article) *models.Article {
	return &models.Article{
		Title:       pbArticle.Title,
		Link:        pbArticle.Link,
		Image:       pbArticle.Image,
		Description: pbArticle.Description,
	}
}
func ConvertArticleInfo2PbArticle(ci *models.Article) *articlepb.Article {
	return &articlepb.Article{
		Id:          int32(ci.ID),
		Title:       ci.Title,
		Link:        ci.Link,
		Image:       ci.Image,
		Description: ci.Description,
	}
}

func (s *server) Insert(ctx context.Context, req *articlepb.InsertRequest) (*articlepb.InsertResponse, error) {
	ci := ConvertPbArticle2ArticleInfor(req.Article)

	err := s.Repo.Insert(ci)
	if err != nil {
		resp := &articlepb.InsertResponse{
			StatusCode: -1,
			Message:    fmt.Sprintf("insert err %v", err),
		}
		return resp, nil

	}

	resp := &articlepb.InsertResponse{
		StatusCode: 1,
		Message:    "OK",
	}

	return resp, nil
}

func (s *server) Read(ctx context.Context, req *articlepb.ReadRequest) (*articlepb.ReadResponse, error) {
	ci, err := s.Repo.Read(uint(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Read phone %s err %v", req.GetId(), err)
	}
	return &articlepb.ReadResponse{
		Article: ConvertArticleInfo2PbArticle(ci),
	}, nil
}
func (s *server) Update(ctx context.Context, req *articlepb.UpdateRequest) (*articlepb.UpdateResponse, error) {
	ci := ConvertPbArticle2ArticleInfor(&articlepb.Article{})
	if err := s.Repo.Update(ci); err != nil {
		log.Fatal(err)
		return nil, err
	}
	updateArticle, err := s.Repo.Read(uint(req.GetNewarticle().Id))
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "try to read update contact %+v err %v", req.GetNewarticle(), err)
	}
	return &articlepb.UpdateResponse{
		Article: ConvertArticleInfo2PbArticle(updateArticle),
	}, nil

}
func (s *server) Delete(ctx context.Context, req *articlepb.DeleteRequest) (*articlepb.DeleteResponse, error) {
	err := s.Repo.Delete(uint(req.Id))
	if err != nil {
		return &articlepb.DeleteResponse{
			Status:  -1,
			Message: fmt.Sprintf("delete error %v", err),
		}, nil
	}

	return &articlepb.DeleteResponse{
		Status:  1,
		Message: fmt.Sprintf("delete article with ID %d successfully", req.GetId()),
	}, nil

}
