/*
 * Copyright (C) 2015-2018 Virgil Security Inc.
 *
 * Lead Maintainer: Virgil Security Inc. <support@virgilsecurity.com>
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *   (1) Redistributions of source code must retain the above copyright
 *   notice, this list of conditions and the following disclaimer.
 *
 *   (2) Redistributions in binary form must reproduce the above copyright
 *   notice, this list of conditions and the following disclaimer in
 *   the documentation and/or other materials provided with the
 *   distribution.
 *
 *   (3) Neither the name of the copyright holder nor the names of its
 *   contributors may be used to endorse or promote products derived
 *   from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR ''AS IS'' AND ANY EXPRESS OR
 * IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT,
 * INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
 * IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 */

package sdk

import (
	"encoding/hex"
	"testing"
	"time"

	"sync"

	"fmt"

	"github.com/stretchr/testify/assert"
	"gopkg.in/virgil.v5/cryptoimpl"
)

func TestCachingJwtProvider(t *testing.T) {

	crypto := cryptoimpl.NewVirgilCrypto()

	pk, err := crypto.GenerateKeypair()
	if err != nil {
		panic(err)
	}

	signer := cryptoimpl.NewVirgilAccessTokenSigner()

	genCount := 0

	prov := NewCachingJwtProvider(func(context *TokenContext) (*Jwt, error) {

		genCount++

		issuedAt := time.Now().UTC().Truncate(time.Second)
		expiresAt := issuedAt.Add(time.Second * 6)

		jwtBody, err := NewJwtBodyContent("app_id", "identity", issuedAt, expiresAt, nil)
		if err != nil {
			return nil, err
		}

		jwtHeader, err := NewJwtHeaderContent(signer.GetAlgorithm(), hex.EncodeToString(pk.PrivateKey().Identifier()))

		if err != nil {
			return nil, err
		}

		unsignedJwt, err := NewJwt(jwtHeader, jwtBody, nil)
		if err != nil {
			return nil, err
		}
		jwtSignature, err := signer.GenerateTokenSignature(unsignedJwt.Unsigned(), pk.PrivateKey())
		if err != nil {
			return nil, err
		}

		return NewJwt(jwtHeader, jwtBody, jwtSignature)

	})

	routines := 100

	wg := &sync.WaitGroup{}
	wg.Add(routines)

	start := time.Now()

	total := 0
	for i := 0; i < routines; i++ {

		go func() {

			for time.Now().Sub(start) < (time.Second * 5) {
				token, err := prov.GetToken(&TokenContext{Identity: "Alice"})
				assert.NotNil(t, token)
				assert.NoError(t, err)
				total++
			}

			wg.Done()
		}()

	}
	wg.Wait()
	assert.Equal(t, 6, genCount)
	fmt.Println("total", total)

}
