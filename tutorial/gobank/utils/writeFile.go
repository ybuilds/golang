package utils

import (
	"errors"
	"fmt"
	"os"
)

func WriteToFile(balance float64) error {
	var balanceText = fmt.Sprint(balance);
	
	var err = os.WriteFile(FilePath, []byte(balanceText), 0777);
	if err != nil {
		fmt.Printf("ERR[%s]\n", err);
		return errors.New("Error writing balance to file");
	}

	return nil;
}