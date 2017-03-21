package utils

import (
  "reflect"
  //"strconv"

  "github.com/urfave/cli"

)

// InputCheckType validates the type of arguments passed to commands. It takes
// a generic argument and an expected type. If the type of the argument matches
// the exptected type it returns true, otherwise it returns false.
//
// NOTE: This function was intended to be the beginning of input validation
// for arguments passed to commands. It is still experimental and currently
// does not work.
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
