package main

type VirtualGuest struct {
                Hostname          string `json:"hostname"`
                ID                int    `json:"id"`
                NetworkComponents []struct {
                                CreateDate       string `json:"createDate"`
                                GuestID          int    `json:"guestId"`
                                ID               int    `json:"id"`
                                MacAddress       string `json:"macAddress"`
                                MaxSpeed         int    `json:"maxSpeed"`
                                ModifyDate       string `json:"modifyDate"`
                                Name             string `json:"name"`
                                NetworkID        int    `json:"networkId"`
                                Port             int    `json:"port"`
                                PrimaryIPAddress string `json:"primaryIpAddress"`
                                Speed            int    `json:"speed"`
                                Status           string `json:"status"`
                                UUID             string `json:"uuid"`
                } `json:"networkComponents"`
                OperatingSystem struct {
                                HardwareID                  interface{} `json:"hardwareId"`
                                ID                          int         `json:"id"`
                                ManufacturerLicenseInstance string      `json:"manufacturerLicenseInstance"`
                                Passwords                   []struct {
                                                CreateDate string      `json:"createDate"`
                                                ID         int         `json:"id"`
                                                ModifyDate string      `json:"modifyDate"`
                                                Password   string      `json:"password"`
                                                Port       interface{} `json:"port"`
                                                Software   struct {
                                                                HardwareID                  interface{} `json:"hardwareId"`
                                                                ID                          int         `json:"id"`
                                                                ManufacturerLicenseInstance string      `json:"manufacturerLicenseInstance"`
                                                                Passwords                   []struct {
                                                                                CreateDate string      `json:"createDate"`
                                                                                ID         int         `json:"id"`
                                                                                ModifyDate string      `json:"modifyDate"`
                                                                                Password   string      `json:"password"`
                                                                                Port       interface{} `json:"port"`
                                                                                Software   interface{} `json:"software"`
                                                                                SoftwareID int         `json:"softwareId"`
                                                                                Username   string      `json:"username"`
                                                                } `json:"passwords"`
                                                                SoftwareLicense struct {
                                                                                ID                  int `json:"id"`
                                                                                SoftwareDescription struct {
                                                                                                ControlPanel                 int         `json:"controlPanel"`
                                                                                                ID                           int         `json:"id"`
                                                                                                LicenseTermValue             interface{} `json:"licenseTermValue"`
                                                                                                LongDescription              string      `json:"longDescription"`
                                                                                                Manufacturer                 string      `json:"manufacturer"`
                                                                                                Name                         string      `json:"name"`
                                                                                                OperatingSystem              int         `json:"operatingSystem"`
                                                                                                ReferenceCode                string      `json:"referenceCode"`
                                                                                                RequiredUser                 string      `json:"requiredUser"`
                                                                                                UpgradeSoftwareDescriptionID interface{} `json:"upgradeSoftwareDescriptionId"`
                                                                                                UpgradeSwDescID              interface{} `json:"upgradeSwDescId"`
                                                                                                Version                      string      `json:"version"`
                                                                                                VirtualLicense               int         `json:"virtualLicense"`
                                                                                                VirtualizationPlatform       int         `json:"virtualizationPlatform"`
                                                                                } `json:"softwareDescription"`
                                                                                SoftwareDescriptionID int `json:"softwareDescriptionId"`
                                                                } `json:"softwareLicense"`
                                                } `json:"software"`
                                                SoftwareID int    `json:"softwareId"`
                                                Username   string `json:"username"`
                                } `json:"passwords"`
                                SoftwareLicense struct {
                                                ID                  int `json:"id"`
                                                SoftwareDescription struct {
                                                                ControlPanel                 int         `json:"controlPanel"`
                                                                ID                           int         `json:"id"`
                                                                LicenseTermValue             interface{} `json:"licenseTermValue"`
                                                                LongDescription              string      `json:"longDescription"`
                                                                Manufacturer                 string      `json:"manufacturer"`
                                                                Name                         string      `json:"name"`
                                                                OperatingSystem              int         `json:"operatingSystem"`
                                                                ReferenceCode                string      `json:"referenceCode"`
                                                                RequiredUser                 string      `json:"requiredUser"`
                                                                UpgradeSoftwareDescriptionID interface{} `json:"upgradeSoftwareDescriptionId"`
                                                                UpgradeSwDescID              interface{} `json:"upgradeSwDescId"`
                                                                Version                      string      `json:"version"`
                                                                VirtualLicense               int         `json:"virtualLicense"`
                                                                VirtualizationPlatform       int         `json:"virtualizationPlatform"`
                                                } `json:"softwareDescription"`
                                                SoftwareDescriptionID int `json:"softwareDescriptionId"`
                                } `json:"softwareLicense"`
                } `json:"operatingSystem"`
                PrimaryBackendIPAddress string `json:"primaryBackendIpAddress"`
                PrimaryIPAddress        string `json:"primaryIpAddress"`
                ProvisionDate           string `json:"provisionDate"`
                Notes           		string `json:"notes"`
}
