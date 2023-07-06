package news

type Service struct {
}

func NewNewsService() *Service {
	return &Service{}
}

type News struct {
	Data NewsData
}
type NewsData struct {
	Items []NewsItems
}

type NewsItems struct {
	Title string
	Text  string
	Date  string
}
