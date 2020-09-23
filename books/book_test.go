package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/chengyayu/ginkgodemo/books"
)

var _ = Describe("Book", func() {

	It("can be loaded from JSON", func() {
		book := books.NewBookFromJSON(`{
				"title":"Les Miserables",
				"author":"Victor Hugo",
				"pages":1488
			}`)

		Expect(book.Title).To(Equal("Les Miserables"))
		Expect(book.Author).To(Equal("Victor Hugo"))
		Expect(book.Pages).To(Equal(1488))
	})

	var (
		longBook  books.Book
		shortBook books.Book
	)

	BeforeEach(func() {
		longBook = books.Book{
			Title:  "The Return Of The King",
			Author: "Ergo Wang",
			Pages:  1200,
		}

		shortBook = books.Book{
			Title:  "JJ World",
			Author: "Alen",
			Pages:  98,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})
})
