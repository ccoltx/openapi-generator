/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type UserApi interface {

	/*
	 * CreateUser Create user
	 * This can only be done by the logged in user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiCreateUserRequest
	 */
	CreateUser(ctx _context.Context) ApiCreateUserRequest

	/*
	 * CreateUserExecute executes the request
	 */
	CreateUserExecute(r ApiCreateUserRequest) (*_nethttp.Response, GenericOpenAPIError)

	/*
	 * CreateUsersWithArrayInput Creates list of users with given input array
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiCreateUsersWithArrayInputRequest
	 */
	CreateUsersWithArrayInput(ctx _context.Context) ApiCreateUsersWithArrayInputRequest

	/*
	 * CreateUsersWithArrayInputExecute executes the request
	 */
	CreateUsersWithArrayInputExecute(r ApiCreateUsersWithArrayInputRequest) (*_nethttp.Response, GenericOpenAPIError)

	/*
	 * CreateUsersWithListInput Creates list of users with given input array
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiCreateUsersWithListInputRequest
	 */
	CreateUsersWithListInput(ctx _context.Context) ApiCreateUsersWithListInputRequest

	/*
	 * CreateUsersWithListInputExecute executes the request
	 */
	CreateUsersWithListInputExecute(r ApiCreateUsersWithListInputRequest) (*_nethttp.Response, GenericOpenAPIError)

	/*
	 * DeleteUser Delete user
	 * This can only be done by the logged in user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param username The name that needs to be deleted
	 * @return ApiDeleteUserRequest
	 */
	DeleteUser(ctx _context.Context, username string) ApiDeleteUserRequest

	/*
	 * DeleteUserExecute executes the request
	 */
	DeleteUserExecute(r ApiDeleteUserRequest) (*_nethttp.Response, GenericOpenAPIError)

	/*
	 * GetUserByName Get user by user name
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param username The name that needs to be fetched. Use user1 for testing.
	 * @return ApiGetUserByNameRequest
	 */
	GetUserByName(ctx _context.Context, username string) ApiGetUserByNameRequest

	/*
	 * GetUserByNameExecute executes the request
	 * @return User
	 */
	GetUserByNameExecute(r ApiGetUserByNameRequest) (User, *_nethttp.Response, GenericOpenAPIError)

	/*
	 * LoginUser Logs user into the system
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiLoginUserRequest
	 */
	LoginUser(ctx _context.Context) ApiLoginUserRequest

	/*
	 * LoginUserExecute executes the request
	 * @return string
	 */
	LoginUserExecute(r ApiLoginUserRequest) (string, *_nethttp.Response, GenericOpenAPIError)

	/*
	 * LogoutUser Logs out current logged in user session
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiLogoutUserRequest
	 */
	LogoutUser(ctx _context.Context) ApiLogoutUserRequest

	/*
	 * LogoutUserExecute executes the request
	 */
	LogoutUserExecute(r ApiLogoutUserRequest) (*_nethttp.Response, GenericOpenAPIError)

	/*
	 * UpdateUser Updated user
	 * This can only be done by the logged in user.
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param username name that need to be deleted
	 * @return ApiUpdateUserRequest
	 */
	UpdateUser(ctx _context.Context, username string) ApiUpdateUserRequest

	/*
	 * UpdateUserExecute executes the request
	 */
	UpdateUserExecute(r ApiUpdateUserRequest) (*_nethttp.Response, GenericOpenAPIError)
}

// UserApiService UserApi service
type UserApiService service

type ApiCreateUserRequest struct {
	ctx _context.Context
	ApiService UserApi
	body *User
}

func (r ApiCreateUserRequest) Body(body User) ApiCreateUserRequest {
	r.body = &body
	return r
}

func (r ApiCreateUserRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.CreateUserExecute(r)
}

/*
 * CreateUser Create user
 * This can only be done by the logged in user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateUserRequest
 */
func (a *UserApiService) CreateUser(ctx _context.Context) ApiCreateUserRequest {
	return ApiCreateUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *UserApiService) CreateUserExecute(r ApiCreateUserRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.CreateUser")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/user"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		executionError.error = "body is required and must be specified"
		return nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}

type ApiCreateUsersWithArrayInputRequest struct {
	ctx _context.Context
	ApiService UserApi
	body *[]User
}

func (r ApiCreateUsersWithArrayInputRequest) Body(body []User) ApiCreateUsersWithArrayInputRequest {
	r.body = &body
	return r
}

func (r ApiCreateUsersWithArrayInputRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.CreateUsersWithArrayInputExecute(r)
}

/*
 * CreateUsersWithArrayInput Creates list of users with given input array
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateUsersWithArrayInputRequest
 */
func (a *UserApiService) CreateUsersWithArrayInput(ctx _context.Context) ApiCreateUsersWithArrayInputRequest {
	return ApiCreateUsersWithArrayInputRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *UserApiService) CreateUsersWithArrayInputExecute(r ApiCreateUsersWithArrayInputRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.CreateUsersWithArrayInput")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/user/createWithArray"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		executionError.error = "body is required and must be specified"
		return nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}

type ApiCreateUsersWithListInputRequest struct {
	ctx _context.Context
	ApiService UserApi
	body *[]User
}

