package post

import (
	authordom "post-app/internal/domain/author"
	dom "post-app/internal/domain/post"
)

// AuthorDTO автор поста.
type AuthorDTO struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// MapAuthorToDTO преобразует автора в DTO.
func MapAuthorToDTO(a *authordom.Author) AuthorDTO {
	return AuthorDTO{
		ID:   a.ID().Value(),
		Name: a.Name().Value(),
	}
}

// PostDTO пост.
type PostDTO struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// MapPostToDTO преобразует пост в DTO.
func MapPostToDTO(p *dom.Post) PostDTO {
	return PostDTO{
		ID:        p.ID().Value(),
		Title:     p.Title().Value(),
		Content:   p.Content().Value(),
		CreatedAt: p.CreatedAt().String(),
	}
}

// AuthorWithPostsDTO автор с постами.
type AuthorWithPostsDTO struct {
	Author AuthorDTO `json:"author"`
	Posts  []PostDTO `json:"posts"`
}

// MapAuthorWithPostsToDTO преобразует автора и его посты в DTO.
func MapAuthorWithPostsToDTO(a *authordom.Author) AuthorWithPostsDTO {
	var postsDto []PostDTO
	for _, p := range a.Posts() {
		postsDto = append(postsDto, MapPostToDTO(p))
	}

	return AuthorWithPostsDTO{MapAuthorToDTO(a), postsDto}
}
