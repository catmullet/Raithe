package main

import (
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"strings"
	api "github.com/catmullet/Raithe/api/queue"
	auth "github.com/catmullet/Raithe/app/auth/model"
	"github.com/theplant/testingutils/assert"
	"net/http"
	"github.com/catmullet/Raithe/app/auth/services"
	"encoding/json"
	"github.com/subosito/gotenv"
)

var testToken = "06baa7369051b6e6ed905c5e80b83c0dfa5f8283c7104df2a2a970633df8eeac99688f23eb402dc49e4db9f008eeb107aea806c6b148f9995e6c5cd0c17ce22c1671bcaecb0a4019b4ee86bdf702ffb4988213d2b1012cae491220e405fc92d52f795659343356fa88d0ce6c59dc0fac0a52c54f6ff7494517403ada78a6860b4f6fa62ad47a0768403e037465716050c0cdbdb56a6a39e2c2eaefc62ba7d0399de6b8ebabf3475d9c3cf571e3557386cf27a57f5339f5209064b6e3d02e601480e75b5df55293994cbd2880293fff08a7e6eefffa393e254c497e97db9e771fd8567755ac224778f5e6a8995fe2af11634ab7ffcfbcf2f344aa0a98d726e9465f4c7d9926ade6155ac96125237b6386a11d9f03b9d4ac75d31d692b91fef51af4d9ee3da6b970aa1336f99248fd7efc8904093881a85b4484b4b587e6db2070cc0bcaddb0d934902987b3cca21b309de6636f46963ce950fb955af000e18f257a983f37e109686c28348cb64d5807257b18a797a519633f64c9c75514ae217c219f2479f96776acb92b15c924b573b957f16aa2eaebdafdb64681ff32afa04d82ff50cf0b7d8e4042d1d8a67404f54de3b4409b15459eefaf37ecc51c2c39d14ecdd17d46ed1ea42f7835421bde598869b29c5cb7729f7cbc2528a61823c5a4a513722fabfabdab93203de2f1e58634a6fda931e678d7864f6cd5db463cacd9"
var registerJSON = `{ "agent_name":"test" }`
var secureToken = &auth.SecurityToken{}
var testMessage = func(token string) string { return `{"queue":"test","message":"Hello World","security_token": {"agent_name": "test","token":"` + token + `"}}`}
var testPop = func(token string) string { return `{"queue":"test","security_token": {"agent_name": "test","token":"` + token + `"}}`}



func bootstrapTests(){
	gotenv.Load("env")
}

func TestAgentRegistration(t *testing.T) {

	bootstrapTests()

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/auth/register", strings.NewReader(registerJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// Register as Agent
	services.RegisterAsAgent(c)

	resp := auth.RegisterResponse{}
	json.Unmarshal([]byte(rec.Body.String()), &resp)
	*secureToken = resp.SecurityToken


	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, len(resp.SecurityToken.Token), len(testToken))
}

func TestMessagePush(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/queue/push", strings.NewReader(testMessage(secureToken.Token)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	api.Push(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestMessagePop(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/queue/pop", strings.NewReader(testPop(secureToken.Token)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	api.Pop(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

