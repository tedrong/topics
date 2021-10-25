package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/topics/controllers"
	"github.com/topics/database"
	"github.com/topics/forms"
	"github.com/topics/logging"
)

var firstName = "first"
var lastName = "last"
var testEmail = "goTest@test.com"
var testPassword = "123456"
var accessToken string
var refreshToken string
var userUUID string
var auth = new(controllers.AuthController)

//TokenAuthMiddleware ...
//JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	//Custom form validator
	binding.Validator = new(forms.DefaultValidator)

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)
		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)
		v1.GET("/user/logout", user.Logout)
		v1.PUT("/user/renew/:uuid", user.Renew)

		/*** START AUTH ***/
		auth := new(controllers.AuthController)
		//Refresh the token when needed to generate new access_token and refresh_token for the user
		v1.POST("/token/refresh", auth.Refresh)

		/*** START DASHBOARD ***/
		dashboard := new(controllers.DashboardController)
		v1.GET("/dashboard/system/info", TokenAuthMiddleware(), dashboard.SystemInfo)
		v1.GET("/dashboard/system/info/history/:day", TokenAuthMiddleware(), dashboard.SystemInfoHistory)
		v1.GET("/dashboard/system/client/type/percentage", TokenAuthMiddleware(), dashboard.ClientTypePercentage)
		v1.GET("/dashboard/system/log/:line", TokenAuthMiddleware(), dashboard.SystemLog)
	}

	return r
}

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}
	logCfg := logging.Config{
		ConsoleLoggingEnabled: false,
		EncodeLogsAsJson:      true,
		FileLoggingEnabled:    true,
		Directory:             "../tmp",
		Filename:              "test.log",
		MaxSize:               32,
		MaxBackups:            2,
		MaxAge:                32,
	}
	zlog := logCfg.Init()
	zlog.Info().Msg("zlog system initialized")
	database.Init(1)

	exitVal := m.Run()
	os.Exit(exitVal)
}

/**
* TestRegister
* Test user registration
*
* Must return response code 200
 */
func TestRegister(t *testing.T) {
	r := SetupRouter()
	registerForm := forms.RegisterForm{
		FirstName: firstName,
		LastName:  lastName,
		Email:     testEmail,
		Password:  testPassword,
	}
	data, _ := json.Marshal(registerForm)

	req, err := http.NewRequest("POST", "/v1/user/register", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

/**
* TestRegisterInvalidEmail
* Test user registration with invalid email
*
* Must return response code 406
 */
func TestRegisterInvalidEmail(t *testing.T) {
	r := SetupRouter()
	registerForm := forms.RegisterForm{
		FirstName: firstName,
		LastName:  lastName,
		Email:     "invalid@email",
		Password:  testPassword,
	}
	data, _ := json.Marshal(registerForm)

	req, err := http.NewRequest("POST", "/v1/user/register", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusNotAcceptable, res.Code)
}

/**
* TestLogin
* Test user login
* and get the access_token and refresh_token stored
*
* Must return response code 200
 */
func TestLogin(t *testing.T) {
	r := SetupRouter()
	var resJSON struct {
		Message string `json:"message"`
		User    struct {
			CreatedAt int64  `json:"created_at"`
			Email     string `json:"email"`
			ID        int64  `json:"id"`
			UUID      string `json:"UUID"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			UpdatedAt int64  `json:"updated_at"`
		} `json:"user"`
		Token struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"token"`
	}
	loginForm := forms.LoginForm{
		Email:    testEmail,
		Password: testPassword,
	}
	data, _ := json.Marshal(loginForm)

	req, err := http.NewRequest("POST", "/v1/user/login", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "test")
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	json.Unmarshal(body, &resJSON)

	accessToken = resJSON.Token.AccessToken
	refreshToken = resJSON.Token.RefreshToken
	userUUID = resJSON.User.UUID

	assert.Equal(t, http.StatusOK, res.Code)
}

/**
* TestInvalidLogin
* Test invalid login
*
* Must return response code 406
 */
func TestInvalidLogin(t *testing.T) {
	r := SetupRouter()
	loginForm := forms.LoginForm{
		Email:    "wrong@email.com",
		Password: testPassword,
	}
	data, _ := json.Marshal(loginForm)

	req, err := http.NewRequest("POST", "/v1/user/login", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "test")
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotAcceptable, res.Code)
}

/**
* TestRenw
* Test renew user information
*
* Must return response code 200
 */
func TestRenewUser(t *testing.T) {
	r := SetupRouter()
	renewForm := forms.RenewForm{
		FirstName: "Renew" + firstName,
		LastName:  "Renew" + lastName,
	}
	data, _ := json.Marshal(renewForm)

	req, err := http.NewRequest("PUT", "/v1/user/renew/"+userUUID, bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

/**
* TestSystemInfo
* Test get system information
*
* Must return response code 200
 */
func TestSystemInfo(t *testing.T) {
	r := SetupRouter()
	req, err := http.NewRequest("GET", "/v1/dashboard/system/info", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

/**
* TestSystemInfoHistory
* Test get system information history
*
* Must return response code 200
 */
func TestSystemInfoHistory(t *testing.T) {
	r := SetupRouter()

	req, err := http.NewRequest("GET", "/v1/dashboard/system/info/history/7", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

/**
* TestClientType
* Test get client type information
*
* Must return response code 200
 */
func TestClientType(t *testing.T) {
	r := SetupRouter()

	req, err := http.NewRequest("GET", "/v1/dashboard/system/client/type/percentage", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

/**
* TestSystemLog
* Test get system logs
*
* Must return response code 200
 */
func TestSystemLog(t *testing.T) {
	r := SetupRouter()

	req, err := http.NewRequest("GET", "/v1/dashboard/system/log/1", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

/**
* TestRefreshToken
* Test refreshing the token with valid refresh_token
*
* Must return response code 200
 */
func TestRefreshToken(t *testing.T) {
	r := SetupRouter()
	tokenForm := forms.Token{
		RefreshToken: refreshToken,
	}
	data, _ := json.Marshal(tokenForm)

	req, err := http.NewRequest("POST", "/v1/token/refresh", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

/**
* TestInvalidRefreshToken
* Test refreshing the token with invalid refresh_token
*
* Must return response code 401
 */
func TestInvalidRefreshToken(t *testing.T) {
	r := SetupRouter()
	//Since we didn't update it in the test before - this will not be valid anymore
	tokenForm := forms.Token{
		RefreshToken: refreshToken,
	}
	data, _ := json.Marshal(tokenForm)

	req, err := http.NewRequest("POST", "/v1/token/refresh", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
}

/**
* TestUserSignout
* Test logout a user
*
* Must return response code 200
 */
func TestUserLogout(t *testing.T) {
	r := SetupRouter()

	req, err := http.NewRequest("GET", "/v1/user/logout", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", accessToken))
	if err != nil {
		t.Error(err)
	}

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

/**
* TestCleanUp
* Deletes the created user with it's articles
*
* Must pass
 */
func TestCleanUp(t *testing.T) {
	var user database.User
	result := database.GetPG(database.DBContent).Where("email = ?", strings.ToLower(testEmail)).Delete(&user)
	if result.Error != nil {
		t.Error(result.Error)
	}
}
