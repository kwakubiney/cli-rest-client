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
	MapData      map[string]string
	FieldKeys    []string
	Where        string
	By           string
	Flag         *flag.FlagSet
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

func ResponseByStatusCode(resp *http.Response, clientHandler *CliHandler) error{
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
	}	
return nil
}

func DispatchHelper(s *CliHandler) error{
	if s.Options.TypeOfObject == "" {
		return errors.New("unrecognizable command")
	}
		switch s.Options.Method {
		case "create":
			if s.Options.TypeOfObject == "" || s.Options.Fields == "" {
				return errors.New("unrecognizable command")
			}
			createFieldKeys, mapofData := command.ParseFields(s.Options.Fields)
			s.Options.FieldKeys, s.Options.MapData = createFieldKeys, mapofData
			log.Println(mapofData)
			if !command.ValidateFields(s.Options.Method, createFieldKeys, s.Options.TypeOfObject) {
				return errors.New("unrecognizable command")
			}
		case "update":
			if s.Options.TypeOfObject == "" || s.Options.Fields == "" || s.Options.Where == "" {
				return errors.New("unrecognizable command")
			}
			updateFieldKeys, mapOfData := command.ParseFields(s.Options.Fields)
			if !command.ValidateFields(s.Options.Method, updateFieldKeys, s.Options.TypeOfObject) {
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
	return nil
	}

func (s *CliHandler) Dispatch() error {
	return DispatchHelper(s)
}

func ApiRequestDispatcher(clientHandler *CliHandler) error {
	requestBody := clientHandler.Options.MapData
	if clientHandler.Options.TypeOfObject == "user"{
	if clientHandler.Options.Method == "create" {
		fmt.Println("Creating user...........")
		fmt.Println("=> POST https://localhost/user\n" +
			"<=")
		resp, err := utils.MakeRequest(fmt.Sprintf("http://127.0.0.1:%s/User", os.Getenv("PORT")), requestBody, "POST")
		if err != nil {
			return err
		}

		ResponseByStatusCode(resp, clientHandler)
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
		ResponseByStatusCode(resp, clientHandler)
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
		ResponseByStatusCode(resp, clientHandler)
		return nil
	}
	return nil
}
if clientHandler.Options.TypeOfObject == "game"{
	if clientHandler.Options.Method == "create" {
		fmt.Println("Creating game...........")
		fmt.Println("=> POST https://localhost/Game\n" +
			"<=")
		resp, err := utils.MakeRequest(fmt.Sprintf("http://127.0.0.1:%s/Game", os.Getenv("PORT")), requestBody, "POST")
		if err != nil {
			return err
		}

		ResponseByStatusCode(resp, clientHandler)
	}

	if clientHandler.Options.Method == "update" {
		fmt.Println("Updating game...........")
		fmt.Println("=> PUT https://localhost/game?title=<title>\n" +
			"<=")
		//TODO: Allow updates based on other fields
		resp, err := utils.MakeRequest(fmt.Sprintf("http://127.0.0.1:%s/Game?title=%s", os.Getenv("PORT"), clientHandler.Options.Where), 
		requestBody, "PUT")
		if err != nil {
			return err
		}
		ResponseByStatusCode(resp, clientHandler)
		return nil
	}

	if clientHandler.Options.Method == "filter" {
		log.Println(clientHandler.Options.By, clientHandler.Options.Where)
		resp, err := utils.MakeRequest(fmt.Sprintf("http://127.0.0.1:%s/Game?%s=%s", os.Getenv("PORT"),
			clientHandler.Options.By, clientHandler.Options.Where), map[string]interface{}{
				"by": clientHandler.Options.By,
			}, "GET")

		if err != nil {
			return err
		}
		ResponseByStatusCode(resp, clientHandler)
		return nil
	}
	return nil
}
return nil
}


