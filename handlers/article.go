package handlers

import (
	"context"
	"strconv"
	"time"

	"net/url"

	pb "UwU/proto"
)

type Article struct {
	ID        uint32    `json:"id"`
	AuthorID  uint32    `json:"author_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (a Article) ToGRPC() *pb.Article {
	return &pb.Article{
		Id:        a.ID,
		AuthorId:  a.AuthorID,
		Title:     a.Title,
		Body:      a.Body,
		CreatedAt: a.CreatedAt.String(),
	}
}

type ArticleInput struct {
	AuthorID uint32 `json:"author_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

func (ai *ArticleInput) FromGRPC(b *pb.ArticleInput) {
	ai.AuthorID = b.AuthorId
	ai.Body = b.Body
	ai.Title = b.Title
}

type ArticleParams struct {
	Limit  uint32
	Page   uint32
	Author string
	Query  string
}

func (ap *ArticleParams) FromGRPC(b *pb.ArticleParam) {
	ap.Author = b.GetAuthor()
	ap.Limit = b.GetLimit()
	if ap.Limit == 0 {
		ap.Limit = 10
	}
	ap.Page = b.GetLimit()
	if ap.Page == 0 {
		ap.Page = 1
	}
	ap.Query = b.GetQuery()
}

func (a ArticleParams) CacheTag() string {
	return (strconv.FormatUint(uint64(a.Limit), 10) + "&" + strconv.FormatUint(uint64(a.Limit), 10) + "&" + url.QueryEscape(a.Author) + "&" + url.QueryEscape(a.Query))
}

func (dp Dependencies) GetAllArticles(in *pb.ArticleParam, stream pb.UwU_GetAllArticlesServer) error {
	ctx := context.TODO()
	params := ArticleParams{}
	params.FromGRPC(in)
	result := FetchArticlesFromCache(ctx, dp, params.CacheTag())
	var err error

	if len(result) == 0 {
		result, err = FetchArticlesFromDB(ctx, dp, params)
		if err != nil {
			return err
		}
		err = PersistArticlesCache(ctx, dp, params, result)
		if err != nil {
			return err
		}
	}

	for _, item := range result {
		if err = stream.Send(item.ToGRPC()); err != nil {
			return err
		}
	}
	return err
}

func (db Dependencies) GetArticle(ctx context.Context, in *pb.ArticleId) (*pb.Article, error) {
	return FetchArticle(ctx, db, in.Id).ToGRPC(), nil
}

func (db Dependencies) CreateArticle(ctx context.Context, in *pb.ArticleInput) (*pb.Nothing, error) {
	newArticle := new(ArticleInput)
	newArticle.FromGRPC(in)

	return &pb.Nothing{}, PersistArticle(ctx, db, *newArticle)
}
