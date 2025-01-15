package unit

import(
	"fmt"
	"github.com/Rtppts/testback/entity"
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestAllTrue(t *testing.T){

	g := NewGomegaWithT(t)

	t.Run(`TestAllTrue`, func(t *testing.T){
		Ratta := entity.Ratta{
			Name: "Ratta",
			ID: "B6512651",
			Phone: "0644044078",
			Email: "0837335326.3d.v2@gamil.com",
		}

		ok, err := govalidator.ValidateStruct(Ratta)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
		// g.Expect(err.Error()).To(Equal(""))
	})
}










func TestNameNull (t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`TestNameNull`, func(t *testing.T){
		Ratta := entity.Ratta{
			Name: "",
			ID: "B6512651",
			Phone: "0644044078",
			Email: "0837335326.3d.v2@gamil.com",
		}

		ok, err := govalidator.ValidateStruct(Ratta)
		fmt.Print("err ", err)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name is required"))
	})
}










func TestEmailIsInvalid(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`EmailInvalid`, func(t *testing.T){
		Ratta := entity.Ratta{
			Name: "Rattapon",
			ID: "B6512651",
			Phone: "0644044078",
			Email: "555555",
		}

		ok, err := govalidator.ValidateStruct(Ratta)
		fmt.Print("err ", err)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Email %s invalid", "is")))
	})
}



