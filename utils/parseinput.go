package utils

import (
  "reflect"
  //"strconv"

  "github.com/urfave/cli"

)

func InputCheckType(c *cli.Context, arg_number, flag_number int, expected_type string) (bool, error) {
  if arg_number != 0 {
    input_type := reflect.TypeOf(c.Args().Get(arg_number))
		switch {
		case expected_type == input_type.String():
			return true, nil
		case expected_type != input_type.String():
			return false, nil
		}
  }
  if flag_number != 0 {
    input_type := reflect.TypeOf(c.Args().Get(flag_number))
		switch {
		case expected_type == input_type.String():
			return true, nil
		case expected_type != input_type.String():
			return false, nil
		}
  }
	return false, nil // Add actual error here
}

//func InputCheckContent(c *cli.Context, )
