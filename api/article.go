package api

type Article struct{}

func (article *Article) FindAllArticle() {}
func (article *Article) FindArticle()    {}

func NewArticle() *Article {
	return &Article{}
}
