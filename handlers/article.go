package handlers

import (
	"net/http"
	"strconv"
	"time"

	"net/url"

	"github.com/gofiber/fiber/v2"
)

type Article struct {
	ID        uint      `json:"id"`
	AuthorID  uint      `json:"author_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleDTO struct {
	AuthorID string `json:"author_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

type ArticleParams struct {
	Limit  int
	Page   int
	Sorted bool
	Author string
}

func (a ArticleParams) CacheTag() string {
	return (strconv.Itoa(a.Limit) + "&" + strconv.Itoa(a.Page) + "&" + url.QueryEscape(a.Author))
}

func GetAllArticles(c *fiber.Ctx) (err error) {
	params := ArticleParams{}

	dep, ok := c.Locals("deps").(*Dependencies)
	if !ok {
		return c.SendStatus(http.StatusInternalServerError)
	}

	params.Limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		params.Limit = 25
	}
	params.Page, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		params.Page = 1
	}

	if c.Query("sorted", "false") == "true" {
		params.Sorted = true
	}

	params.Author = c.Query("author", "")
	result := FetchArticlesFromCache(c.Context(), *dep, params.CacheTag())

	if len(result) == 0 {
		result = FetchArticlesFromDB(c.Context(), *dep, params)
	}

	return c.JSON(result)
}

func GetArticle(c *fiber.Ctx) (err error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return
	}
	if id < 0 {
		id -= (2 * id)
	}

	dep, ok := c.Locals("deps").(*Dependencies)
	if !ok {
		return c.SendStatus(http.StatusInternalServerError)
	}

	result, err := FetchArticle(c.Context(), *dep, uint(id))
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func CreateArticle(c *fiber.Ctx) error {
	newArticle := ArticleDTO{}

	dep, ok := c.Locals("deps").(*Dependencies)
	if !ok {
		return c.SendStatus(http.StatusInternalServerError)
	}

	if c.BodyParser(&newArticle) != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	return PersistArticle(c.Context(), *dep, newArticle)
}
