package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/drunkleen/TimeFarm-Bot/types"
	"io"
	"net/http"
	"os"
)

func GetAndSaveToken(queryData string) error {
	url := "https://tg-bot-tap.laborx.io/api/v1/auth/validate-init/v2"
	payload := map[string]string{
		"initData": queryData,
		"platform": "ios",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	var tokenInfo map[string]any
	if err := json.Unmarshal(body, &tokenInfo); err != nil {
		return fmt.Errorf("json decode error: %v", err)
	}

	tokenFile, err := os.OpenFile("./configs/tokens.conf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open ./configs/tokens.conf file: %v", err)
	}
	defer tokenFile.Close()

	if _, err := tokenFile.WriteString(fmt.Sprintf("%s\n", tokenInfo["token"])); err != nil {
		return fmt.Errorf("failed to write ./configs/tokens.conf to file: %v", err)
	}
	return nil
}

func CheckFarmingStatus(token string) (types.FarmingStatusResponse, error) {
	url := "https://tg-bot-tap.laborx.io/api/v1/farming/info"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return types.FarmingStatusResponse{}, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return types.FarmingStatusResponse{}, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.FarmingStatusResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.FarmingStatusResponse{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var farmingInfo types.FarmingStatusResponse
	if err := json.Unmarshal(body, &farmingInfo); err != nil {
		return types.FarmingStatusResponse{}, fmt.Errorf("json decode error: %v", err)
	}
	return farmingInfo, nil
}

func StartFarming(token string) (map[string]any, error) {
	url := "https://tg-bot-tap.laborx.io/api/v1/farming/start"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var farmingInfo map[string]any
	if err := json.Unmarshal(body, &farmingInfo); err != nil {
		return nil, fmt.Errorf("json decode error: %v", err)
	}
	return farmingInfo, nil
}

func FinishFarming(token string) (map[string]any, error) {
	url := "https://tg-bot-tap.laborx.io/api/v1/farming/finish"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var farmingInfo map[string]any
	if err := json.Unmarshal(body, &farmingInfo); err != nil {
		return nil, fmt.Errorf("json decode error: %v", err)
	}
	return farmingInfo, nil
}

func CheckTasks(token string) ([]types.CheckTaskItem, error) {
	url := "https://tg-bot-tap.laborx.io/api/v1/tasks"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var CheckTasksResponse []types.CheckTaskItem
	if err := json.Unmarshal(body, &CheckTasksResponse); err != nil {
		return nil, fmt.Errorf("json decode error: %v", err)
	}

	return CheckTasksResponse, nil
}

func SubmitTasks(token, taskID string) (map[string]any, error) {
	url := "https://tg-bot-tap.laborx.io/api/v1/tasks/" + taskID + "/submissions"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var ClaimTokenResponse map[string]any
	if err := json.Unmarshal(body, &ClaimTokenResponse); err != nil {
		return nil, fmt.Errorf("json decode error: %v", err)
	}
	return ClaimTokenResponse, nil
}

func ClaimTasks(token, taskID string) (map[string]any, error) {
	url := "https://tg-bot-tap.laborx.io/api/v1/tasks/" + taskID + "/claims"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var ClaimTokenResponse map[string]any
	if err := json.Unmarshal(body, &ClaimTokenResponse); err != nil {
		return nil, fmt.Errorf("json decode error: %v", err)
	}
	return ClaimTokenResponse, nil
}

func UpgradeLevel(token string) (map[string]any, error) {
	url := "https://tg-bot-tap.laborx.io/api/v1/me/level/upgrade"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var UpgradeLevelResponse map[string]any
	if err := json.Unmarshal(body, &UpgradeLevelResponse); err != nil {
		return nil, fmt.Errorf("json decode error: %v", err)
	}
	return UpgradeLevelResponse, nil
}
