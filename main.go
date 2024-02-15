package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)


func chown_r (file string, username string, group string) {
	cmd := exec.Command("chown", fmt.Sprintf("%s:%s", username, group), "-R", file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println("Error running chown for", file, ":", err)
	} else {
		log.Println("chowned", file)
	}
}

func k_dest (user string) {
	cmd := exec.Command("sudo", "-u", user, "kdestroy", "-A")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println("Error running kdestroy for", user, ":", err)
	} else {
		log.Println("kdestroy:", user)
	}
}

func chmod_r (file string, perms string) {
	cmd := exec.Command("chmod", perms, "-R", file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println("Error running chmod for", file, ":", err)
	} else {
		log.Println("chmoded", file)
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

				username := "root"

				// chown_r(dirPath, username, username)

				// chmod_r(dirPath, "700")

				cmd := exec.Command("kinit", "-k", "-t", keytabFile)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				out, err := cmd.CombinedOutput()
				if string(out) == bad_msg {
					k_dest(username)
					loops += 1
				} else {
					break
				}
				if err != nil {
					log.Println("Error running kinit for", username, ":", err)
				} else {
					log.Println("Ran kinit for", username)
				}
			}
		}
		
		
		time.Sleep(3 * time.Hour)
	}

}
