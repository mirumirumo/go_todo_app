package auth

import (
	"bytes"
	"testing"
)

func TestEmbed(t *testing.T) {
	want := []byte("ssh-rsa")
	if !bytes.Contains(rawPubKey, want) {
		t.Errorf("want %s, but got %s", want, rawPubKey)
	}

	want = []byte("-----BEGIN OPENSSH PRIVATE KEY-----")
	if !bytes.Contains(rawPrivKey, want) {
		t.Errorf("want %s, but got %s", want, rawPrivKey)
	}

}
