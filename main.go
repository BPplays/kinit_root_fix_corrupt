package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)



func k_dest () {
	cmd := exec.Command("kdestroy", "-A")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println("Error running kdestroy root", ":", err)
	} else {
		log.Println("kdestroy:", "root")
	}
}


func main() {
	bad_msg := "kinit: Failed to store credentials: Internal credentials cache error while getting initial credentials"
	var loops int64

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))


	// file := "/etc/krb5.keytab"
	for {

		// dirPath := filepath.Join("/etc/krb5")
		keytabFile := filepath.Join("/etc/krb5.keytab")

		if _, err := os.Stat(keytabFile); err == nil {
			loops = 0
			for {
				if loops > 5 {
					break
				}

				// username := "root"

				// chown_r(dirPath, username, username)

				// chmod_r(dirPath, "700")

				cmd := exec.Command("kinit", "-k", "-t", keytabFile)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				out, err := cmd.CombinedOutput()
				if string(out) == bad_msg {
					k_dest()
					loops += 1
				} else {
					break
				}
				if err != nil {
					log.Println("Error running kinit for", "root", ":", err)
				} else {
					log.Println("Ran kinit for", "root")
				}
			}
		}
		
		
		time.Sleep(3 * time.Hour)
	}

}
