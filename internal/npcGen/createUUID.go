package npcgen

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// Appropriated from github.com/google/uuid
func encodeHex(dst []byte, uuid [16]byte) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

// Appropriated from github.com/google/uuid
func CreateUUIDv4() (string, error) {
	uuid := [16]byte{}
	rand_reader := rand.Reader

	// Reads in uuid len amount of bytes of random numbers
	_, err := io.ReadFull(rand_reader, uuid[:])
	if err != nil {
		return "", err
	}

	// Changes the bits specified for the Version & Variant
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x3f // Variant 10

	// Creates a buffer and takes the bytes from the uuid and
	// encodes them into hexadecimal, replacing the appropriate
	// char with "-" to output the uuid as a string
	var buf [36]byte
	encodeHex(buf[:], uuid)

	return string(buf[:]), nil
}
