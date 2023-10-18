package models

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm/schema"
)

type StringSlice []string

func (stringSlice *StringSlice) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) error {
	value := dbValue.(string)
	*stringSlice = strings.Split(value, "\n")
	return nil
}

func (stringSlice StringSlice) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return strings.Join(stringSlice, "\n"), nil
}

type Card struct {
	ID    string      `gorm:"primaryKey"`
	Front StringSlice `gorm:"not null"`
	Back  StringSlice `gorm:"not null"`
}

func (c Card) Path() string {
	return fmt.Sprintf("/cards/%s", c.ID)
}
