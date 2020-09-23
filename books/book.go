package books

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (p Book) CategoryByLength() string {
	if p.Pages >= 300 {
		return "NOVEL"
	}
	return "SHORT STORY"
}

func NewBookFromJSON(s string) Book {
	return Book{}
}
