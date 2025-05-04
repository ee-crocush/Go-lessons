package post

import "time"

// PostID идентификатор поста.
type PostID struct {
	value int32
}

// NewPostID создает новый идентификатор поста.
func NewPostID(id int32) (PostID, error) {
	if id < 1 {
		return PostID{}, ErrInvalidPostID
	}
	return PostID{value: id}, nil
}

// Value возвращает значение идентификатора.
func (t PostID) Value() int32 { return t.value }

// Equal сравнивает два идентификатора.
func (t PostID) Equal(other PostID) bool { return t.value == other.value }

// PostTitle титул поста.
type PostTitle struct {
	value string
}

// NewPostTitle создает новый титул поста.
func NewPostTitle(text string) (PostTitle, error) {
	if len(text) > 0 {
		return PostTitle{text}, nil
	}

	return PostTitle{}, ErrEmptyPostTitle
}

// Value возвращает значение титула.
func (t PostTitle) Value() string { return t.value }

// PostContent содержание поста.
type PostContent struct {
	value string
}

// NewPostTitle создает содержание поста.
func NewPostContent(text string) (PostContent, error) {
	if len(text) > 0 {
		return PostContent{text}, nil
	}

	return PostContent{}, ErrEmptyPostContent
}

// Value возвращает значение содержания.
func (c PostContent) Value() string { return c.value }

// Timestamp временная метка поста.
type Timestamp struct {
	value time.Time
}

// NewTimestamp создает новую временную метку.
func NewTimestamp() Timestamp {
	return Timestamp{time.Now().UTC()}
}

// FromUnixSeconds создаёт Timestamp из секунд
func FromUnixSeconds(s int64) Timestamp {
	return Timestamp{value: time.Unix(s, 0)}
}

// Time возвращает значение временной метки.
func (t Timestamp) Time() time.Time {
	return t.value
}

// String возвращает строковое значение в формате 2006-01-02 15:04:05
func (t Timestamp) String() string {
	return t.value.Format(time.DateTime)
}
