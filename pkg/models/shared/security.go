// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Security struct {
	OauthClientCredentials *string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *Security) GetOauthClientCredentials() *string {
	if o == nil {
		return nil
	}
	return o.OauthClientCredentials
}
