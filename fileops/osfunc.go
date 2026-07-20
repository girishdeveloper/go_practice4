package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Temp directory on this machine:", os.TempDir())
	userHomeDirectory, err := os.UserHomeDir()
	raiseError(err)
	fmt.Println("home directory of the current user:", userHomeDirectory)
	cacheRootDirectory, err := os.UserCacheDir()
	raiseError(err)
	fmt.Println("The default root directory for user-specific cache directory on this machine:", cacheRootDirectory)
	userConfigDirectory, err := os.UserConfigDir()
	raiseError(err)
	fmt.Println("The default root directory for user-specific configuration data on this machine:", userConfigDirectory)

	f1, err := os.Lstat("/home/girish/Downloads")
	raiseError(err)
	f2, err := os.Lstat("/home/girish/Documents")
	raiseError(err)
	fmt.Println("Are these same files?", os.SameFile(f1, f2))
}
func raiseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
