package main

import (
	"errors"
	"fmt"
)

type TagLine struct {
	org_id     string
	asset_name string
	version    string
	field_key  string
	field_val  string
	field_type string
}

func newTagLine(line map[string]string) (*TagLine, error) {
	t := TagLine{}
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
	if val, ok := line["version"]; ok {
		t.version = val
	} else {
		return nil, errors.New("\t[ERROR] version attribute is missing")
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
	if val, ok := line["fieldType"]; ok {
		t.field_type = val
	} else {
		return nil, errors.New("\t[ERROR] fieldType attribute is missing")
	}

	return &t, nil
}

func (c *ExchangeClient) handleTags(file string) error {
	tags, err := CSVFileToMap(file)
	if err != nil {
		return err
	}

	if len(tags) <= 0 {
		fmt.Printf("\tTags file is empty. Skipping.\n")
		return nil
	}

	for _, line := range tags {
		tag, err := newTagLine(line)
		if err != nil {
			return err
		}
		if err := c.handleTag(tag); err != nil {
			return err
		}
	}

	return nil
}

func (c *ExchangeClient) handleTag(tag *TagLine) error {
	ctx := *c.ctx
	if tag == nil {
		return errors.New("\t[ERROR] tag line is empty")
	}

	fmt.Printf("Processing Tag \n\torg:  %s\n\tasset: %s\n\tversion: %s\n\ttype : %s\n\tkey: %s\n\tvalue: %s\n\n", tag.org_id, tag.asset_name, tag.version, tag.field_type, tag.field_key, tag.field_val)

	parentOrg := ctx.Value(ANYPOINT_ORG_KEY).(string)

	if err := c.addCustomField(parentOrg, tag.field_type, tag.field_key, tag.field_key); err != nil {
		return err
	}

	if err := c.putCustomFieldValue(tag.org_id, tag.asset_name, tag.version, tag.field_key, tag.field_val); err != nil {
		return err
	}

	return nil
}
