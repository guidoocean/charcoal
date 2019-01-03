package content_test

import (
	. "github.com/egoholic/charcoal/entity/content"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	title1 = "Title1"
	body1  = "Body 1"
)

var _ = Describe("Content Entity", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("return pointer on Content", func() {
				Expect(New(title1, body1)).To(BeAssignableToTypeOf(&Content{}))
			})
		})
	})
	Context("accessors", func() {
		Describe("Title()", func() {
			It("returns title", func() {
				content := New(title1, body1)
				Expect(content.Title()).To(Equal(title1))
			})
		})
		Describe("AssignTitle()", func() {
			Context("when persisted", func() {
				It("returns title", func() {
					content := New(title1, body1)
					Expect(content.Title()).To(Equal(title1))
					newTitle := "New Title"
					content.AssignTitle(newTitle)
					Expect(content.Title()).To(Equal(newTitle))
				})
			})
			Context("when not persisted", func() {
				It("returns title", func() {
					content := New(title1, body1)
					Expect(content.Title()).To(Equal(title1))
				})
			})
		})
		Describe("Body()", func() {
			It("returns body", func() {
				content := New(title1, body1)
				Expect(content.Body()).To(Equal(body1))
			})
		})
		Describe("AssignBody()", func() {
			Context("when persisted", func() {
				It("returns body", func() {
					content := New(title1, body1)
					Expect(content.Body()).To(Equal(body1))
				})
			})
			Context("when not persisted", func() {
				It("returns body", func() {
					content := New(title1, body1)
					Expect(content.Body()).To(Equal(body1))
				})
			})
		})
	})
})
