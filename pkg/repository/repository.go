// Go project by arthur
// blihUI
// 2018

package repository

import (
	"blihUI/pkg/blih"
	"blihUI/pkg/data"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type repositoryInformation string

func (s repositoryInformation) String() string {
	if s == "" {
		return "(unknown)"
	}
	return string(s)
}

type repositoryDescription string

func (s repositoryDescription) String() string {
	if s == "" {
		return "(none)"
	}
	return string(s)
}

type Repository struct {
	name        string
	uuid, url   repositoryInformation
	description repositoryDescription
	public      bool
	creation    time.Time
	acl         map[string]string
}

func (r *Repository) Name() string {
	return string(r.name)
}

func (r *Repository) UUID() string {
	return string(r.uuid)
}

func (r *Repository) URL() string {
	return string(r.url)
}

func (r *Repository) Description() string {
	return string(r.description)
}

func (r *Repository) Public() bool {
	return r.public
}

func (r *Repository) Creation() time.Time {
	return r.creation
}

func (r *Repository) ACL() map[string]string {
	return r.acl
}

func (r Repository) String() string {
	return fmt.Sprintf("%s (UUID: %s, description: %s, public: %v, url: %s, creation: %s)",
		r.name, r.uuid, r.description, r.public, r.url, r.creation)
}

func List(b *blih.BLIH) ([]Repository, error) {
	repositories, err := b.Request("repositories", "GET", nil)
	if err != nil {
		return nil, err
	}
	list, ok := repositories["repositories"].(map[string]interface{})
	if ok != true {
		err = errors.New("could not convert repositories to map[string]interface{}")
		return nil, err
	}
	var repos []Repository
	for key := range list {
		repo := Repository{name: key}
		repos = append(repos, repo)
	}
	return repos, nil
}

func Create(name, description string, b *blih.BLIH) error {
	d := data.Data{"name": name, "type": "git"}
	if description != "" {
		d["description"] = description
	}
	_, err := b.Request("repositories", "POST", &d)
	return err
}

func Delete(name string, b *blih.BLIH) error {
	_, err := b.Request("repository/"+name, "DELETE", nil)
	return err
}

func Info(name string, b *blih.BLIH) (*Repository, error) {
	repository, err := b.Request("repository/"+name, "GET", nil)
	if err != nil {
		return nil, err
	}
	infos, ok := repository["message"].(map[string]interface{})
	if ok != true {
		err := errors.New(fmt.Sprintf("could not convert infos of %s to map[string]interface{}", name))
		return nil, err
	}
	repo := &Repository{
		name:        name,
		uuid:        repositoryInformation(infos["uuid"].(string)),
		description: repositoryDescription(infos["description"].(string)),
		url:         repositoryInformation(infos["url"].(string)),
	}
	if repo.description == "None" {
		repo.description = ""
	}
	repo.public, err = strconv.ParseBool(infos["public"].(string))
	if err != nil {
		err = errors.New(fmt.Sprintf("could not convert infos[\"public\"] of %s to bool: %s", name, err))
		return nil, err
	}
	timestamp, err := strconv.ParseInt(infos["creation_time"].(string), 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("could not convert infos[\"creation_time\"] of %s to int: %s", name, err))
		return nil, err
	}
	repo.creation = time.Unix(timestamp, 0)
	return repo, err
}

func SetACL(name, acluser, acl string, b *blih.BLIH) error {
	d := data.Data{"user": acluser, "acl": acl}
	_, err := b.Request("repository/"+name+"/acls", "POST", &d)
	return err
}

func GetACL(name string, b *blih.BLIH) (map[string]string, error) {
	repository, err := b.Request("repository/"+name+"/acls", "GET", nil)
	if err != nil {
		return nil, err
	}
	acls := make(map[string]string)
	for key, value := range repository {
		acls[key] = value.(string)
	}
	return acls, err
}
