package types

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

// TypesupportHandler TODO
type TypesupportHandler struct {
	handlerMap map[string]TypesupportHandlerInterface
}

var instance *TypesupportHandler
var once sync.Once
var nameSanitizerPattern = regexp.MustCompile("^[\\*]")
var emptyString = ""

// GetTypesupportHandler TODO
func GetTypesupportHandler() *TypesupportHandler {
	once.Do(func() {
		instance = &TypesupportHandler{}

		instance.handlerMap = make(map[string]TypesupportHandlerInterface)
	})
	return instance
}

// GetHandlerByMSGInterface TODO
func (typeManager *TypesupportHandler) GetHandlerByMSGInterface(msg MSGInterface) (TypesupportHandlerInterface, error) {

	// Variable declaration
	var datatype reflect.Type
	var datatypeFullString string
	var datatypeString string

	// Variable Instantiaion
	datatype = reflect.TypeOf(msg)
	datatypeFullString = datatype.String()
	datatypeString = strings.Split(datatypeFullString, ".")[0]

	return typeManager.GetHandlerByName(datatypeString)
}

// GetHandlerByName TODO
func (typeManager *TypesupportHandler) GetHandlerByName(key string) (TypesupportHandlerInterface, error) {

	// Variable declaration
	var err error
	var handler TypesupportHandlerInterface

	// Sanitize the key (required to avoid issues with copies or pointers)
	key = sanitizeKey(key)

	// Attempt to retrieve respective handler
	handler, exists := typeManager.handlerMap[key]
	if !exists {
		err = fmt.Errorf("unknown datatype \"%s\" handler, please check handler your dependencies", key)
	}

	return handler, err
}

// setHandler TODO
func (typeManager *TypesupportHandler) setHandler(key string, handler TypesupportHandlerInterface) error {

	// Variable declaration
	var err error

	// Sanitize the key (required to avoid issues with copies or pointers)
	key = sanitizeKey(key)

	// Check for duplicated registrations
	_, exists := typeManager.handlerMap[key]
	if exists {
		err = fmt.Errorf("overriding typesupport package \"%s\" handler", key)
	}

	// Register new handler
	typeManager.handlerMap[key] = handler

	return err
}

// Package key sanitizer, to avoid issues with struct copies vs pointers
func sanitizeKey(key string) string {

	// Removes any "*" at the start of any string
	if nameSanitizerPattern.MatchString(key) {
		key = nameSanitizerPattern.ReplaceAllString(key, emptyString)
	}

	return key
}
