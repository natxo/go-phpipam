package phpipam

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type IpamSubnetRequest struct {
	Description string `json:"description"`
}

type Ipamsubnetresponse struct {
	Code    int     `json:"code"`
	Success bool    `json:"success"`
	Message string  `json:"message"`
	ID      string  `json:"id"`
	Data    string  `json:"data"`
	Time    float64 `json:"time"`
}

type IpamSubnet struct {
	ID                    string `json:"id"`
	Subnet                string `json:"subnet"`
	Mask                  string `json:"mask"`
	SectionID             string `json:"sectionId"`
	Description           string `json:"description"`
	LinkedSubnet          string `json:"linked_subnet"`
	FirewallAddressObject string `json:"firewallAddressObject"`
	VrfID                 string `json:"vrfId"`
	MasterSubnetID        string `json:"masterSubnetId"`
	AllowRequests         string `json:"allowRequests"`
	VlanID                string `json:"vlanId"`
	ShowName              string `json:"showName"`
	Device                string `json:"device"`
	Permissions           string `json:"permissions"`
	PingSubnet            string `json:"pingSubnet"`
	DiscoverSubnet        string `json:"discoverSubnet"`
	DNSrecursive          string `json:"DNSrecursive"`
	DNSrecords            string `json:"DNSrecords"`
	NameserverID          string `json:"nameserverId"`
	ScanAgent             string `json:"scanAgent"`
	IsFolder              string `json:"isFolder"`
	IsFull                string `json:"isFull"`
	Tag                   string `json:"tag"`
	Threshold             string `json:"threshold"`
	Location              string `json:"location"`
	EditDate              string `json:"editDate"`
	LastScan              string `json:"lastScan"`
	LastDiscovery         string `json:"lastDiscovery"`
	ResolveDNS            string `json:"resolveDNS"`
	CustomerID            string `json:"customer_id"`
	Nameservers           struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Namesrv1    string `json:"namesrv1"`
		Description string `json:"description"`
		Permissions string `json:"permissions"`
		EditDate    string `json:"editDate"`
	} `json:"nameservers"`
	GatewayID string `json:"gatewayId"`
	Gateway   struct {
		IPAddr string `json:"ip_addr"`
		ID     string `json:"id"`
	} `json:"gateway"`
}

type Subnet struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		ID                    string      `json:"id"`
		Subnet                string      `json:"subnet"`
		Mask                  string      `json:"mask"`
		Sectionid             string      `json:"sectionId"`
		Description           string      `json:"description"`
		LinkedSubnet          interface{} `json:"linked_subnet"`
		Firewalladdressobject interface{} `json:"firewallAddressObject"`
		Vrfid                 interface{} `json:"vrfId"`
		Mastersubnetid        string      `json:"masterSubnetId"`
		Allowrequests         string      `json:"allowRequests"`
		Vlanid                interface{} `json:"vlanId"`
		Showname              string      `json:"showName"`
		Device                interface{} `json:"device"`
		Permissions           string      `json:"permissions"`
		Pingsubnet            string      `json:"pingSubnet"`
		Discoversubnet        string      `json:"discoverSubnet"`
		Resolvedns            string      `json:"resolveDNS"`
		Dnsrecursive          string      `json:"DNSrecursive"`
		Dnsrecords            string      `json:"DNSrecords"`
		Nameserverid          string      `json:"nameserverId"`
		Scanagent             string      `json:"scanAgent"`
		CustomerID            interface{} `json:"customer_id"`
		Isfolder              string      `json:"isFolder"`
		Isfull                string      `json:"isFull"`
		Ispool                string      `json:"isPool"`
		Tag                   string      `json:"tag"`
		Threshold             string      `json:"threshold"`
		Location              interface{} `json:"location"`
		Editdate              string      `json:"editDate"`
		Lastscan              interface{} `json:"lastScan"`
		Lastdiscovery         interface{} `json:"lastDiscovery"`
		Calculation           struct {
			Type           string `json:"Type"`
			IPAddress      string `json:"IP address"`
			Network        string `json:"Network"`
			Broadcast      string `json:"Broadcast"`
			SubnetBitmask  string `json:"Subnet bitmask"`
			SubnetNetmask  string `json:"Subnet netmask"`
			SubnetWildcard string `json:"Subnet wildcard"`
			MinHostIP      string `json:"Min host IP"`
			MaxHostIP      string `json:"Max host IP"`
			NumberOfHosts  string `json:"Number of hosts"`
			SubnetClass    bool   `json:"Subnet Class"`
		} `json:"calculation"`
	} `json:"data"`
	Time float64 `json:"time"`
}

