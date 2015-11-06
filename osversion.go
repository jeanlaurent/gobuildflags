// +build !windows

package main

func GetWindowsVersion() (string, error) {
	return "You're not on windows, funny man", nil
}

