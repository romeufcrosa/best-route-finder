package v1

import (
	"testing"

	"github.com/romeufcrosa/best-route-finder/tests"

	. "github.com/onsi/gomega"
	"github.com/romeufcrosa/best-route-finder/configurations"
	"github.com/romeufcrosa/best-route-finder/domain/entities"
	"github.com/romeufcrosa/best-route-finder/providers"
	"github.com/romeufcrosa/best-route-finder/tests/fixtures"
)

func TestAddEdgeSucess(t *testing.T) {
	RegisterTestingT(t)

	sqlPool, err := configurations.Pool()
	if err != nil {
		t.Fatal(err.Error())
	}
	tests.ClearTables(sqlPool)

	params := providers.NewParams(sqlPool)
	providers.Configure(params, func() bool {
		return true
	})

	edge := entities.Edge{}
	err = fixtures.LoadJSONInto("testdata/an_edge.json", &edge)
	if err != nil {
		t.Fatal(err.Error())
	}

	interactor, err := providers.GetRoutesProvider()

	result, err := interactor.AddEdge(ctx, edge)
	Expect(err).To(BeNil(), "Should not return an error")
	Expect(result).ToNot(BeNil(), "Should return something")
}
