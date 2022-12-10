package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const Port = 8090
const ShellToUse = "bash"

func runCommand(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return stdout.String(), stderr.String(), errors.New(fmt.Sprintf("Error running command: %s", err))
	}

	return stdout.String(), stderr.String(), nil
}

func runShortcut(shortcutName string) error {
	command := fmt.Sprintf("open 'shortcuts://run-shortcut?name=%s'", shortcutName)
	log.Printf("Command: %s\n", command)

	out, errOut, err := runCommand(command)
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	log.Println("--- stdout ---")
	log.Println(out)

	log.Println("--- stderr ---")
	log.Printf("%v\n\n", errOut)

	return err
}

func runShortcutHandler(w http.ResponseWriter, req *http.Request) {
	shortcutName := req.URL.Query().Get("name")
	err := runShortcut(shortcutName)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/execute", runShortcutHandler)

	log.Printf("Server started on port %d", Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
