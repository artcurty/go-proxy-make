package e2e

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/wiremock/go-wiremock"
	"net/http"
	"os"
	"testing"
	"time"
)

var wiremockClient *wiremock.Client

func setupWireMockStubs(client *wiremock.Client) {
	client.StubFor(wiremock.Get(wiremock.URLPathEqualTo("/proxy-order")).
		WillReturnResponse(
			wiremock.NewResponse().
				WithStatus(http.StatusOK).
				WithBody(`{"id":"xop1","product":"ps5","quantity":"1"}`).
				WithHeader("Content-Type", "application/json")))

	client.StubFor(wiremock.Post(wiremock.URLPathEqualTo("/proxy-order")).
		WillReturnResponse(
			wiremock.NewResponse().
				WithStatus(http.StatusOK).
				WithBody(`{"id":"xop1","product":"ps5","quantity":"1"}`).
				WithHeader("Content-Type", "application/json")))

	client.StubFor(wiremock.Put(wiremock.URLPathEqualTo("/payment-order")).
		WillReturnResponse(
			wiremock.NewResponse().
				WithStatus(http.StatusOK).
				WithBody(`{"orderId":"xop1","orderStatus":"paid"}`).
				WithHeader("Content-Type", "application/json")))
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "wiremock/wiremock",
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForListeningPort("8080/tcp"),
		HostConfigModifier: func(hc *container.HostConfig) {
			hc.PortBindings = map[nat.Port][]nat.PortBinding{
				"8080/tcp": {{HostPort: "8080"}},
			}
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	defer container.Terminate(ctx)

	host, err := container.Host(ctx)
	if err != nil {
		panic(err)
	}

	wiremockClient = wiremock.NewClient("http://" + host + ":8080")
	setupWireMockStubs(wiremockClient)

	time.Sleep(2 * time.Second)

	code := m.Run()

	wiremockClient.Reset()

	os.Exit(code)
}
