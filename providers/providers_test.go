package providers

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
)

type mProvider struct {
	myData string
}

func TestRegisterError(t *testing.T) {
	RegisterTestingT(t)

	register := Provider("my_provider")

	err := Register(register, func() (interface{}, error) {
		return nil, errors.New("provider error")
	})
	Expect(err).NotTo(BeNil(), "Should return an error")

	provider, err := Get(register)
	Expect(err).To(Equal(ErrProviderNotFound), "Should return an error")
	Expect(provider).To(BeNil(), "Should return the provider")
}

func TestProviderManagement(t *testing.T) {
	RegisterTestingT(t)

	register := Provider("my_provider")
	expected := mProvider{"my_data"}

	err := Register(register, func() (interface{}, error) {
		return expected, nil
	})
	Expect(err).To(BeNil(), "Should return no error")

	provider, err := Get(register)
	Expect(err).To(BeNil(), "Should return no error")
	Expect(provider).To(Equal(expected), "Should set the provider")

	err = Register(register, func() (interface{}, error) {
		return expected, nil
	})
	Expect(err).NotTo(BeNil(), "Should return an error when re-registering a provider")
}
