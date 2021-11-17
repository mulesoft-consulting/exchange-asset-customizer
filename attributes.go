package main

import (
	"errors"
	"fmt"
)

type AttributeLine struct {
	org_id     string
	asset_name string
	field_key  string
	field_val  string
}

func newAttributeLine(line map[string]string) (*AttributeLine, error) {
	t := AttributeLine{}
	if val, ok := line["orgId"]; ok {
		t.org_id = val
	} else {
		return nil, errors.New("\t[ERROR] orgId attribute is missing")
	}
	if val, ok := line["assetName"]; ok {
		t.asset_name = val
	} else {
		return nil, errors.New("\t[ERROR] assetName attribute is missing")
	}
	if val, ok := line["fieldKey"]; ok {
		t.field_key = val
	} else {
		return nil, errors.New("\t[ERROR] fieldKey attribute is missing")
	}
	if val, ok := line["fieldVal"]; ok {
		t.field_val = val
	} else {
		return nil, errors.New("\t[ERROR] fieldVal attribute is missing")
	}

	return &t, nil
}

func (c *ExchangeClient) handleAttributes(file string) error {
	tags, err := CSVFileToMap(file)
	if err != nil {
		return err
	}

	if len(tags) <= 0 {
		fmt.Printf("\tAttributes file is empty. Skipping.\n")
		return nil
	}

	for _, line := range tags {
		attr, err := newAttributeLine(line)
		if err != nil {
			return err
		}
		if err := c.handleAttribute(attr); err != nil {
			return err
		}
	}

	return nil
}

func (c *ExchangeClient) handleAttribute(tag *AttributeLine) error {
	if tag == nil {
		return errors.New("\t[ERROR] tag line is empty")
	}

	fmt.Printf("Processing Attribute \n\torg:  %s\n\tasset: %s\n\tkey: %s\n\tvalue: %s\n\n", tag.org_id, tag.asset_name, tag.field_key, tag.field_val)

	if err := c.patchAssetAttributes(tag.org_id, tag.asset_name, tag.field_key, tag.field_val); err != nil {
		return err
	}

	return nil
}
