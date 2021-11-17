package main

import (
	"errors"
	"fmt"
	"strings"
)

var CATEGORIES_VAL_SPLIT_CHAR = "+"

type CategoryLine struct {
	org_id     string
	asset_name string
	version    string
	field_key  string
	field_val  string
}

func newCategoryLine(line map[string]string) (*CategoryLine, error) {
	c := &CategoryLine{}
	if val, ok := line["orgId"]; ok {
		c.org_id = val
	} else {
		return nil, errors.New("\t[ERROR] orgId attribute is missing")
	}
	if val, ok := line["assetName"]; ok {
		c.asset_name = val
	} else {
		return nil, errors.New("\t[ERROR] assetName attribute is missing")
	}
	if val, ok := line["version"]; ok {
		c.version = val
	} else {
		return nil, errors.New("\t[ERROR] version attribute is missing")
	}
	if val, ok := line["fieldKey"]; ok {
		c.field_key = val
	} else {
		return nil, errors.New("\t[ERROR] fieldKey attribute is missing")
	}
	if val, ok := line["fieldVal"]; ok {
		c.field_val = val
	} else {
		return nil, errors.New("\t[ERROR] fieldVal attribute is missing")
	}

	return c, nil
}

func (c *ExchangeClient) handleCategories(file string) error {
	categories, err := CSVFileToMap(file)
	if err != nil {
		return err
	}
	if len(categories) <= 0 {
		fmt.Printf("\tcategory file is empty.\n")
		return nil
	}

	for _, line := range categories {
		cat, err := newCategoryLine(line)
		if err != nil {
			return err
		}
		if err := c.handleCategory(cat); err != nil {
			return err
		}
	}
	return nil
}

func (c *ExchangeClient) handleCategory(category *CategoryLine) error {
	if category == nil {
		return errors.New("\t[ERROR] category line is empty")
	}

	fmt.Printf("Processing Category \n\torg: %s\n\tasset: %s\n\tversion: %s\n\tkey: %s\n\tvalue: %s\n\n", category.org_id, category.asset_name, category.version, category.field_key, category.field_val)

	if err := c.createCustomCategory(category.org_id, category.asset_name, category.version, category.field_key, strings.Split(category.field_val, CATEGORIES_VAL_SPLIT_CHAR)); err != nil {
		return err
	}

	return nil
}
