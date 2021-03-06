// The MIT License (MIT)
//
// Copyright (c) 2013-2015 Oryx(ossrs)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"github.com/ossrs/go-oryx/core"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ctx := core.NewContext()
	core.Conf = core.NewConfig(ctx)

	core.Info = core.NewLogPlus(log.New(ioutil.Discard, core.LogInfoLabel, log.LstdFlags))
	core.Trace = core.NewLogPlus(log.New(ioutil.Discard, core.LogTraceLabel, log.LstdFlags))
	core.Warn = core.NewLogPlus(log.New(ioutil.Discard, core.LogWarnLabel, log.LstdFlags))
	core.Error = core.NewLogPlus(log.New(ioutil.Discard, core.LogErrorLabel, log.LstdFlags))

	os.Exit(m.Run())
}
