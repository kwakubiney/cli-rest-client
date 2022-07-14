package cli

import (
	"errors"
	"flag"
	"log"

	"github.com/kwakubiney/canonical-take-home/command"
)

type Options struct {
	Method       string
	TypeOfObject string
	Fields       string
	Help         *bool
	MapData      map[string]string
	FieldKeys    []string
}

type CliHandler struct {
	Options *Options
}

func NewCliHandler(opts *Options) *CliHandler {
	return &CliHandler{
		Options: opts,
	}
}

func (s *CliHandler) Dispatch() error{
	if *s.Options.Help || s.Options.Method == "" {
		return errors.New("unrecognizable command. check --help")
	}
	if s.Options.TypeOfObject == "user" {
		switch s.Options.Method {
		case "create":
			if s.Options.TypeOfObject == "" || s.Options.Fields == "" {
				return errors.New("unrecognizable command. check --help")
			}
			createFieldKeys, mapofData := command.ParseFields(s.Options.Fields)
			s.Options.FieldKeys, s.Options.MapData = createFieldKeys, mapofData
			if !command.ValidateCreateandUpdateUserFields(s.Options.Method, createFieldKeys) {
				return errors.New("urecognizable command. check --help")
			}
		case "update":
			{
				if s.Options.TypeOfObject == "" || s.Options.Fields == "" {
					return errors.New("unrecognizable command. check --help")
				}
				updateFieldKeys, mapOfData := command.ParseFields(s.Options.Fields)
				if !command.ValidateCreateandUpdateUserFields(s.Options.Method, updateFieldKeys) {
					return errors.New("unrecognizable command. check --help")
				}
				log.Println(updateFieldKeys, mapOfData)
				s.Options.FieldKeys, s.Options.MapData = updateFieldKeys, mapOfData
			}
		case "delete":
			{
				if s.Options.TypeOfObject == "" || s.Options.Fields == "" {
					flag.Usage()
					return errors.New("unrecognizable command. check --help")
				}
				deleteFieldKeys, mapofData := command.ParseFields(s.Options.Fields)
				if !command.ValidateCreateandUpdateUserFields(s.Options.Method, deleteFieldKeys) {
					return errors.New("unrecognizable command. check --help")
				}
				s.Options.FieldKeys, s.Options.MapData = deleteFieldKeys, mapofData
			}

		default:
			return errors.New("unrecognizable command. Check --help")
		}
	}
	return nil
}
