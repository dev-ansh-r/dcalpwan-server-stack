package goproto

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"google.golang.org/genproto/protobuf/field_mask"
)

var nameTagRegex *regexp.Regexp

func init() {
	nameTagRegex = regexp.MustCompile("name\\=[a-zA-Z_0-9]+")
}

// GoFieldsPaths converts protobuf FieldMask paths to Go fields paths.
//
// This implementation does not support separation by ",", but only paths separated by ".".
func GoFieldsPaths(pb *field_mask.FieldMask, v any) []string {
	var newFields []string
	if len(pb.GetPaths()) == 0 {
		return newFields
	}

	goFields := goFieldsFromProtoMasks(reflect.ValueOf(v))
	for _, field := range pb.Paths {
		goName, ok := goFields[field]
		if ok {
			newFields = append(newFields, goName)
		} else {
			newFields = append(newFields, field)
		}
	}

	return newFields
}

func goFieldsFromProtoMasks(v reflect.Value) map[string]string {
	if !v.IsValid() {
		return nil
	}

	fields := make(map[string]string)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fields
	}

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("protobuf")
		if tag == "" {
			continue
		}

		protoName := strings.TrimPrefix(nameTagRegex.FindString(tag), "name=")
		if protoName == "" {
			continue
		}
		goName := v.Type().Field(i).Name
		fields[protoName] = goName

		subFields := goFieldsFromProtoMasks(v.Field(i))
		if len(subFields) == 0 {
			continue
		}

		for k, v := range subFields {
			fields[fmt.Sprintf("%s.%s", protoName, k)] = fmt.Sprintf("%s%s%s", goName, ".", v)
		}
	}

	return fields
}
