package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadBalanceFromFile() (float64, error){
	data, err := os.ReadFile(FilePath);
	if err != nil {
		fmt.Printf("ERR[%s]\n", err);
		return 0, errors.New("Error in reading file");
	}

	var balanceText = string(data);
	balance, err := strconv.ParseFloat(balanceText, 64);
	if err != nil {
		fmt.Printf("ERR[%s]\n", err);
		return 0, errors.New("Error parsing balance from file");
	}

	return balance, nil;
}