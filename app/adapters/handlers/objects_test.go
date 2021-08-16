package handlers

import (
	"cache/objects"
	"cache/objects/ports"
	"cache/objects/ports/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type objectsHandlerSuite struct {
	suite.Suite
	engine         *gin.Engine
	handler        *ObjectsHandler
	objectsManager *mocks.ObjectManager
}

func (s *objectsHandlerSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	s.engine = gin.Default()

	s.objectsManager = new(mocks.ObjectManager)
	s.handler = newObjectsHandler(s.objectsManager)
}

func TestObjectsHandlerSuiteInit(t *testing.T) {
	suite.Run(t, new(objectsHandlerSuite))
}

func (s *objectsHandlerSuite) TestObjectsHandler_Save_bad_request() {
	rawData := `{"user": "test"}`
	obj := objects.NewObject(rawData)

	s.objectsManager.
		On("Save", mock.Anything, mock.Anything).
		Return(obj, nil).
		Once()

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request, _ = http.NewRequest(http.MethodPost, "/object", strings.NewReader(rawData))
	s.handler.Save(c)
	s.Equal(http.StatusBadRequest, recorder.Code)
}

func (s *objectsHandlerSuite) TestObjectsHandler_Save_InsufficientStorage() {
	rawData := `{"user": "test"}`
	params := gin.Params{gin.Param{
		Key:   "key",
		Value: "1",
	}}

	s.objectsManager.
		On("Save", mock.Anything, mock.Anything).
		Return(nil, ports.ObjectNotStored).
		Once()

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request, _ = http.NewRequest(http.MethodPost, "/object", strings.NewReader(rawData))
	c.Params = params
	s.handler.Save(c)
	s.Equal(http.StatusInsufficientStorage, recorder.Code)
}

func (s *objectsHandlerSuite) TestObjectsHandler_Save_Successfully() {
	rawData := `{"user": "test"}`
	obj := objects.NewObject(rawData)

	params := gin.Params{gin.Param{
		Key:   "key",
		Value: "1",
	}}

	s.objectsManager.
		On("Save", mock.Anything, mock.Anything).
		Return(obj, nil).
		Once()

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request, _ = http.NewRequest(http.MethodPost, "/object", strings.NewReader(rawData))

	c.Params = params
	s.handler.Save(c)
	s.Equal(http.StatusOK, recorder.Code)
}
