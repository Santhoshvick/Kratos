package handler

import (
	"errors"
	"log"
	"strconv"
)


func BalanceValidation(availableBalance string) error{
	avalBalance, err := strconv.ParseInt(availableBalance, 10, 64)
	if err != nil {
		log.Fatalf("Issue with conveting the available Balance %v", err)
	}
	if availableBalance==""{
		return errors.New("the available balance should not be less than 0 or should not be negative")
	}

	if avalBalance < 0 {
		return errors.New("the available balance should not be less than 0 or should not be negative")
	}

	return nil

}


func CurrencyValidation(currency string) error{
	arr:=[]string{"RUPEE","USD","EURO","DIRHAM"}

	for i:=0;i<len(arr);i++{
		val:=arr[i]
		if val==currency{
			return errors.New("please verify the currency Once.Allowed Curencies are USD  EURO  RUPEE   DIRHAM")
		}
	}
	return nil

}


func StatusValidations(status string) error{
	arr:=[]string{"ACTIVE","SUSPENDED","CLOSED"}
		for i:=0;i<len(arr);i++{
		val:=arr[i]
		if val==status{
			return errors.New("please verify the status Once.Allowed Status are ACTIVE  SUSPENDED  CLOSED ")
		}
	}
	return nil


}


