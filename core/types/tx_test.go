// Copyright(c) 2018 DSiSc Group. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"github.com/DSiSc/craft/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var emptyTx *types.Transaction

func TestNewTransaction(t *testing.T) {
	assert := assert.New(t)
	b := Address{
		0xb2, 0x6f, 0x2b, 0x34, 0x2a, 0xab, 0x24, 0xbc, 0xf6, 0x3e,
		0xa2, 0x18, 0xc6, 0xa9, 0x27, 0x4d, 0x30, 0xab, 0x9a, 0x15,
	}
	emptyTx = NewTransaction(
		0,
		&b,
		big.NewInt(0), 0, big.NewInt(0),
		nil,
		b,
	)
	notoTx := NewTransaction(
		0,
		nil,
		nil, 0, big.NewInt(0),
		nil,
		b,
	)
	assert.NotNil(emptyTx)
	assert.NotNil(notoTx)
}

func TestTxHash(t *testing.T) {
	assert := assert.New(t)
	hash := TxHash(emptyTx)
	expect := types.Hash{0xf3, 0x2c, 0x26, 0xa4, 0xee, 0x93, 0x3d, 0x72, 0x80, 0x40, 0xa5, 0xb1, 0x1b, 0x8d, 0xd3, 0x94,
		0x31, 0x83, 0xec, 0x50, 0x36, 0xfd, 0xac, 0xf9, 0x35, 0x2, 0x1, 0x1a, 0xab, 0x95, 0xb8, 0xb5}
	assert.Equal(expect, hash)
}
