package common

import "syscall"

func EnvString(key, fallback string) string {
	if v, ok := syscall.Getenv(key); ok {
		return v
	} 
	return fallback
	
}