package service

import (
	"article-db/model"
	"article-db/pb"
	"context"
	"trpc.group/trpc-go/trpc-go/log"
)

func (ServiceImpl) GetArticleById(ctx context.Context, req *pb.GetArticleByIdRequest) (*pb.GetArticleByIdResponse, error) {
	log.Info("要查询的文章id:", req.GetId())
	article, err := model.GetArticleById(uint(req.GetId()))
	if err != nil {
		return &pb.GetArticleByIdResponse{
			Base: &pb.BaseResponse{
				Code: 1,
				Msg:  "未找到对应文章",
			},
			Id:      0,
			Title:   "",
			Content: "",
		}, nil
	}
	return &pb.GetArticleByIdResponse{
		Base: &pb.BaseResponse{
			Code: 0,
			Msg:  "",
		},
		Id:      uint64(article.ID),
		Title:   article.Title,
		Content: article.Content,
	}, nil
}
