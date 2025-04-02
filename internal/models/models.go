package models

type Article struct {
    ID      string `json:"id"`
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
    Author  string `json:"author"`
}

type User struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}