package function

import (
	"encoding/json"
	"errors"
	"fmt"
	"function/pkg"
	"io"
	"net/http"

	sdk "github.com/AbdulahadAbduqahhorov/ucode-sdk"
)

/*
Answer below questions before starting the function.

When the function invoked?
  - table_slug -> AFTER | BEFORE | HTTP -> CREATE | UPDATE | MULTIPLE_UPDATE | DELETE | APPEND_MANY2MANY | DELETE_MANY2MANY

What does it do?
- Explain the purpose of the function.(O'zbekcha yozilsa ham bo'ladi.)
*/

func Handler(params *pkg.Params) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ucodeApi = sdk.New(&sdk.Config{
				BaseURL:      params.Config.BaseUrl,
				BotToken:     params.Config.BotToken,
				AccountIds:   []string{params.Config.ChatId},
				FunctionName: params.Config.Name,
			})

			request       = sdk.Request{}
			response      = sdk.Response{}
			errorResponse = sdk.ResponseError{}

			returnError = func(errorResponse sdk.ResponseError) {
				response = sdk.Response{
					Status: "error",
					Data:   map[string]any{"message": errorResponse.ClientErrorMessage, "error": errorResponse.ErrorMessage, "description": errorResponse.Description},
				}
				params.Log.Err(errors.New(errorResponse.ErrorMessage)).Msg(errorResponse.ClientErrorMessage)
				handleResponse(w, response, errorResponse.StatusCode)
			}
		)

		{
			requestByte, err := io.ReadAll(r.Body)
			if err != nil {
				errorResponse.ClientErrorMessage = "Error on getting request body"
				errorResponse.ErrorMessage = err.Error()
				errorResponse.StatusCode = http.StatusInternalServerError
				returnError(errorResponse)
				return
			}

			if err = json.Unmarshal(requestByte, &request); err != nil {
				errorResponse.ClientErrorMessage = "Error on unmarshal request"
				errorResponse.ErrorMessage = err.Error()
				errorResponse.StatusCode = http.StatusInternalServerError
				returnError(errorResponse)
				return
			}
			fmt.Println(string(requestByte))

			appId, ok := request.Data["app_id"].(string)
			if !ok {
				errorResponse.ClientErrorMessage = "App ID is required"
				errorResponse.ErrorMessage = "App ID is required"
				errorResponse.StatusCode = http.StatusUnauthorized
				returnError(errorResponse)
				return
			}
			ucodeApi.Cfg.SetAppId(appId)
		}

		response.Status = "done"
		handleResponse(w, response, 200)
	}
}

func handleResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	bodyByte, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`
			{
				"error": "Error marshalling response"
			}
		`))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(bodyByte)
}
