package main

import (
  "fmt"
  "os"
  "github.com/codegangsta/cli"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "strings"
)

func main() {
  app := cli.NewApp()
  app.Name = "sl"
  app.Usage = "a softlayer command line utility"
  app.Action = func(c *cli.Context) {
	println("do something")
  }

  app.Commands = []cli.Command {
	{
		Name:	"devices",
		Aliases:[]string{"d"},
		Usage:	"List all devices within softlayer",
		Action: func(c *cli.Context) {
			config, err := ioutil.ReadFile("/home/vagrant/.sl/config")
			if err != nil {
			    println("Unable to read configuration file /home/vagrant/.sl/config")
			}
			configString := strings.TrimSpace(string(config))
			urlString := fmt.Sprint("https://", configString, "/rest/v3/SoftLayer_Account/VirtualGuests.json")

			// Build the request
			req, err := http.NewRequest("GET", urlString, nil)
			if err != nil {
			    println("Unable to create request")
			}

			req.ParseForm()
			req.Form.Add("objectMask","id;hostname;provisionDate;primaryIpAddress;primaryBackendIpAddress;operatingSystem.passwords;networkComponents;notes")

			// Send the request via a client
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
			    println("Unable to perform request")
			}

			// Defer the closing of the body
			defer resp.Body.Close()
			// Read the content into a byte array
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
			    println("Unable to read response")
			}

			// Fill the record with the data from the JSON
			var virtualGuests []VirtualGuest
		    err = json.Unmarshal(body, &virtualGuests)
		    if err != nil {
		        // An error occurred while converting our JSON to an object
		        println("Unable to unmarshal response")
		    }

			fmt.Printf("%-12s", "ID")
			fmt.Printf("%-30s", "Hostname")
			fmt.Printf("%-18s", "Private IP")
			fmt.Printf("%-18s", "Public IP")
			fmt.Printf("%-30s", "Provision Date")
			fmt.Printf("%-30s\n", "Notes")

			for _,each := range virtualGuests {
			    fmt.Printf("%-12d", each.ID)
			    fmt.Printf("%-30s", each.Hostname)
			    fmt.Printf("%-18s", each.PrimaryBackendIPAddress)
			    fmt.Printf("%-18s", each.PrimaryIPAddress)
			    fmt.Printf("%-30s", each.ProvisionDate)
			    fmt.Printf("%-30s\n", each.Notes)
			}
		},
	},
	{
		Name:	"domains",
		Aliases:[]string{"do"},
		Usage:	"List all domains within softlayer",
		Action: func(c *cli.Context) {
			println("...listing domains")
		},
	},
	{
		Name:	"vlans",
		Aliases:[]string{"v"},
		Usage:	"List all VLANs",
		Action: func(c *cli.Context) {
			println("...listing VLANs")
		},
	},
  }

  app.Run(os.Args)
}

func promptForUsername() string {
	fmt.Println("Username> ")
	return ""
}

func promptForApiKey() string {
	fmt.Println("Username> ")
	return ""
}

func saveDetails(username string, api_key string) {
	println("Saving details...")
}
