package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/romeufcrosa/best-route-finder/configurations"
	"github.com/romeufcrosa/best-route-finder/configurations/readers"
	domain "github.com/romeufcrosa/best-route-finder/domain/entities"
	"github.com/romeufcrosa/best-route-finder/providers"
	"github.com/romeufcrosa/best-route-finder/tests"
	"github.com/romeufcrosa/best-route-finder/tests/fixtures"
)

var (
	ctx    = context.TODO()
	reader readers.FileReader
)

func TestEdgesSuccess(t *testing.T) {
	RegisterTestingT(t)

	payload, err := fixtures.LoadJSON("testdata/edges_post_req.json")
	if err != nil {
		t.Fatal(err.Error())
	}

	var (
		returned, expected domain.Edge
	)
	response, err := fixtures.LoadJSON("testdata/edges_post_res.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	json.Unmarshal(response, &expected)

	req, _ := http.NewRequest("POST", "/api/v1/edges", bytes.NewBuffer(payload))

	httptestRecorder := httptest.NewRecorder()
	server := Router()
	server.ServeHTTP(httptestRecorder, req)

	json.Unmarshal([]byte(httptestRecorder.Body.String()), &returned)
	Expect(httptestRecorder.Code).To(Equal(http.StatusOK), "Should return a 200 OK")
	Expect(expected.From).To(Equal(returned.From), "Should return the inserted edge")
}

func TestNodesSuccess(t *testing.T) {
	RegisterTestingT(t)

	var (
		returned, expected domain.Node
	)

	payload, err := fixtures.LoadJSON("testdata/nodes_post_req.json")
	if err != nil {
		t.Fatal(err.Error())
	}

	response, err := fixtures.LoadJSON("testdata/nodes_post_res.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	json.Unmarshal(response, &expected)

	req, _ := http.NewRequest("POST", "/api/v1/nodes", bytes.NewBuffer(payload))

	httptestRecorder := httptest.NewRecorder()
	server := Router()
	server.ServeHTTP(httptestRecorder, req)

	json.Unmarshal([]byte(httptestRecorder.Body.String()), &returned)
	Expect(httptestRecorder.Code).To(Equal(http.StatusOK), "Should return a 200 OK")
	Expect(expected.Name).To(Equal(returned.Name), "Should return the inserted node")
}

func TestRoutesSuccess(t *testing.T) {
	RegisterTestingT(t)

	var (
		srcNode            = 1
		destNode           = 2
		expected, returned domain.Route
	)

	expected = domain.Route{
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

	path := fmt.Sprintf("/api/v1/routes/from/%d/to/%d", srcNode, destNode)

	req, _ := http.NewRequest("GET", path, bytes.NewBuffer([]byte{}))

	httptestRecorder := httptest.NewRecorder()
	server := Router()
	server.ServeHTTP(httptestRecorder, req)

	json.Unmarshal([]byte(httptestRecorder.Body.String()), &returned)
	Expect(httptestRecorder.Code).To(Equal(http.StatusOK), "Should return a 200 OK")
	Expect(expected).To(Equal(returned), "Should return the inserted node")
}

func init() {
	reader = tests.ConfsReader()
	configurations.Load(ctx, reader)

	sqlPool, _ := configurations.Pool()
	params := providers.NewParams(sqlPool)

	providers.Configure(params, func() bool {
		return true
	})

	providers.RegisterRepositoryProviders()
}
