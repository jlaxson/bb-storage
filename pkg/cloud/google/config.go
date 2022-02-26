package google

import (
	"net/http"

	bb_http "github.com/buildbarn/bb-storage/pkg/http"
	google_pb "github.com/buildbarn/bb-storage/pkg/proto/configuration/cloud/google"
	"github.com/buildbarn/bb-storage/pkg/util"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func OptionsFromConfiguration(configuration *google_pb.GoogleClientConfiguration, name string) ([]option.ClientOption, error) {
	options := []option.ClientOption{}

	roundTripper, err := bb_http.NewRoundTripperFromConfiguration(configuration.GetHttpClient())
	if err != nil {
		return options, util.StatusWrap(err, "Failed to create HTTP client")
	}
	options = append(options, option.WithHTTPClient(&http.Client{
		Transport: bb_http.NewMetricsRoundTripper(roundTripper, name),
	}))

	if credentialsOptions := configuration.GetCredentials(); credentialsOptions != nil {
		switch credentialsType := credentialsOptions.(type) {
		case *google_pb.GoogleClientConfiguration_CredentialsJson:
			// Use static credentials.
			staticCredentials := credentialsType.CredentialsJson
			options = append(options, option.WithCredentialsJSON([]byte(staticCredentials)))
		default:
			return options, status.Error(codes.InvalidArgument, "Unknown credentials options type provided")
		}
	}

	if scopes := configuration.GetScopes(); len(scopes) > 0 {
		options = append(options, option.WithScopes(scopes...))
	}
	return options, nil
}
