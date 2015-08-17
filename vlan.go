package main

type Vlan struct {
                ID            int `json:"id"`
                PrimarySubnet struct {
                                AddressSpace         string `json:"addressSpace"`
                                BroadcastAddress     string `json:"broadcastAddress"`
                                Cidr                 int    `json:"cidr"`
                                Gateway              string `json:"gateway"`
                                ID                   int    `json:"id"`
                                IsCustomerOwned      bool   `json:"isCustomerOwned"`
                                IsCustomerRoutable   bool   `json:"isCustomerRoutable"`
                                ModifyDate           string `json:"modifyDate"`
                                Netmask              string `json:"netmask"`
                                NetworkIdentifier    string `json:"networkIdentifier"`
                                NetworkVlanID        int    `json:"networkVlanId"`
                                SortOrder            string `json:"sortOrder"`
                                SubnetType           string `json:"subnetType"`
                                TotalIPAddresses     int    `json:"totalIpAddresses"`
                                UsableIPAddressCount int    `json:"usableIpAddressCount"`
                                Version              int    `json:"version"`
                } `json:"primarySubnet"`
                VlanNumber int `json:"vlanNumber"`
}
