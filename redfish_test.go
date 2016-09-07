package redfish

import (
	"fmt"
	"log"
	"os"
	"testing"

	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	apiclientRedfish "github.com/codedellemc/gorackhd-redfish/client"
	"github.com/codedellemc/gorackhd-redfish/client/redfish_v1"
	"github.com/codedellemc/gorackhd-redfish/models"
)

func TestRedfishGetAccountsOperation(t *testing.T) {

	// create the transport
	transport := rc.New("localhost:9090", "/redfish/v1", []string{"http"})

	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclientRedfish.New(transport, strfmt.Default)

	//use any function to do REST operations
	resp, err := client.RedfishV1.GetAccounts(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

func TestRedfishGetPowerStatus(t *testing.T) {

	// create the transport
	transport := rc.New("localhost:9090", "/redfish/v1", []string{"http"})

	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclientRedfish.New(transport, strfmt.Default)

	params := redfish_v1.NewGetSystemParams()
	params = params.WithIdentifier("57154fe9d67951e70958c213")
	resp, err := client.RedfishV1.GetSystem(params)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", resp.Payload)
	//fmt.Printf("%#v\n", resp.Payload.ID)
	fmt.Printf("%#v\n", resp.Payload.PowerState)
}

func TestRedfishShutdown(t *testing.T) {

	// create the transport
	transport := rc.New("localhost:9090", "/redfish/v1", []string{"http"})

	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclientRedfish.New(transport, strfmt.Default)

	resetType := "GracefulShutdown"
	action := &models.RackHDResetActionResetAction{
		ResetType: &resetType,
	}
	params := redfish_v1.NewDoResetParams()
	params = params.WithIdentifier("57154fe9d67951e70958c213")
	params = params.WithPayload(action)
	resp, err := client.RedfishV1.DoReset(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp)
}
