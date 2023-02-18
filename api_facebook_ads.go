
/*
 * Mailchimp Marketing API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.0.55
 * Contact: apihelp@mailchimp.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package mailchimp

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type FacebookAdsApiService service

/*
FacebookAdsApiService List facebook ads
Get list of Facebook ads.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *FacebookAdsApiGetAllFacebookAdsOpts - Optional Parameters:
     * @param "Fields" (optional.Interface of []string) -  A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
     * @param "ExcludeFields" (optional.Interface of []string) -  A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
     * @param "Count" (optional.Int32) -  The number of records to return. Default value is 10. Maximum value is 1000
     * @param "Offset" (optional.Int32) -  Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0.
     * @param "SortField" (optional.String) -  Returns files sorted by the specified field.
     * @param "SortDir" (optional.String) -  Determines the order direction for sorted results.

@return InlineResponse2008
*/

type FacebookAdsApiGetAllFacebookAdsOpts struct { 
	Fields optional.Interface
	ExcludeFields optional.Interface
	Count optional.Int32
	Offset optional.Int32
	SortField optional.String
	SortDir optional.String
}

func (a *FacebookAdsApiService) GetAllFacebookAds(ctx context.Context, localVarOptionals *FacebookAdsApiGetAllFacebookAdsOpts) (InlineResponse2008, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse2008
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/facebook-ads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeFields.IsSet() {
		localVarQueryParams.Add("exclude_fields", parameterToString(localVarOptionals.ExcludeFields.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortField.IsSet() {
		localVarQueryParams.Add("sort_field", parameterToString(localVarOptionals.SortField.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortDir.IsSet() {
		localVarQueryParams.Add("sort_dir", parameterToString(localVarOptionals.SortDir.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse2008
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
FacebookAdsApiService Get facebook ad info
Get details of a Facebook ad.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param outreachId The outreach id.
 * @param optional nil or *FacebookAdsApiGetFacebookAdsIdOpts - Optional Parameters:
     * @param "Fields" (optional.Interface of []string) -  A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
     * @param "ExcludeFields" (optional.Interface of []string) -  A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.

@return InlineResponse2009
*/

type FacebookAdsApiGetFacebookAdsIdOpts struct { 
	Fields optional.Interface
	ExcludeFields optional.Interface
}

func (a *FacebookAdsApiService) GetFacebookAdsId(ctx context.Context, outreachId string, localVarOptionals *FacebookAdsApiGetFacebookAdsIdOpts) (InlineResponse2009, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InlineResponse2009
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/facebook-ads/{outreach_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"outreach_id"+"}", fmt.Sprintf("%v", outreachId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.ExcludeFields.IsSet() {
		localVarQueryParams.Add("exclude_fields", parameterToString(localVarOptionals.ExcludeFields.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse2009
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

