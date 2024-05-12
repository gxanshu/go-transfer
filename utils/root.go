package utils

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func GetLocalIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			if ipnet.IP.IsLoopback() {
				continue
			}
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("no suitable IP address found")
}

func GetFilePath(fileName string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory", err)
		return "", err
	}

	filePath := wd + "/" + fileName
	return filePath, nil
}
