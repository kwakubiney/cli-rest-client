package cli

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"io"

	"github.com/kwakubiney/canonical-take-home/command"
	"net/http"
	"github.com/kwakubiney/canonical-take-home/internal/utils"
)


type Options struct {
	Method       string
	TypeOfObject string
	Fields       string
	Help         *bool
	MapData      map[string]string
	FieldKeys    []string
	Where        string
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
		return errors.New("unrecognizable command")
	}
	
	if s.Options.TypeOfObject == "user" {
		switch s.Options.Method {
		case "create":
			if s.Options.TypeOfObject == "" || s.Options.Fields == "" {
				return errors.New("unrecognizable command")
			}
			createFieldKeys, mapofData := command.ParseFields(s.Options.Fields)
			s.Options.FieldKeys, s.Options.MapData = createFieldKeys, mapofData
			if !command.ValidateCreateandUpdateUserFields(s.Options.Method, createFieldKeys) {
				return errors.New("unrecognizable command")
			}
		case "update":
			{
			if s.Options.TypeOfObject == "" || s.Options.Fields == "" || s.Options.Where == ""{
				return errors.New("unrecognizable command")
			}
			updateFieldKeys, mapOfData := command.ParseFields(s.Options.Fields)
			if !command.ValidateCreateandUpdateUserFields(s.Options.Method, updateFieldKeys) {
				return errors.New("unrecognizable command")
			}
			log.Println(updateFieldKeys, mapOfData)
			s.Options.FieldKeys, s.Options.MapData = updateFieldKeys, mapOfData
			}

		case "delete":
			{
			if s.Options.TypeOfObject == "" || s.Options.Fields == "" {
				flag.Usage()
				return errors.New("unrecognizable command")
				}
			deleteFieldKeys, mapofData := command.ParseFields(s.Options.Fields)
			if !command.ValidateCreateandUpdateUserFields(s.Options.Method, deleteFieldKeys) {
				return errors.New("unrecognizable command")
				}
			s.Options.FieldKeys, s.Options.MapData = deleteFieldKeys, mapofData
			}

		default:
			return errors.New("unrecognizable command")
		}
	}
	return nil
}


func ApiRequestDispatcher(clientHandler *CliHandler) error{
	requestBody := clientHandler.Options.MapData
	if clientHandler.Options.Method == "create"{
		resp, err := utils.MakeRequest(fmt.Sprintf("http://127.0.0.1:%s/User" ,os.Getenv("PORT")), requestBody, "POST")
		if err != nil{
			return err	
		}
		if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
	return nil
	}

	if clientHandler.Options.Method == "update"{
		resp, err := utils.MakeRequest(fmt.Sprintf("http://127.0.0.1:%s/User" ,os.Getenv("PORT")), requestBody, "PUT")
		if err != nil{
			return err	
		}
		if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
	return nil
	}

	return nil
}