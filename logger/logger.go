// Copyright 2021 Jack lei. All rights reserved.
// Use of this source code is governed by a mit license that can be found
// in the LICENSE file.
// use github.com/sirupsen/logrus
package logger

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	// do something here to set environment depending on an environment variable
	// or command-line flag
	//if Environment == "production" {
	//	log.SetFormatter(&log.JSONFormatter{})
	//} else {
	// The TextFormatter is default, you don't actually have to do this.
	log.SetFormatter(&log.TextFormatter{})
	//}
}
