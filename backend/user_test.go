package backend

import (
	"testing"
	//"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

type User struct {
	Name     string `valid:"required~Name ห้ามว่าง"`
	Email    string `valid:"email,required~Email ห้ามว่าง"`
	Age      int    `valid:"range(10|18)~Age ผิดพลาด"`
	Phone    string `valid:"required~Phone ห้ามว่าง,matches(^[0]\\d{9}$)~Phone ห้ามเกิน 10 ตัว"`
	Password string `valid:"matches(^[0-9]\\d{9}$)~Password ต้องมีขนาด 10 ตัว"`
}

// Name ว่าง
func TestName(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ab := User{
		Name:     "",
		Email:    "james@gmail.com",
		Age:      11,
		Phone:    "0888888888",
		Password: "1234567890",
	}

	ok, err := govalidator.ValidateStruct(ab)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Name ห้ามว่าง"))
}

// email ผิด
func TestEmail(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ab := User{
		Name:     "Natthapong",
		Email:    "james",
		Age:      18,
		Phone:    "0888888888",
		Password: "1234567890",
	}

	ok, err := govalidator.ValidateStruct(ab)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Email: james does not validate as email"))
}

// email ห้ามว่าง
func TestEmailtwo(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ab := User{
		Name:     "Natthapong",
		Email:    "",
		Age:      18,
		Phone:    "0888888888",
		Password: "1234567890",
	}

	ok, err := govalidator.ValidateStruct(ab)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Email ห้ามว่าง"))
}

// Age เกิน
func TestAge(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ab := User{
		Name:     "Natthapong",
		Email:    "james@gmail.com",
		Age:      19,
		Phone:    "0888888888",
		Password: "1234567890",
	}

	ok, err := govalidator.ValidateStruct(ab)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Age ผิดพลาด"))
}

// Phone ห้ามว่าง
func TestPhone(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ab := User{
		Name:     "Natthapong",
		Email:    "james@gmail.com",
		Age:      18,
		Phone:    "",
		Password: "1234567890",
	}

	ok, err := govalidator.ValidateStruct(ab)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Phone ห้ามว่าง"))
}

// Phone ห้าม < 10ตัว และ ขึ้นต้นด้วย 0
func TestPhoneTwo(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ab := User{
		Name:     "Natthapong",
		Email:    "james@gmail.com",
		Age:      18,
		Phone:    "18888888881",
		Password: "1234567890",
	}

	ok, err := govalidator.ValidateStruct(ab)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Phone ห้ามเกิน 10 ตัว"))
}

// Password ต้องมีขนาด 10 ตัว
func TestPass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ab := User{
		Name:     "Natthapong",
		Email:    "james@gmail.com",
		Age:      18,
		Phone:    "0888888888",
		Password: "12345",
	}

	ok, err := govalidator.ValidateStruct(ab)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Password ต้องมีขนาด 10 ตัว"))
}

// ใส่ทุกค่า
func TestAll(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ab := User{
		Name:     "Natthapong",
		Email:    "james@gmail.com",
		Age:      18,
		Phone:    "0888888888",
		Password: "1234567890",
	}

	ok, err := govalidator.ValidateStruct(ab)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())
}