package abr

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseURL = "https://abr.business.gov.au/ABRXMLSearch/AbrXmlSearch.asmx"

type ABRXMLSearchClient struct {
	GUID int
}

func (c *ABRXMLSearchClient) SearchByABN(abn string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	resource := "SearchByABNv202001"

	// build params
	params := url.Values{}
	params.Add("searchString", abn)
	params.Add("includeHistoricalDetails", "Y")
	params.Add("authenticationGuid", fmt.Sprintf("%d", c.GUID))
	values := params.Encode()

	// join resource and encoded values
	reqURL := base.JoinPath(resource)
	reqURL.RawQuery = values

	resp, err := http.Get(reqURL.String())
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
