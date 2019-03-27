package v1

import (
	"testing"

	"github.com/romeufcrosa/best-route-finder/tests"

	. "github.com/onsi/gomega"
	domain "github.com/romeufcrosa/best-route-finder/domain/entities"
	"github.com/romeufcrosa/best-route-finder/providers"
)

func TestGetRoutesSuccess(t *testing.T) {
	RegisterTestingT(t)

	sqlPool := tests.GetPool()
	tests.SetupRoutesScenario(sqlPool)

	params := providers.NewParams(sqlPool)
	providers.Configure(params, func() bool {
		return true
	})

	interactor, err := providers.GetRoutesProvider()

	result, err := interactor.GetRoute(ctx, 1, 2)
	Expect(err).To(BeNil(), "Should not return an error")
	Expect(result).ToNot(BeNil(), "Should return something")

	expected := domain.Route{
		Voyage: []*domain.Node{
			{
				ID:   1,
				Name: "Node_A",
			},
			{
				ID:   4,
				Name: "Node_D",
			},
			{
				ID:   2,
				Name: "Node_B",
			},
		},
		Cost:     22,
		Duration: 2,
	}

	Expect(expected).To(Equal(result), "should match calculated optimal route")
}
