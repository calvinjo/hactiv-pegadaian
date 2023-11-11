package config

const (
	Success             = "200"
	InternalServerError = "500"
	InvalidFieldFormat  = "400"
	Failed              = "412"
	DataNotFound        = "400"
	Exceed              = "429"
	Invalid             = "422"

	MaxLimit = 10000000
)

var listResponse = map[string]string{
	Success:             "Success",
	InternalServerError: "Internal Server Error",
	InvalidFieldFormat:  "Invalid Field Format",
	Failed:              "Failed",
}

func MessageResponse(responseCode string) string {
	return listResponse[responseCode]
}
