package plugin

import (
	"encoding/base64"
	"fmt"
)

func pluginReadTransformer(data map[string]interface{}) error {
	val, ok := data["Value"]
	if !ok {
		return fmt.Errorf("'Value' not found in reading data")
	}
	v, ok := val.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unable to convert %T to map[string]interface", val)
	}
	if len(v) != 1 {
		return fmt.Errorf("unexpected size of values map")
	}

	var value interface{}
	for _, x := range v {
		value = x
		break
	}

	data["value"] = value
	delete(data, "Value")
	return nil
}

func pluginWriteDataTransformer(data map[string]interface{}) error {
	c, ok := data["context"]
	if !ok {
		// If no context is specified, there is nothing to transform.
		return nil
	}

	ctx, ok := c.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unable to convert %T to map[string]interface{}", c)
	}

	d, ok := ctx["data"]
	if !ok {
		// If no data in the context, nothing to do.
		return nil
	}

	b, ok := d.(string)
	if !ok {
		return fmt.Errorf("unable to convert %T to []byte", d)
	}

	res, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return err
	}

	ctx["data"] = string(res)
	return nil
}
