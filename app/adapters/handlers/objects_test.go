package handlers

import (
	"cache/app/conf"
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
	config := &conf.Config{
		Slots:  10,
		TTL:    0,
		Policy: "REJECT",
	}
	s.objectsManager = new(mocks.ObjectManager)
	s.handler = newObjectsHandler(config, s.objectsManager)
}

func TestObjectsHandlerSuiteInit(t *testing.T) {
	suite.Run(t, new(objectsHandlerSuite))
}

func (s *objectsHandlerSuite) TestObjectsHandler_Save_bad_request() {
	rawData := `{"user": "test"}`
	obj := objects.NewObject(rawData)

	s.objectsManager.Mock.
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

	s.objectsManager.Mock.
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

	s.objectsManager.Mock.
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

func (s *objectsHandlerSuite) TestObjectsHandler_Delete_BadRequest() {

	s.objectsManager.Mock.
		On("Delete", mock.Anything).
		Return(ports.ObjectNotFound).
		Once()

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request, _ = http.NewRequest(http.MethodDelete, "/object", nil)
	s.handler.DeleteByKey(c)
	s.Equal(http.StatusBadRequest, recorder.Code)
}

func (s *objectsHandlerSuite) TestObjectsHandler_Delete_ObjectNotFound() {
	params := gin.Params{gin.Param{
		Key:   "key",
		Value: "1",
	}}

	s.objectsManager.Mock.
		On("Delete", mock.Anything).
		Return(ports.ObjectNotFound).
		Once()

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request, _ = http.NewRequest(http.MethodDelete, "/object", nil)
	c.Params = params
	s.handler.DeleteByKey(c)

	s.False(c.Writer.Written())
	s.Equal(http.StatusNotFound, c.Writer.Status())
}

func (s *objectsHandlerSuite) TestObjectsHandler_GetByKey_BadRequest() {

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request, _ = http.NewRequest(http.MethodDelete, "/object", nil)

	s.handler.GetByKey(c)

	s.Equal(http.StatusBadRequest, recorder.Code)
}

func (s *objectsHandlerSuite) TestObjectsHandler_GetByKey_Successfully() {

	params := gin.Params{gin.Param{
		Key:   "key",
		Value: "1",
	}}

	rawData := `{"user": "test"}`
	obj := objects.NewObject(rawData)

	s.objectsManager.Mock.On("GetByKey", mock.Anything).
		Return(obj, nil).Once()

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = params
	c.Request, _ = http.NewRequest(http.MethodDelete, "/object", nil)

	s.handler.GetByKey(c)
	s.Equal(http.StatusOK, recorder.Code)
}
