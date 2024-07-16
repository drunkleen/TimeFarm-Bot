package main

import (
	"flag"
	"fmt"
	"github.com/drunkleen/TimeFarm-Bot/requests"
	"github.com/drunkleen/TimeFarm-Bot/utils"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	tokenList           []string
	generateTokenEnable bool
	//upgradeEnable       bool
	taskAutoClaimEnable bool
)

func main() {
	utils.ClearScreen()
	utils.PrintLogo()
	parseArgs()

	if generateTokenEnable {
		queries, err := utils.ParseQueries()
		if err != nil {
			panic(err)
		}
	query:
		for _, queryID := range queries {
			err := requests.GetAndSaveToken(queryID)
			if err != nil {
				fmt.Printf(utils.Red("Error generating token for query:\n%s\n"), queryID)
				continue query
			}
			err = utils.DeleteQuery(queryID)
			if err != nil {
				fmt.Printf(utils.Red("Error deleting query from list:\n%s\n"), queryID)
			}
		}
	}
	var err error
	tokenList, err = utils.ParseTokens()
	if err != nil {
		log.Fatal(err)
	}
	runLoop()
}

func runLoop() {
	var startTime time.Time = time.Now()
	var printString string
	for {
		for tokenIndex, tokenVal := range tokenList {
			printString += fmt.Sprintf(utils.Green("--------------------- Account: %d ----------------------\n"), tokenIndex)
			status, err := requests.CheckFarmingStatus(tokenVal)
			if err != nil {
				log.Printf("Error checking farming status for token %d: %v\n", tokenIndex+1, err)
				continue
			}

			if taskAutoClaimEnable {
				tasksResp, err := requests.CheckTasks(tokenVal)
				if err != nil {
					printString += fmt.Sprintf(utils.Red("Error checking tasks %v\n"), err)
				}

			taskResponses:
				for _, t := range tasksResp {
					if t.Submission.Status == "CLAIMED" || t.Title == "Connect TON wallet" {
						continue taskResponses
					}
					if t.Submission.Status == "SUBMITTED" {
						_, err := requests.ClaimTasks(tokenVal, t.Id)
						if err != nil {
							printString += fmt.Sprintf(utils.Red("Error claiming task: %s\n"), t.Title)
						}
						continue taskResponses
					}
					_, err := requests.SubmitTasks(tokenVal, t.Id)
					if err != nil {
						printString += fmt.Sprintf(utils.Red("Error submitting task: %v\n"), t.Title)
					}
				}
			}

			_, _ = requests.UpgradeLevel(tokenVal)

			farmingResponse, err := requests.FinishFarming(tokenVal)
			if err != nil {
				printString += fmt.Sprintf(utils.Red("Error finishing farming %v\n"), err)
				continue
			}

			if errorValue, ok := farmingResponse["error"].(map[string]any); ok {
				if message, ok := errorValue["message"].(string); ok {
					if message == "Farming didn't start" {
						_, err := requests.StartFarming(tokenVal)
						if err != nil {
							printString += fmt.Sprintf("Error starting farming")
						}
						printString += fmt.Sprintf("Farming started successfully!\n")
					}
				}
			}

			startedAt, err := time.Parse(time.RFC3339, status.ActiveFarmingStartedAt)
			if err != nil {
				fmt.Printf("Error parsing time: %v\n", err)
				continue
			}

			duration := time.Duration(status.FarmingDurationInSec) * time.Second
			BalanceFloat, _ := strconv.ParseFloat(status.Balance, 64)

			printString += fmt.Sprintf(
				"Balance: %.0f | Start: %v | End: %v\n",
				BalanceFloat,
				startedAt.Format(time.Kitchen),
				utils.FormatLeftDuration(startedAt.Add(duration).Sub(time.Now())),
			)
		}

		utils.ClearScreen()
		utils.PrintLogo()
		fmt.Printf(utils.Cyan("                  Current Time: %v\n"), time.Now().UTC().Format(time.Kitchen))
		fmt.Println(printString)
		fmt.Println(utils.Green("-------------------------------------------------------"))

		fmt.Printf(utils.Cyan("                      Up Time: %v\n"), utils.FormatLeftDuration(time.Since(startTime)))
		printString = ""
		time.Sleep(5 * time.Second)

	}
}

func parseArgs() {
	generateToken := flag.String("generate-query", "", "Do you want generate tokens from ./configs/query.conf? (y/N): ")
	//autoUpgradeClock := flag.String("upgrade", "", "Do you want auto upgrade clock? (y/N)")
	tasks := flag.String("task", "", "Do you want to auto claim tasks? (Y/n): ")
	flag.Parse()

	if *generateToken == "" {
		fmt.Print("Do you want generate tokens from ./configs/query.conf? (y/N): ")
		var generateTokenInput string
		fmt.Scanln(&generateTokenInput)
		generateTokenInput = strings.TrimSpace(strings.ToLower(generateTokenInput))

		if generateTokenInput == "y" || generateTokenInput == "yes" {
			generateTokenEnable = true
		}
	}
	//if *autoUpgradeClock == "" {
	//	fmt.Print("Do you want auto upgrade clock? (y/N): ")
	//	var upgradeInput string
	//	fmt.Scanln(&upgradeInput)
	//	upgradeInput = strings.TrimSpace(strings.ToLower(upgradeInput))
	//
	//	if upgradeInput == "y" || upgradeInput == "yes" {
	//		upgradeEnable = true
	//	}
	//}
	if *tasks == "" {
		fmt.Print("Do you want to auto claim tasks? (Y/n): ")
		var taskInput string
		fmt.Scanln(&taskInput)
		taskInput = strings.TrimSpace(strings.ToLower(taskInput))

		if taskInput == "" || taskInput == "y" || taskInput == "yes" {
			taskAutoClaimEnable = true
		}
	}
}
