package v1

import (
	"context"
	"net/http"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/romeufcrosa/best-route-finder/configurations"
	"github.com/romeufcrosa/best-route-finder/configurations/readers"
	"github.com/romeufcrosa/best-route-finder/domain/entities"
	"github.com/romeufcrosa/best-route-finder/providers"
	"github.com/romeufcrosa/best-route-finder/tests"
	"github.com/romeufcrosa/best-route-finder/tests/fixtures"
)

var (
	ctx      = context.TODO()
	httpPool = &http.Client{
		Transport: http.DefaultTransport,
	}
	reader readers.FileReader
)

func TestAddNodeSucess(t *testing.T) {
	RegisterTestingT(t)

	sqlPool, err := configurations.Pool()
	if err != nil {
		t.Fatal(err.Error())
	}
	params := providers.NewParams(sqlPool)
	providers.Configure(params, func() bool {
		return true
	})

	node := entities.Node{}
	err = fixtures.LoadJSONInto("testdata/a_node.json", &node)
	if err != nil {
		t.Fatal(err.Error())
	}

	interactor, err := providers.GetRoutesProvider()

	result, err := interactor.AddNode(ctx, node)
	Expect(err).To(BeNil(), "Should not return an error")
	Expect(result).ToNot(BeNil(), "Should return something")
}

func init() {
	reader = tests.ConfsReader()
	configurations.Load(ctx, reader)

	providers.RegisterRepositoryProviders()
}
