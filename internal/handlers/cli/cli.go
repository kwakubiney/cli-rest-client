package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"bytes"
	"encoding/json"
	"net/http"

	"github.com/kwakubiney/canonical-take-home/command"
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
	By           string
	Flag         flag.FlagSet
}

type CliHandler struct {
	Options *Options
}

func NewCliHandler(opts *Options) *CliHandler {
	return &CliHandler{
		Options: opts,
	}
}

func PrettyPrint(response []byte) error {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, response, "", "\t")
	if err != nil {
		log.Println("JSON parse error: ", err)
		return err
	}
	fmt.Println(prettyJSON.String())
	return nil
}

func (s *CliHandler) Dispatch() error {
	if s.Options.TypeOfObject == "" {
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
			if s.Options.TypeOfObject == "" || s.Options.Fields == "" || s.Options.Where == "" {
				return errors.New("unrecognizable command")
			}
			updateFieldKeys, mapOfData := command.ParseFields(s.Options.Fields)
			if !command.ValidateCreateandUpdateUserFields(s.Options.Method, updateFieldKeys) {
				return errors.New("unrecognizable command")
			}

			s.Options.FieldKeys, s.Options.MapData = updateFieldKeys, mapOfData
	

		//handle validation on db level
		case "filter":
			if s.Options.TypeOfObject == "" || s.Options.Where == "" || s.Options.By == "" {
				return errors.New("unrecognizable command")
			}
			mapOfFields := make(map[string]string)
			s.Options.MapData = mapOfFields
			s.Options.MapData[s.Options.By] = s.Options.Where

		default:
			return errors.New("unrecognizable command")
		}
	}
	return nil
}

func ApiRequestDispatcher(clientHandler *CliHandler) error {
	requestBody := clientHandler.Options.MapData
	if clientHandler.Options.Method == "create" {
		fmt.Println("Creating user...........")
		fmt.Println("=> POST https://localhost/user\n" +
			"<=")
		resp, err := utils.MakeRequest(fmt.Sprintf("http://127.0.0.1:%s/User", os.Getenv("PORT")), requestBody, "POST")
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusCreated {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Println("<=== Resp Status: 201")
			fmt.Println("===> Response:")
			PrettyPrint(bodyBytes)
		} else {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Printf("===> Response:%d\n", resp.StatusCode)
			PrettyPrint(bodyBytes)
			clientHandler.Options.Flag.Usage()
		}

	}

	if clientHandler.Options.Method == "update" {
		fmt.Println("Updating user...........")
		fmt.Println("=> PUT https://localhost/user?username=<name>\n" +
			"<=")
		//TODO: Allow updates based on other fields
		resp, err := utils.MakeRequest(fmt.Sprintf("http://127.0.0.1:%s/User?username=%s", os.Getenv("PORT"), clientHandler.Options.Where), 
		requestBody, "PUT")
		if err != nil {
			return err
		}
		if resp.StatusCode == http.StatusOK {
			fmt.Println("<=== Resp Status: 200 OK")
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Println("===> Response:")
			PrettyPrint(bodyBytes)
		} else {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Printf("===> Response:%d\n", resp.StatusCode)
			PrettyPrint(bodyBytes)
			clientHandler.Options.Flag.Usage()
		}
		return nil
	}

	if clientHandler.Options.Method == "filter" {
		resp, err := utils.MakeRequest(fmt.Sprintf("http://127.0.0.1:%s/User?%s=%s", os.Getenv("PORT"),
			clientHandler.Options.By, clientHandler.Options.Where), map[string]interface{}{
				"by": clientHandler.Options.By,
			}, "GET")

		if err != nil {
			return err
		}
		if resp.StatusCode == http.StatusOK {
			fmt.Println("<=== Rsp Status: 200 OK")
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			fmt.Println("===> Response:")
			PrettyPrint(bodyBytes)

			return nil
		} else {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Printf("===> Response:%d\n", resp.StatusCode)
			PrettyPrint(bodyBytes)
			clientHandler.Options.Flag.Usage()
		}
		return nil
	}
	return nil
}
