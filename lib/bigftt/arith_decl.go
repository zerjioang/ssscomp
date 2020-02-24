// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bigfft

import "math/big"

// implemented in arith_$GOARCH.s
func addVV(z, x, y []big.Word) (c big.Word)
func subVV(z, x, y []big.Word) (c big.Word)
func addVW(z, x []big.Word, y big.Word) (c big.Word)
func subVW(z, x []big.Word, y big.Word) (c big.Word)
func shlVU(z, x []big.Word, s uint) (c big.Word)
func mulAddVWW(z, x []big.Word, y, r big.Word) (c big.Word)
func addMulVVW(z, x []big.Word, y big.Word) (c big.Word)
