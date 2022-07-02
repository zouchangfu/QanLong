package utils

import (
	"context"
	"fmt"
	"github.com/Ullaakut/nmap/v2"
	systemReq "github.com/zouchangfu/QanLong/model/business"
	"log"
	"strconv"
	"strings"
	"time"
)

func StartDetect(t systemReq.Task) (systemReq.AssetDetectResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	var detectResult systemReq.AssetDetectResult

	scanner, err := nmap.NewScanner(
		nmap.WithBinaryPath("D:\\dev\\Nmap\\nmap.exe"),
		nmap.WithTargets(t.Target),
		nmap.WithCustomArguments("-O"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
		return detectResult, err
	}

	// start detect
	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
		return detectResult, err
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		fmt.Printf("Host %q,OS: %q:\n", host.Addresses[0], host.OS.Matches[0].Name)
		detectResult.Ip = host.Addresses[0].String()
		detectResult.OperatingSystem = host.OS.Matches[0].Name
		var portArr []string
		for _, port := range host.Ports {
			portArr = append(portArr, strconv.Itoa(int(port.ID)))
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
		allPort := strings.Join(portArr, ",")
		detectResult.OpenPort = allPort
		fmt.Printf("%v", detectResult)
	}

	return detectResult, nil
}
