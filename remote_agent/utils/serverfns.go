package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func RegisterWithServer(url string) error {
	pu := &PsUtils{}
	macId, err := pu.GetSelfMachineId()
	if err != nil {
		return fmt.Errorf("failed to get machine id: %w", err)
	}

	rb := &RegPostReqBody{
		Version:   "1.0.0",
		MachineId: *macId,
	}

	jb, err := json.Marshal(rb)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	fmt.Printf("%#v", rb)

	br := bytes.NewReader(jb)
	resp, err := http.Post(url, "application/json", br)
	if err != nil {
		fmt.Println("Error while registering with server", err)
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading response body", err)
		return err
	}
	fmt.Println("Response from server:", string(body))

	configFile := "/agent.toml"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	agentHomePath := filepath.Join(homeDir, ".roray_panmon")
	err = os.MkdirAll(agentHomePath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create roray_panmon config directory in user home directory: %w", err)
	}
	file, err := os.Create(agentHomePath + configFile)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()
	respMsg := struct {
		NatsUrl     string `json:"natsUrl"`
		NatsSubject string `json:"subj"`
	}{}
	err = json.Unmarshal(body, &respMsg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// Create TOML content
	tomlContent := fmt.Sprintf("nats_url = %q\nnats_subject = %q\n", respMsg.NatsUrl, respMsg.NatsSubject)

	// Write TOML content to file
	_, err = file.WriteString(tomlContent)
	if err != nil {
		return fmt.Errorf("failed to write to config file: %w", err)
	}

	fmt.Printf("Registration details saved to %s\n", configFile)
	return nil
}
