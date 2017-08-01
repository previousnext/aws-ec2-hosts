package main

import (
	"log"
	"time"

	"github.com/lextoumbourou/goodhosts"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cliRegion = kingpin.Flag("region", "Region to query instances").Default("ap-southeast-2").String()
	cliTpl    = kingpin.Flag("tpl", "Template for host file entry name").Default("{{ .Name }}").String()
	cliFilter = kingpin.Flag("filter", "Filter instances by tag and value eg. TAG=value").Required().String()
	cliTag    = kingpin.Flag("tag", "Tag value for hosts record eg. Name=foo will result in '127.0.0.1 foo'").Required().String()
)

func main() {
	kingpin.Parse()

	log.Println("starting EC2 hosts sync")

	limiter := time.Tick(time.Minute)

	for {
		<-limiter

		// Lookup the EC2 instances with a tag.
		instances, err := queryInstanaces(*cliRegion, *cliFilter, *cliTag)
		if err != nil {
			log.Println("failed to query instances:", err)
			continue
		}

		// Load this instances host file so we can update it.
		hostFile, err := goodhosts.NewHosts()
		if err != nil {
			log.Println("failed to load hosts file:", err)
			continue
		}

		// Add the records to the hosts file.
		for _, instance := range instances {
			name, err := hostname(*cliTpl, instance.Name)
			if err != nil {
				log.Println("unable create hostname:", err)
				continue
			}

			// Clear out the hostfile records which relate to our EC2 query.
			for _, line := range hostFile.Lines {
				if contains(line.Hosts, name) {
					hostFile.Remove(line.IP, name)
				}
			}

			log.Println("adding host:", name, "/", instance.IP)

			err = hostFile.Add(instance.IP, name)
			if err != nil {
				log.Println("unable to sync hosts record:", instance.Name, "/", instance.IP)
				continue
			}
		}

		// Write the file back to disk.
		err = hostFile.Flush()
		if err != nil {
			log.Println("failed to write file hosts file:", err)
			continue
		}

		log.Println("updated hosts file")
	}
}
