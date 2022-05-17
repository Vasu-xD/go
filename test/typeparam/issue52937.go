// build -buildmode=plugin

//go:build !js
// +build !js

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func main()      {}
func F[T any]()  {}
func G[T any](T) {}
