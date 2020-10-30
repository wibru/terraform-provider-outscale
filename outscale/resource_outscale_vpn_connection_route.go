package outscale

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oscgo "github.com/outscale/osc-sdk-go/v2"
)

func resourceOutscaleVPNConnectionRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceOutscaleVPNConnectionRouteCreate,
		Read:   resourceOutscaleVPNConnectionRouteRead,
		Delete: resourceOutscaleVPNConnectionRouteDelete,
		Importer: &schema.ResourceImporter{
			State: resourceOutscaleOAPIVPNConnectionRouteImportState,
		},

		Schema: map[string]*schema.Schema{
			"destination_ip_range": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpn_connection_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceOutscaleVPNConnectionRouteCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	destinationIPRange := d.Get("destination_ip_range").(string)
	vpnConnectionID := d.Get("vpn_connection_id").(string)

	req := oscgo.CreateVpnConnectionRouteRequest{
		DestinationIpRange: destinationIPRange,
		VpnConnectionId:    vpnConnectionID,
	}

	_, _, err := conn.VpnConnectionApi.CreateVpnConnectionRoute(context.Background()).CreateVpnConnectionRouteRequest(req).Execute()
	if err != nil {
		return fmt.Errorf("Error creating Outscale VPN Conecction Route: %s", err)
	}

	d.SetId(fmt.Sprintf("%s:%s", destinationIPRange, vpnConnectionID))

	return resourceOutscaleVPNConnectionRouteRead(d, meta)
}

func resourceOutscaleVPNConnectionRouteRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	destinationIPRange, vpnConnectionID := resourceOutscaleVPNConnectionRouteParseID(d.Id())

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"pending"},
		Target:     []string{"available", "failed"},
		Refresh:    vpnConnectionRouteRefreshFunc(conn, &destinationIPRange, &vpnConnectionID),
		Timeout:    10 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	r, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for Outscale VPN Connection Route(%s) to become ready: %s", d.Id(), err)
	}

	resp := r.(oscgo.ReadVpnConnectionsResponse)

	if err := d.Set("request_id", resp.ResponseContext.GetRequestId()); err != nil {
		return err
	}

	return nil
}

func resourceOutscaleVPNConnectionRouteDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	destinationIPRange, vpnConnectionID := resourceOutscaleVPNConnectionRouteParseID(d.Id())

	req := oscgo.DeleteVpnConnectionRouteRequest{
		DestinationIpRange: destinationIPRange,
		VpnConnectionId:    vpnConnectionID,
	}

	_, _, err := conn.VpnConnectionApi.DeleteVpnConnectionRoute(context.Background()).DeleteVpnConnectionRouteRequest(req).Execute()
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"deleting"},
		Target:     []string{"deleted", "failed"},
		Refresh:    vpnConnectionRouteRefreshFunc(conn, &destinationIPRange, &vpnConnectionID),
		Timeout:    10 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for Outscale VPN Connection Route(%s) to become deleted: %s", vpnConnectionID, err)
	}

	return nil
}

func vpnConnectionRouteRefreshFunc(conn *oscgo.APIClient, destinationIPRange, vpnConnectionID *string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {

		filter := oscgo.ReadVpnConnectionsRequest{
			Filters: &oscgo.FiltersVpnConnection{
				RouteDestinationIpRanges: &[]string{*destinationIPRange},
				VpnConnectionIds:         &[]string{*vpnConnectionID},
			},
		}

		resp, _, err := conn.VpnConnectionApi.ReadVpnConnections(context.Background()).ReadVpnConnectionsRequest(filter).Execute()
		if err != nil {
			switch {
			case strings.Contains(fmt.Sprint(err), "RequestLimitExceeded:"):
				return nil, "pending", nil
			case strings.Contains(fmt.Sprint(err), "404"):
				return nil, "deleted", nil
			default:
				return nil, "failed", fmt.Errorf("Error on vpnConnectionRouteRefresh: %s", err)
			}
		}

		if len(resp.GetVpnConnections()) == 0 {
			return nil, "failed", fmt.Errorf("error on vpnConnectionRouteRefresh: there are not vpn connections with id %v", vpnConnectionID)
		}
		vpnConnection := resp.GetVpnConnections()[0]

		routes, ok := vpnConnection.GetRoutesOk()
		if ok {
			for _, route := range *routes {
				if route.GetDestinationIpRange() == *destinationIPRange {
					return resp, route.GetState(), nil
				}
			}
		}

		return resp, "pending", nil
	}
}

func resourceOutscaleVPNConnectionRouteParseID(ID string) (string, string) {
	parts := strings.SplitN(ID, ":", 2)
	return parts[0], parts[1]
}

func resourceOutscaleOAPIVPNConnectionRouteImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	conn := meta.(*OutscaleClient).OSCAPI

	parts := strings.SplitN(d.Id(), "_", 2)
	if len(parts) != 2 {
		return nil, errors.New("import format error: to import a Outscale VPN connection Route, use the format {vpn_connection_id}_{destination_ip_range}")
	}

	vpnConnectionID := parts[0]
	destinationIPRange := parts[1]

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"pending"},
		Target:     []string{"available", "failed"},
		Refresh:    vpnConnectionRouteRefreshFunc(conn, &destinationIPRange, &vpnConnectionID),
		Timeout:    10 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err := stateConf.WaitForState()
	if err != nil {
		return nil, fmt.Errorf("Error waiting for Outscale VPN Connection Route import(%s) to become ready: %s", d.Id(), err)
	}

	if err != nil {
		if strings.Contains(fmt.Sprint(err), "NotFound") {
			log.Printf("[WARN] VPN Connection route %q could not be found. Removing Route from state.", vpnConnectionID)
			return nil, err
		}
		return nil, err
	}

	if err := d.Set("vpn_connection_id", vpnConnectionID); err != nil {
		return nil, fmt.Errorf("error setting `%s` for Outscale VPN Connection Route(%s): %s", "vpn_connection_id", vpnConnectionID, err)
	}
	if err := d.Set("destination_ip_range", destinationIPRange); err != nil {
		return nil, fmt.Errorf("error setting `%s` for Outscale VPN Connection Route(%s): %s", "destination_ip_range", destinationIPRange, err)
	}

	d.SetId(fmt.Sprintf("%s:%s", destinationIPRange, vpnConnectionID))

	return []*schema.ResourceData{d}, nil
}
