package redfish

import (
	"fmt"
	"log"
	"os"
	"testing"

	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/strfmt"

	apiclientRedfish "github.com/emccode/gorackhd-redfish/client"
	"github.com/emccode/gorackhd-redfish/client/redfish_v1"
	"github.com/emccode/gorackhd-redfish/models"
)

func TestRedfishGetRolesOperation(t *testing.T) {

	// create the transport
	transport := httptransport.New("localhost:9090", "/redfish/v1", []string{"http"})

	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclientRedfish.New(transport, strfmt.Default)

	//use any function to do REST operations
	resp, err := client.RedfishV1.GetRoles(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}

func TestRedfishGetPowerStatus(t *testing.T) {

	// create the transport
	transport := httptransport.New("localhost:9090", "/redfish/v1", []string{"http"})

	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclientRedfish.New(transport, strfmt.Default)

	//use any function to do REST operations
	//resp, err := client.RedfishV1.GetPower(&redfish_v1.GetPowerParams{Identifier: "52:54:be:ef:51:1b"})
	resp, err := client.RedfishV1.GetSystem(&redfish_v1.GetSystemParams{Identifier: "57154fe9d67951e70958c213"})
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", resp.Payload)
	//fmt.Printf("%#v\n", resp.Payload.ID)
	fmt.Printf("%#v\n", resp.Payload.PowerState)
}

func TestRedfishShutdown(t *testing.T) {

	// create the transport
	transport := httptransport.New("localhost:9090", "/redfish/v1", []string{"http"})

	// configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
	if os.Getenv("GORACKHD_ENDPOINT") != "" {
		transport.Host = os.Getenv("GORACKHD_ENDPOINT")
	}

	// create the API client, with the transport
	client := apiclientRedfish.New(transport, strfmt.Default)

	action := &models.RackHDResetAction{
		ResetType: "GracefulShutdown",
	}
	resp, err := client.RedfishV1.DoReset(&redfish_v1.DoResetParams{Identifier: "57154fe9d67951e70958c213", Payload: action})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp)
}
