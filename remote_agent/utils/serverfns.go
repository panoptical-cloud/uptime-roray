package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func RegisterWithServer(url string) error {
	resp, err := http.Get(url)
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

	configFile := "/agent.conf"
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
		NatsUrl string `json:"natsUrl"`
	}{}
	err = json.Unmarshal(body, &respMsg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}
	_, err = file.WriteString(fmt.Sprintf("%s\n", respMsg))
	if err != nil {
		return fmt.Errorf("failed to write to config file: %w", err)
	}
	fmt.Printf("Registration details saved to %s\n", configFile)
	return nil
}
