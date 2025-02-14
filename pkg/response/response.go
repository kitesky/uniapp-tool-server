package response

import (
	"github.com/gin-gonic/gin"
)

type option struct {
	Status  int          `json:"status"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
	Context *gin.Context `json:"-"`
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(c *gin.Context) *option {
	return &option{Context: c}
}

func (s *option) SetStatus(status int) *option {
	s.Status = status
	return s
}

func (s *option) SetCode(code int) *option {
	s.Code = code
	return s
}

func (s *option) SetMessage(message string) *option {
	s.Message = message
	return s
}

func (s *option) SetData(data interface{}) *option {
	s.Data = data
	return s
}

func (s *option) GetStatus() int {
	if s.Status == 0 {
		return 200
	}

	return s.Status
}

func (s *option) GetCode() int {
	if s.Code == 0 {
		return 200
	}

	return s.Code
}

func (s *option) GetMessage() string {
	return s.Message
}

func (s *option) GetData() interface{} {
	return s.Data
}

func (s *option) Success() {
	message := s.GetMessage()
	if message == "" {
		message = "ok"
	}

	data := response{
		Code:    s.GetCode(),
		Message: message,
		Data:    s.GetData(),
	}

	s.Context.JSON(s.GetStatus(), data)
}

func (s *option) Error() {
	code := s.GetCode()

	// 当code=0|200 重置code=400
	if code == 0 || code == 200 {
		code = 400
	}

	data := response{
		Code:    code,
		Message: s.GetMessage(),
	}

	s.Context.JSON(s.GetStatus(), data)
}
