package handlers

import (
	"context"
	"time"

	"github.com/georgysavva/scany/pgxscan"
)

func FetchArticlesFromDB(ctx context.Context, db Dependencies, param ArticleParams) (result []Article) {
	query := "SELECT articles.id,articles.author_id,articles.title,articles.body,articles.created_at FROM articles"

	if param.Author != "" {
		query += " LEFT JOIN authors ON articles.author_id == authors.id AND authors.name == $3 LIMIT $1 OFFSET $2"
		pgxscan.Select(ctx, db.DB, &result, query, param.Limit, param.Limit*param.Page)
	} else {
		query += "LIMIT $1 OFFSET $2"
		pgxscan.Select(ctx, db.DB, &result, query, param.Limit, param.Limit*param.Page)
	}
	return
}

func FetchArticlesFromCache(ctx context.Context, db Dependencies, cacheTag string) (result []Article) {
	db.Cache.Get(ctx, cacheTag, &result)
	return
}

func PersistArticleCache(ctx context.Context, db Dependencies) error {
	return nil
}

func PersistArticlesCache(ctx context.Context, db Dependencies) error {
	return nil
}

func InvalidateCache(ctx context.Context, db Dependencies) error {
	return nil
}

func FetchArticle(ctx context.Context, db Dependencies, id uint) (result *Article, err error) {
	row, err := db.DB.Query(ctx, "SELECT id,author_id,title,body,created_at FROM articles WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return
	}
	err = pgxscan.ScanOne(&result, row)
	return
}

func PersistArticle(ctx context.Context, db Dependencies, dto ArticleDTO) (err error) {
	_, err = db.DB.Exec(ctx, "INSERT INTO articles(author_id,title,body,created_at) VALUES $1,$2,$3,$4", dto.AuthorID, dto.Title, dto.Body, time.Now())
	return
}
