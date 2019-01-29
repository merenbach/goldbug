// Copyright 2018 Andrew Merenbach
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

const defaultAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type cipherRequest struct {
	Alphabet string `json:"alphabet"`
	Message  string `json:"message" binding:"required"`
	Strict   bool   `json:"strict"`
}

type atbashCipherRequest struct {
	cipherRequest
}

type rot13CipherRequest struct {
	cipherRequest
}

type caesarCipherRequest struct {
	cipherRequest
	Shift int `json:"shift" binding:"required"`
}

type decimationCipherRequest struct {
	cipherRequest
	Multiplier int `json:"multiplier" binding:"required"`
}

type affineCipherRequest struct {
	cipherRequest
	Shift      int `json:"shift" binding:"required"`
	Multiplier int `json:"multiplier" binding:"required"`
}

type keywordCipherRequest struct {
	cipherRequest
	Keyword string `json:"keyword" binding:"required"`
}

type vigenereFamilyCipherRequest struct {
	cipherRequest
	Countersign string `json:"countersign" binding:"required"`
}

type vigenereCipherRequest struct {
	vigenereFamilyCipherRequest
	TextAutoclave bool `json:"textAutoclave"`
	KeyAutoclave  bool `json:"keyAutoclave"`
}

type trithemiusCipherRequest struct {
	cipherRequest
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	v1 := router.Group("/v1")
	{
		v1.POST("/encipher/rot13", func(c *gin.Context) {
			var r rot13CipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewRot13Cipher(r.Alphabet).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/rot13", func(c *gin.Context) {
			var r rot13CipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewRot13Cipher(r.Alphabet).Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/atbash", func(c *gin.Context) {
			var r atbashCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewAtbashCipher(r.Alphabet).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/atbash", func(c *gin.Context) {
			var r atbashCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewAtbashCipher(r.Alphabet).Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/caesar", func(c *gin.Context) {
			var r caesarCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewCaesarCipher(r.Alphabet, r.Shift).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/caesar", func(c *gin.Context) {
			var r caesarCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewCaesarCipher(r.Alphabet, r.Shift).Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/decimation", func(c *gin.Context) {
			var r decimationCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewDecimationCipher(r.Alphabet, r.Multiplier).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/decimation", func(c *gin.Context) {
			var r decimationCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewDecimationCipher(r.Alphabet, r.Multiplier).Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/affine", func(c *gin.Context) {
			var r affineCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewAffineCipher(r.Alphabet, r.Multiplier, r.Shift).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/affine", func(c *gin.Context) {
			var r affineCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewAffineCipher(r.Alphabet, r.Multiplier, r.Shift).Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/keyword", func(c *gin.Context) {
			var r keywordCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewKeywordCipher(r.Alphabet, r.Keyword).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/keyword", func(c *gin.Context) {
			var r keywordCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewKeywordCipher(r.Alphabet, r.Keyword).Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/vigenere", func(c *gin.Context) {
			var r vigenereCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			var cipher *VigenereFamilyCipher
			if !r.TextAutoclave && !r.KeyAutoclave {
				cipher = NewVigenereCipher(r.Countersign, r.Alphabet)
			} else if r.TextAutoclave && !r.KeyAutoclave {
				cipher = NewVigenereTextAutoclaveCipher(r.Countersign, r.Alphabet)
			} else if !r.TextAutoclave && r.KeyAutoclave {
				cipher = NewVigenereKeyAutoclaveCipher(r.Countersign, r.Alphabet)
			} else if r.TextAutoclave && r.KeyAutoclave {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Text autoclave and key autoclave cannot be specified simultaneously.",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": cipher.Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/vigenere", func(c *gin.Context) {
			var r vigenereCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			var cipher *VigenereFamilyCipher
			if !r.TextAutoclave && !r.KeyAutoclave {
				cipher = NewVigenereCipher(r.Countersign, r.Alphabet)
			} else if r.TextAutoclave && !r.KeyAutoclave {
				cipher = NewVigenereTextAutoclaveCipher(r.Countersign, r.Alphabet)
			} else if !r.TextAutoclave && r.KeyAutoclave {
				cipher = NewVigenereKeyAutoclaveCipher(r.Countersign, r.Alphabet)
			} else if r.TextAutoclave && r.KeyAutoclave {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Text autoclave and key autoclave cannot be specified simultaneously.",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": cipher.Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/beaufort", func(c *gin.Context) {
			var r vigenereFamilyCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewBeaufortCipher(r.Countersign, r.Alphabet).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/beaufort", func(c *gin.Context) {
			var r vigenereFamilyCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewBeaufortCipher(r.Countersign, r.Alphabet).Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/gronsfeld", func(c *gin.Context) {
			var r vigenereFamilyCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewGronsfeldCipher(r.Countersign, r.Alphabet).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/gronsfeld", func(c *gin.Context) {
			var r vigenereFamilyCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewGronsfeldCipher(r.Countersign, r.Alphabet).Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/variantbeaufort", func(c *gin.Context) {
			var r vigenereFamilyCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewVariantBeaufortCipher(r.Countersign, r.Alphabet).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/variantbeaufort", func(c *gin.Context) {
			var r vigenereFamilyCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewVariantBeaufortCipher(r.Countersign, r.Alphabet).Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/dellaporta", func(c *gin.Context) {
			var r vigenereFamilyCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewDellaPortaCipher(r.Countersign, r.Alphabet).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/dellaporta", func(c *gin.Context) {
			var r vigenereFamilyCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewDellaPortaCipher(r.Countersign, r.Alphabet).Decipher(r.Message, r.Strict),
			})
		})

		v1.POST("/encipher/trithemius", func(c *gin.Context) {
			var r trithemiusCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewTrithemiusCipher(r.Alphabet).Encipher(r.Message, r.Strict),
			})
		})
		v1.POST("/decipher/trithemius", func(c *gin.Context) {
			var r trithemiusCipherRequest
			if err := c.BindJSON(&r); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if r.Alphabet == "" {
				r.Alphabet = defaultAlphabet
			}
			c.JSON(http.StatusOK, gin.H{
				"message": NewTrithemiusCipher(r.Alphabet).Decipher(r.Message, r.Strict),
			})
		})
	}

	router.Run(":" + port)
}
