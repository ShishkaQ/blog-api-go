package main

import (
    "blog-api/internal/handlers"
    "blog-api/internal/middleware"
    "blog-api/internal/storage"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    store := storage.New()

    authHandler := handlers.NewAuthHandler(store)
    articleHandler := handlers.NewArticleHandler(store)

    // Public routes
    public := r.Group("/api")
    {
        public.POST("/login", authHandler.Login)
        public.GET("/articles/:id", articleHandler.GetArticle)
    }

    // Protected routes
    protected := r.Group("/api")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.POST("/articles", articleHandler.CreateArticle)
    }

    r.Run(":8080")
}