package handlers

import (
	"context"
	"time"

	"encoding/json"

	"github.com/georgysavva/scany/pgxscan"
)

const exp = 1 * time.Hour

func FetchArticlesFromDB(ctx context.Context, db Dependencies, param ArticleParams) (result []Article, err error) {

	if (param.Author != "") && (param.Query != "") {
		query := `
		SELECT 
		   articles.id,articles.author_id,articles.title,articles.body,articles.created_at 
		FROM articles 
		LEFT JOIN authors ON 
		articles.author_id=authors.id 
		WHERE authors.name=$3 
		AND articles.article_fts @@ to_tsquery('english',$4) 
		ORDER BY created_at DESC LIMIT $1 OFFSET $2`
		err = pgxscan.Select(ctx, db.DB, &result, query, param.Limit, param.Limit*(param.Page-1), param.Author, param.Query)
	} else if param.Query != "" {
		query := `
		SELECT 
		   id,author_id,title,body,created_at 
		FROM articles 
		WHERE article_fts @@ to_tsquery('english',$3)
		ORDER BY created_at DESC LIMIT $1 OFFSET $2 `
		err = pgxscan.Select(ctx, db.DB, &result, query, param.Limit, param.Limit*(param.Page-1), param.Query)
	} else if param.Author != "" {
		query := `
		SELECT 
			articles.id,articles.author_id,articles.title,articles.body,articles.created_at 
		FROM articles
		INNER JOIN authors ON 
		articles.author_id=authors.id 
		AND authors.name=$3 
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2`
		err = pgxscan.Select(ctx, db.DB, &result, query, param.Limit, param.Limit*(param.Page-1), param.Author)
	} else {
		query := `
		SELECT 
			id,author_id,title,body,created_at 
		FROM articles
		ORDER BY created_at DESC LIMIT $1 OFFSET $2 `
		err = pgxscan.Select(ctx, db.DB, &result, query, param.Limit, param.Limit*(param.Page-1))
	}
	return
}

func FetchArticlesFromCache(ctx context.Context, db Dependencies, cacheTag string) (result []Article) {
	getCache(ctx, db, cacheTag, &result)
	return
}

func PersistArticleCache(ctx context.Context, db Dependencies) (err error) {
	return nil
}

func PersistArticlesCache(ctx context.Context, db Dependencies, param ArticleParams, data []Article) (err error) {
	err = setCache(ctx, db, param.CacheTag(), data)
	return
}

func InvalidateCache(ctx context.Context, db Dependencies) {
	db.Cache.FlushDB(ctx)
}

func FetchArticle(ctx context.Context, db Dependencies, id uint32) (result *Article) {
	res := []Article{}

	_ = pgxscan.Select(ctx, db.DB, &res, "SELECT id,author_id,title,body,created_at FROM articles WHERE id=$1 LIMIT 1", id)
	if len(res) == 0 {
		result = &Article{}
	} else {
		result = &res[0]
	}
	return
}

func PersistArticle(ctx context.Context, db Dependencies, dto ArticleInput) (err error) {
	_, err = db.DB.Exec(ctx, "INSERT INTO articles(author_id,title,body,created_at) VALUES ($1,$2,$3,$4)", dto.AuthorID, dto.Title, dto.Body, time.Now())
	InvalidateCache(ctx, db)
	return
}

func setCache(ctx context.Context, db Dependencies, key string, value interface{}) (err error) {
	p, err := json.Marshal(value)
	if err != nil {
		return
	}

	db.Cache.Set(ctx, key, string(p), exp)
	return
}

func getCache(ctx context.Context, db Dependencies, key string, value interface{}) (err error) {
	p, err := db.Cache.Get(ctx, key).Result()
	if err != nil {
		return
	}

	return json.Unmarshal([]byte(p), value)
}
