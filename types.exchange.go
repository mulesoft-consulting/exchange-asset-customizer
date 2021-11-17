package main

type anypointContextKey string

func (c anypointContextKey) String() string {
	return "anypoint " + string(c)
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	RedirectUrl string `json:"redirectUrl"`
}

type CustomFieldKeyBody struct {
	DataType    string `json:"dataType"`
	DisplayName string `json:"displayName"`
	TagKey      string `json:"tagKey"`
}

type CustomFieldValueBody struct {
	TagValue string `json:"tagValue"`
}

type CustomCategoryValueBody struct {
	TagValue []string `json:"tagValue"`
}

type PatchAttributeBody struct {
	ContactName  string `json:"contactName,omitempty"`
	ContactEmail string `json:"contactEmail,omitempty"`
}

func NewPatchAttributeBody(key string, value string) *PatchAttributeBody {
	if key == "contactName" {
		return &PatchAttributeBody{
			ContactName: value,
		}
	} else if key == "contactEmail" {
		return &PatchAttributeBody{
			ContactEmail: value,
		}
	}
	return nil
}
