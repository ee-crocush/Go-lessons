package post

import uc "post-app/internal/usecase/post"

// AuthorDTO автора.
type AuthorDTO struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// PostDTO поста.
type PostDTO struct {
	ID        int32  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// MapGetByIDUseCaseToRequest преобразует данные из usecase получения поста по ID в dto.
func MapGetByIDUseCaseToRequest(in uc.GetByIDOutputDTO) GetPostByIDResponse {
	author := AuthorDTO{
		ID:   in.Author.ID,
		Name: in.Author.Name,
	}
	post := PostDTO{
		ID:        in.Post.ID,
		Title:     in.Post.Title,
		Content:   in.Post.Content,
		CreatedAt: in.Post.CreatedAt,
	}

	return GetPostByIDResponse{Author: author, Post: post}
}

// AuthorWithPostsDTO автор с постами.
type AuthorWithPostsDTO struct {
	Author AuthorDTO `json:"author"`
	Posts  []PostDTO `json:"posts"`
}

// MapGetAllUseCaseToRequest преобразует данные из usecase получения постов в dto.
func MapGetAllUseCaseToRequest(in uc.GetAllOutputDTO) GetAllPostsResponse {
	var result []AuthorWithPostsDTO
	for _, post := range in.Posts {
		author := AuthorDTO{
			ID:   post.Author.ID,
			Name: post.Author.Name,
		}

		var posts []PostDTO
		for _, p := range post.Posts {
			posts = append(
				posts, PostDTO{
					ID:        p.ID,
					Title:     p.Title,
					Content:   p.Content,
					CreatedAt: p.CreatedAt,
				},
			)
		}
		result = append(
			result, AuthorWithPostsDTO{
				Author: author,
				Posts:  posts,
			},
		)

	}

	return GetAllPostsResponse{Data: result}
}

// MapGetByAuthorIDUseCaseToRequest преобразует данные из usecase получения постов по ID автора в dto.
func MapGetByAuthorIDUseCaseToRequest(in uc.GetByAuthorIDOutputDTO) GetPostByAuthorIDResponse {
	author := AuthorDTO{
		ID:   in.Author.ID,
		Name: in.Author.Name,
	}

	var posts []PostDTO
	for _, p := range in.Posts {
		posts = append(
			posts, PostDTO{
				ID:        p.ID,
				Title:     p.Title,
				Content:   p.Content,
				CreatedAt: p.CreatedAt,
			},
		)
	}

	result := AuthorWithPostsDTO{
		Author: author,
		Posts:  posts,
	}

	return GetPostByAuthorIDResponse{Data: result}
}
