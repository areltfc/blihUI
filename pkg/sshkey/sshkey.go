// Go project by arthur
// blihUI
// 2018

package sshkey

import "blihUI/pkg/blih"

type SSHKey struct {
	name, key string
}

func Delete(name string, b *blih.BLIH) error {
	_, err := b.Request("sshkey/"+name, "DELETE", nil)
	return err
}
