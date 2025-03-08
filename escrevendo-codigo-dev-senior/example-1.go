package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Junior Approach - Incomplete Implementation
func fetchUserDataJunior(userID string) (map[string]interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("/api/users/%s", userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Senior Approach - Complete Implementation
func fetchUserDataSenior(userID string) (map[string]interface{}, error) {
	if userID == "" {
		return nil, errors.New("User ID is required")
	}

	resp, err := http.Get(fmt.Sprintf("/api/users/%s", userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error! status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	// Validate expected data structure
	_, existsID := data["id"]
	_, existsName := data["name"]

	if !existsID || !existsName {
		return nil, fmt.Errorf("invalid user data received: %s", data)
	}

	return data, nil
}

func main() {
	// example of junior implementation.
	juniorData, juniorErr := fetchUserDataJunior("123")
	if juniorErr != nil {
		log.Println("Junior Error:", juniorErr)
	} else {
		log.Println("Junior Data:", juniorData)
	}

	// example of senior implementation
	seniorData, seniorErr := fetchUserDataSenior("123")

	if seniorErr != nil {
		log.Println("Senior Error:", seniorErr)
		// error handling and recovery if necessary
	}
	log.Println("Senior Data:", seniorData)

	// testing for error case
	seniorData, seniorErr = fetchUserDataSenior("")
	if seniorErr != nil {
		log.Println("Senior Error (empty ID):", seniorErr)
	}
}
