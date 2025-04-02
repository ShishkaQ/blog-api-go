package storage

import "blog-api/internal/models"

type MemoryStorage struct {
    articles  map[string]models.Article
    users     map[string]models.User
}

func New() *MemoryStorage {
    return &MemoryStorage{
        articles: make(map[string]models.Article),
        users: map[string]models.User{
            "admin": {Username: "admin", Password: "admin123"},
        },
    }
}

// Добавляем методы для работы со статьями
func (s *MemoryStorage) CreateArticle(article models.Article) {
    s.articles[article.ID] = article
}

func (s *MemoryStorage) GetArticle(id string) (models.Article, bool) {
    article, exists := s.articles[id]
    return article, exists
}

// Методы для работы с пользователями
func (s *MemoryStorage) ValidateUser(username, password string) bool {
    user, exists := s.users[username]
    return exists && user.Password == password
}

func (s *MemoryStorage) GetUser(username string) (models.User, bool) {
    user, exists := s.users[username]
    return user, exists
}