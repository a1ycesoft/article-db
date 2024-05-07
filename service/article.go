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

func (ServiceImpl) InsertArticle(ctx context.Context, req *pb.InsertArticleRequest) (*pb.InsertArticleResponse, error) {
	//	log.Info(req.Title, req.Content)
	err := model.InsertArticle(&req.Title, &req.Content)
	if err != nil {
		return &pb.InsertArticleResponse{Base: &pb.BaseResponse{
			Code: 1,
			Msg:  "插入文章失败 " + err.Error(),
		}}, nil
	}
	return &pb.InsertArticleResponse{Base: &pb.BaseResponse{
		Code: 0,
		Msg:  "插入文章成功",
	}}, nil
}

func (ServiceImpl) QueryArticleByKeyword(ctx context.Context, req *pb.QueryArticleByKeywordRequest) (*pb.QueryArticleByKeywordResponse, error) {
	articles, err := model.QueryArticleByKeyword(req.GetKeyword(), req.GetPageNum(), req.GetPageSize())
	if err != nil {
		return &pb.QueryArticleByKeywordResponse{
			Base: &pb.BaseResponse{
				Code: 1,
				Msg:  "服务器内部错误,查询失败",
			},
			Articles: nil,
		}, nil
	}
	arr := make([]*pb.Article, len(articles))
	for i, v := range articles {
		arr[i] = &pb.Article{
			Id:      int64(v.ID),
			Title:   v.Title,
			Content: v.Content,
		}
	}
	return &pb.QueryArticleByKeywordResponse{
		Base: &pb.BaseResponse{
			Code: 0,
			Msg:  "查询成功",
		},
		Articles: arr,
	}, nil
}
