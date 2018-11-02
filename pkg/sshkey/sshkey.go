// Go project by arthur
// blihUI
// 2018

package sshkey

import (
	"blihUI/pkg/blih"
	"fmt"
)

type SSHKey struct {
	name, key string
}

func (s SSHKey) String() string {
	return fmt.Sprintf("%s %s", s.key, s.name)
}

func Delete(name string, b *blih.BLIH) error {
	_, err := b.Request("sshkey/"+name, "DELETE", nil)
	return err
}

func List(b *blih.BLIH) ([]SSHKey, error) {
	list, err := b.Request("sshkeys", "GET", nil)
	if err != nil {
		return nil, err
	}
	var keys []SSHKey
	for key, value := range list {
		keys = append(keys, SSHKey{name: key, key: value.(string)})
	}
	return keys, nil
}
