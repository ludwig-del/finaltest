package test

import (
	"myapp/entity"
	"testing"
	. "github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
)

func Test(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Sample Test", func(t *testing.T) {
		approved := entity.Product{
			Name:        "Sample Product",
			Price:       19.99,
			Quantity:    10,
			Description: "This is a sample product.",
		}
		ok, err := govalidator.ValidateStruct(approved)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