func (r ApiCreateUsersWithListInputRequest) Body(body []User) ApiCreateUsersWithListInputRequest {
	r.body = &body
	return r
}

func (r ApiCreateUsersWithListInputRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.CreateUsersWithListInputExecute(r)
}

/*
 * CreateUsersWithListInput Creates list of users with given input array
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateUsersWithListInputRequest
 */
func (a *UserApiService) CreateUsersWithListInput(ctx _context.Context) ApiCreateUsersWithListInputRequest {
	return ApiCreateUsersWithListInputRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *UserApiService) CreateUsersWithListInputExecute(r ApiCreateUsersWithListInputRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.CreateUsersWithListInput")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/user/createWithList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		executionError.error = "body is required and must be specified"
		return nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}

type ApiDeleteUserRequest struct {
	ctx _context.Context
	ApiService UserApi
	username string
}


func (r ApiDeleteUserRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.DeleteUserExecute(r)
}

/*
 * DeleteUser Delete user
 * This can only be done by the logged in user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param username The name that needs to be deleted
 * @return ApiDeleteUserRequest
 */
func (a *UserApiService) DeleteUser(ctx _context.Context, username string) ApiDeleteUserRequest {
	return ApiDeleteUserRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

/*
 * Execute executes the request
 */
func (a *UserApiService) DeleteUserExecute(r ApiDeleteUserRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.DeleteUser")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/user/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", _neturl.PathEscape(parameterToString(r.username, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}

type ApiGetUserByNameRequest struct {
	ctx _context.Context
	ApiService UserApi
	username string
}


func (r ApiGetUserByNameRequest) Execute() (User, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.GetUserByNameExecute(r)
}

/*
 * GetUserByName Get user by user name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param username The name that needs to be fetched. Use user1 for testing.
 * @return ApiGetUserByNameRequest
 */
func (a *UserApiService) GetUserByName(ctx _context.Context, username string) ApiGetUserByNameRequest {
	return ApiGetUserByNameRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

/*
 * Execute executes the request
 * @return User
 */
func (a *UserApiService) GetUserByNameExecute(r ApiGetUserByNameRequest) (User, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.GetUserByName")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/user/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", _neturl.PathEscape(parameterToString(r.username, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiLoginUserRequest struct {
	ctx _context.Context
	ApiService UserApi
	username *string
	password *string
}

func (r ApiLoginUserRequest) Username(username string) ApiLoginUserRequest {
	r.username = &username
	return r
}
func (r ApiLoginUserRequest) Password(password string) ApiLoginUserRequest {
	r.password = &password
	return r
}

func (r ApiLoginUserRequest) Execute() (string, *_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.LoginUserExecute(r)
}

/*
 * LoginUser Logs user into the system
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiLoginUserRequest
 */
func (a *UserApiService) LoginUser(ctx _context.Context) ApiLoginUserRequest {
	return ApiLoginUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return string
 */
func (a *UserApiService) LoginUserExecute(r ApiLoginUserRequest) (string, *_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.LoginUser")
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarPath := localBasePath + "/user/login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.username == nil {
		executionError.error = "username is required and must be specified"
		return localVarReturnValue, nil, executionError
	}
	if r.password == nil {
		executionError.error = "password is required and must be specified"
		return localVarReturnValue, nil, executionError
	}

	localVarQueryParams.Add("username", parameterToString(*r.username, ""))
	localVarQueryParams.Add("password", parameterToString(*r.password, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarReturnValue, localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, executionError
}

type ApiLogoutUserRequest struct {
	ctx _context.Context
	ApiService UserApi
}


func (r ApiLogoutUserRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.LogoutUserExecute(r)
}

/*
 * LogoutUser Logs out current logged in user session
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiLogoutUserRequest
 */
func (a *UserApiService) LogoutUser(ctx _context.Context) ApiLogoutUserRequest {
	return ApiLogoutUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *UserApiService) LogoutUserExecute(r ApiLogoutUserRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.LogoutUser")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/user/logout"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}

type ApiUpdateUserRequest struct {
	ctx _context.Context
	ApiService UserApi
	username string
	body *User
}

func (r ApiUpdateUserRequest) Body(body User) ApiUpdateUserRequest {
	r.body = &body
	return r
}

func (r ApiUpdateUserRequest) Execute() (*_nethttp.Response, GenericOpenAPIError) {
	return r.ApiService.UpdateUserExecute(r)
}

/*
 * UpdateUser Updated user
 * This can only be done by the logged in user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param username name that need to be deleted
 * @return ApiUpdateUserRequest
 */
func (a *UserApiService) UpdateUser(ctx _context.Context, username string) ApiUpdateUserRequest {
	return ApiUpdateUserRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

/*
 * Execute executes the request
 */
func (a *UserApiService) UpdateUserExecute(r ApiUpdateUserRequest) (*_nethttp.Response, GenericOpenAPIError) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		executionError       GenericOpenAPIError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserApiService.UpdateUser")
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarPath := localBasePath + "/user/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", _neturl.PathEscape(parameterToString(r.username, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		executionError.error = "body is required and must be specified"
		return nil, executionError
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		executionError.error = err.Error()
		return nil, executionError
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		executionError.error = err.Error()
		return localVarHTTPResponse, executionError
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, executionError
}
