package genesisfile

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


// sometimes we define a convenience type in a library which breaks
// our code. The _right_ answer to this problem would be to recursively
// unwrap typedefs, trying each of them until we found one we could handle,
// a primitive, or a structdef
//
// we're not doing that
//
// the _expedient_ answer is to create a map of convenience typenames, and
// what we unpack them to

var typeExceptions map[string]string

func init() {
	typeExceptions = map[string]string{
		"eai.RateTable":  "[]eai.RTRow",
		"sv.EAIFeeTable": "[]sv.EAIFee",
	}
}
