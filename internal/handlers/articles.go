package handlers

import (
    "blog-api/internal/models"
    "blog-api/internal/storage"
    "github.com/gin-gonic/gin"
    "net/http"
)

type ArticleHandler struct {
    store *storage.MemoryStorage
}

func NewArticleHandler(s *storage.MemoryStorage) *ArticleHandler {
    return &ArticleHandler{store: s}
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
    var article models.Article
    if err := c.ShouldBindJSON(&article); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    article.Author = c.GetString("username")
    h.store.CreateArticle(article)
    c.JSON(http.StatusCreated, article)
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
    id := c.Param("id")
    article, exists := h.store.GetArticle(id)
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
        return
    }
    c.JSON(http.StatusOK, article)
}