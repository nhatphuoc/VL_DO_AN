package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type ModelBool bool

func (n *ModelBool) Scan(value interface{}) error {
	inter, ok:= value.(int64)

	if ok!=true {
		return errors.New(fmt.Sprintf("Error scan ModelBool,%T",value))
	}

	if inter == 0{
		*n = ModelBool(false)
	} else {
		*n = ModelBool(true)

	}

	return nil
}

func (n *ModelBool) Value() (driver.Value, error) {
	value := bool(*n)
	if value {
		return true,nil
	} else {
		return false,nil
	}
}

func (n *ModelBool) MarshalJSON() ([]byte, error) {
	var v string
	if bool(*n) {
		v="true"
	} else{
		v="false"
	}
	return []byte(fmt.Sprintf("%s", v)), nil
}

func (n *ModelBool) UnmarshalJSON(data []byte) error {
	str := string(data)
	fmt.Print(data)
	if str == "true" {
		*n = ModelBool(true)
	} else {
		*n = ModelBool(false)
	}
	return nil
}