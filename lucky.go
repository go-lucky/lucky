// Copyright 2021 Jack lei. All rights reserved.
// Use of this source code is governed by a mit license that can be found
// in the LICENSE file.
package lucky

import (
	"log"
	"net/http"
)

type lucky struct {
	UserName string
}

func (l *lucky) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json;  charset=utf-8")
	w.WriteHeader(http.StatusOK)
	log.Printf("l ")
	w.Write([]byte("{\"fuck\":\"1\"}"))
}
