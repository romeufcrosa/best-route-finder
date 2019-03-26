package internal

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
)

func TestRegisterAndLoad(t *testing.T) {
	RegisterTestingT(t)

	expected := mProvider{"my_data"}
	registration := func() (interface{}, error) {
		return expected, nil
	}
	factory := NewProviderFactory(registration)
	err := factory.Register()
	Expect(err).To(BeNil(), "Should return no error on register")

	result, err := factory.Load()
	Expect(err).To(BeNil(), "Should return no error on load")
	Expect(result).To(Equal(expected), "Should return the corresponding provider")
}

func TestRegisterAndLoadError(t *testing.T) {
	RegisterTestingT(t)

	registration := func() (interface{}, error) {
		return nil, errors.New("I am Error")
	}
	factory := NewProviderFactory(registration)
	err := factory.Register()
	Expect(err).ToNot(BeNil(), "Should return an error on register")

	result, err := factory.Load()
	Expect(err).ToNot(BeNil(), "Should return an error on load")
	Expect(result).To(BeNil(), "Should not return any provider")
}

type mProvider struct {
	myData string
}