type Simplesubnet struct {
	ID                    string `json:"id"`
	Subnet                string `json:"subnet"`
	Mask                  string `json:"mask"`
	SectionID             string `json:"sectionId"`
	Description           string `json:"description"`
	LinkedSubnet          string `json:"linked_subnet"`
	FirewallAddressObject string `json:"firewallAddressObject"`
	VrfID                 string `json:"vrfId"`
	MasterSubnetID        string `json:"masterSubnetId"`
	AllowRequests         string `json:"allowRequests"`
	VlanID                string `json:"vlanId"`
	ShowName              string `json:"showName"`
	Device                string `json:"device"`
	Permissions           string `json:"permissions"`
	PingSubnet            string `json:"pingSubnet"`
	DiscoverSubnet        string `json:"discoverSubnet"`
	DNSrecursive          string `json:"DNSrecursive"`
	DNSrecords            string `json:"DNSrecords"`
	NameserverID          string `json:"nameserverId"`
	ScanAgent             string `json:"scanAgent"`
	IsFolder              string `json:"isFolder"`
	IsFull                string `json:"isFull"`
	Tag                   string `json:"tag"`
	Threshold             string `json:"threshold"`
	Location              string `json:"location"`
	EditDate              string `json:"editDate"`
	LastScan              string `json:"lastScan"`
	LastDiscovery         string `json:"lastDiscovery"`
	ResolveDNS            string `json:"resolveDNS"`
	CustomerID            string `json:"customer_id"`
	Nameservers           struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Namesrv1    string `json:"namesrv1"`
		Description string `json:"description"`
		Permissions string `json:"permissions"`
		EditDate    string `json:"editDate"`
	} `json:"nameservers"`
	GatewayID string `json:"gatewayId"`
	Gateway   struct {
		IPAddr string `json:"ip_addr"`
		ID     string `json:"id"`
	} `json:"gateway"`
}
type Subnets struct {
	Code    int            `json:"code"`
	Success bool           `json:"success"`
	Data    []Simplesubnet `json:"data"`
	Time    float64        `json:"time"`
}

type Searchresult struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    []struct {
		ID          string `json:"id"`
		Subnet      string `json:"subnet"`
		Mask        string `json:"mask"`
		SectionID   string `json:"sectionId"`
		Description string `json:"description"`
	} `json:"data"`
	Time    float64 `json:"time"`
	Message string  `json:"message"`
}

// GetSubnet Client pointer method to get all phpipam subnet data using subnetID
// string, returns Subnet struct and error
func (c *Client) GetSubnet(subnetID string) (Subnet, error) {
	var subnetData Subnet
	req, _ := http.NewRequest("GET", c.ServerURL+"/api/"+c.Application+"/subnets/"+subnetID+"/", nil)
	body, err := c.Do(req)
	if err != nil {
		return subnetData, err
	}
	err = json.Unmarshal([]byte(body), &subnetData)
	if err != nil {
		return subnetData, err
	}
	if subnetData.Code != 200 {
		code := string(subnetData.Code)
		return subnetData, errors.New(code)
	}
	return subnetData, nil
}

// Searchsubnet client pointer method to get subnet data searching using cidr
// subnet, returns Searchresult and error
func (c *Client) Searchsubnet(parentsubnet string) (Searchresult, error) {
	var searchdata Searchresult
	req, err := http.NewRequest("GET", c.ServerURL+"/api/"+c.Application+"/subnets/cidr/"+parentsubnet, nil)
	if err != nil {
		return searchdata, err
	}
	body, err := c.Do(req)
	if err != nil {
		return searchdata, err
	}
	err = json.Unmarshal([]byte(body), &searchdata)
	if err != nil {
		return searchdata, err
	}
	if searchdata.Code != 200 {
		return searchdata, errors.New(searchdata.Message)
	}
	return searchdata, nil
}

// get all children subnets of a specific parent
func (c *Client) Getallchildrensubnets(subnetid string) ([]Simplesubnet, error) {
	var subnetsdata Subnets
	req, err := http.NewRequest("GET", c.ServerURL+"/api/"+c.Application+"/subnets/"+subnetid+"/slaves/", nil)
	if err != nil {
		return nil, err
	}
	body, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(body), &subnetsdata)
	if err != nil {
		return nil, err
	}
	if subnetsdata.Code != 200 {
		return nil, errors.New("status code not 200 while getting all children subnets")
	}
	return subnetsdata.Data, nil
}

func (c *Client) Createsubnet(rootsubnetid, mask, name string) (subnetdata string, subnetid string, err error) {
	var subnet IpamSubnetRequest
	var subnetresp Ipamsubnetresponse
	subnet.Description = name
	data, err := json.Marshal(subnet)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest("POST", c.ServerURL+"/api/"+c.Application+"/subnets/"+rootsubnetid+"/first_subnet/"+mask+"/", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}
	body, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(body), &subnetresp)
	if err != nil {
		return subnetresp.Data, subnetresp.ID, err
	}
	if subnetresp.Success != true {
		return subnetresp.Data, subnetresp.ID, errors.New(subnetresp.Message)
	}
	fmt.Println("subnet created:", subnetresp.Data)
	return subnetresp.Data, subnetresp.ID, nil

}
