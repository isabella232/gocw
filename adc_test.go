// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gocw_test

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/google/gocw"
)

func TestProcessData(t *testing.T) {
	adc := gocw.Adc{}
	// input, output captured from OpenADCInterface.py:processData() at runtime.
	const dataStr = "ace748ee09e2091a4be047edf7e1595e2ee0583e10e1482628e6b9060ddf6922" +
		"40e7594630dff8123be1c891f2e0d821fee2f9b244e1283237e688fa0adf29123ee7393e2ddfd80a" +
		"3ae1988deee0d8320ee2a9a242e1282a34e6b9020ae0b8f23ee6f99a37e2591e80df47e60ce5e8c5" +
		"eadf888633e4289a02e1585a1ca1f9221ba3683a1da1082616a1f829eba0281a009f98fe37a2e785" +
		"e6a729a656a6e9c675a659be6ea0683e51a0875df6a2e84247a4e95651a6b94656a3e8d674a42892" +
		"5aa4690650a5e976569c9726029fc79df3a5d94a329d3826619eb7e1cea578e603a0c9a2649fb759" +
		"daa3b811eda6799e5b9dc77a0d9f3779f1a3281e02a5891616a6495a4aa0e9ae6c9fd751d6a3c815" +
		"efa68992599d47660da077bdfba6696a3e9de83e6a9f07f9daa5e8f607a0e9aa699f1779dfa027cd" +
		"dda178ce07a4b96246a6b992539d583a6a9f27f5d4a5a8f607a0f9aa669ff759daa40829f2a6c9a2" +
		"5c9d977a10a0a7cdffa6b97a429e58466d9de7d1e2a0880df4a4886232a5592e5aa6c9b2649d8766" +
		"0ea047c1fea65966389d6836699f37f9d4a5d8f207a119b269a00759dca40829f3a6b9a25e9de786" +
		"0f9e77c9f3a58899dfa749966a9fe852369fe7b9dda00855caa598da1ba6493a58a488e634a0d8e6" +
		"4e9f7912139f7809faa368fa16a4786a15a047f1faa4186df8a1a84635a017fa4aa1b7f207a1b8e2"
	var expected = []float64{
		0.0302734375, 0.0283203125, 0.013671875, 0.052734375, 0.021484375, 0.0087890625,
		0.015625, -0.0205078125, 0.009765625, 0.0302734375, 0.0, 0.005859375, 0.001953125,
		0.0537109375, 0.0615234375, -0.0068359375, -0.025390625, -0.0302734375, 0.044921875,
		0.083984375, 0.1025390625, 0.111328125, 0.1142578125, 0.1103515625, 0.107421875,
		0.107421875, 0.1083984375, 0.0986328125, 0.0791015625, 0.0146484375, 0.005859375,
		-0.009765625, -0.0400390625, 0.0078125, 0.0693359375, 0.015625, 0.044921875,
		0.0791015625, 0.0830078125, 0.076171875, 0.083984375, 0.0791015625, 0.1044921875,
		0.11328125, 0.0517578125, 0.060546875, 0.087890625, 0.03515625, 0.064453125,
		0.078125, 0.0634765625, 0.068359375, 0.083984375, 0.0908203125, 0.091796875,
		0.001953125, -0.0537109375, -0.0537109375, -0.0126953125, -0.0244140625,
		-0.00390625, 0.048828125, 0.080078125, 0.0908203125, 0.0947265625, 0.0087890625,
		-0.0439453125, -0.048828125, -0.0078125, -0.0205078125, 0.0029296875, 0.0556640625,
		0.0849609375, 0.09765625, 0.1015625, 0.01171875, -0.037109375, -0.041015625,
		-0.0048828125, -0.0185546875, 0.00390625, 0.0576171875, 0.0888671875, 0.1005859375,
		0.1005859375, 0.0126953125, -0.033203125, -0.03515625, -0.0146484375, -0.033203125,
		-0.0126953125, 0.001953125, 0.0068359375, 0.048828125, 0.021484375, 0.0673828125,
		0.0859375, 0.072265625, 0.083984375, 0.09765625, 0.10546875, 0.1044921875,
		0.013671875, -0.041015625, -0.04296875, -0.0029296875, -0.0166015625, 0.0048828125,
		0.05859375, 0.0869140625, 0.09765625, 0.1015625, 0.0126953125, -0.0380859375,
		-0.04296875, -0.0048828125, -0.0166015625, 0.0068359375, 0.060546875, 0.087890625,
		0.099609375, 0.103515625, 0.0146484375, -0.033203125, -0.037109375, -0.001953125,
		-0.015625, 0.0068359375, 0.0595703125, 0.091796875, 0.1025390625, 0.103515625,
		0.013671875, -0.0322265625, -0.033203125, -0.0146484375, -0.0341796875, -0.0126953125,
		0.001953125, 0.0068359375, 0.0498046875, 0.0224609375, 0.068359375, 0.0859375,
		0.0732421875, 0.0810546875, 0.09765625, 0.1044921875, 0.103515625, 0.013671875,
		-0.0419921875, -0.04296875, -0.0029296875, -0.013671875, 0.0068359375, 0.0595703125,
		0.087890625, 0.099609375, 0.103515625, 0.0146484375, -0.037109375, -0.041015625,
		-0.0009765625, -0.013671875, 0.009765625, 0.0625, 0.08984375, 0.1015625, 0.10546875,
		0.015625, -0.033203125, -0.0380859375, -0.0009765625, -0.0126953125, 0.009765625,
		0.064453125, 0.091796875, 0.1044921875, 0.1064453125, 0.0166015625, -0.0263671875,
		-0.029296875, -0.01171875, -0.033203125, -0.01171875, 0.0029296875, 0.0078125,
		0.048828125, 0.0234375, 0.0703125, 0.087890625, 0.0732421875, 0.0830078125, 0.09765625,
		0.10546875, 0.10546875, 0.013671875, -0.0380859375, -0.0390625, -0.001953125, -0.015625,
		0.00390625, 0.0546875, 0.0869140625, 0.0986328125, 0.1025390625, 0.0126953125,
		-0.041015625, -0.04296875, -0.001953125, -0.0126953125, 0.0068359375, 0.05859375,
		0.0908203125, 0.1025390625, 0.10546875, 0.0166015625, -0.03515625, -0.041015625,
		0.0, -0.0126953125, 0.009765625, 0.0625, 0.091796875, 0.1015625, 0.1044921875,
		0.0146484375, -0.0302734375, -0.033203125, -0.0126953125, -0.013671875, -0.0244140625,
		-0.0322265625, 0.037109375, 0.0859375, 0.103515625, 0.0986328125, 0.11328125,
		0.052734375, 0.01953125, -0.001953125, -0.0341796875, -0.017578125, -0.001953125,
		-0.052734375, 0.0205078125, 0.0, 0.0263671875, 0.052734375, 0.0869140625, 0.0859375,
		0.076171875, 0.09765625, 0.05078125, 0.0556640625, 0.0703125, 0.076171875,
		0.0556640625, 0.0126953125, 0.0185546875, 0.06640625, -0.0087890625, -0.005859375,
		0.001953125, -0.0087890625, 0.021484375, 0.060546875, 0.052734375, 0.0205078125,
		0.025390625, 0.0693359375, -0.005859375, -0.00390625, 0.00390625, -0.0078125,
		0.0263671875, 0.0634765625, 0.0517578125, 0.0166015625, 0.025390625, 0.072265625,
		-0.001953125, 0.0009765625, 0.0068359375, -0.00390625, 0.0263671875}

	data, err := hex.DecodeString(dataStr)
	if err != nil {
		t.Fatal("Failed to decode input hex string")
	}
	actual := adc.ProcessTraceData(data)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual processed data did not match expected")
	}
}
