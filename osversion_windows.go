package main

import (
	"golang.org/x/sys/windows/registry"
	"log"
)

func GetWindowsVersion() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	version, _, err := k.GetStringValue("SystemRoot")
	return version, err
}

