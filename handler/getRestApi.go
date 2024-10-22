package handler

var Books = map[string][]string{
	"PARKJUNGMIN":     {"요즘 사는 맛", "쓸만한 인간", "쓰고 싶다 쓰고 싶지 않다"},
	"SEEMEEZ HARUKEE": {"작별의 건너편1", "작별의 건너편2", "작별의 건너편3"},
	"KIMHOYUN":        {"불편한 편의점1", "불편한 편의점2"},
	"BRANDEN BURZURD": {"골든 티켓"},
}

func GetBooks(author string) []string {
	books := Books[author]

	return books
}
