package middlewares

import (
	"bytes"
	"fmt"
	"github.com/gbrayhan/microservices-go/tools"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()

	buf := make([]byte, 4096)
	num, err := c.Request.Body.Read(buf)
	if err != nil {
		panic(err)
		return
	}
	reqBody := string(buf[0:num])
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody)))

	// TODO: Viper time_zone in config.json
	loc, _ := time.LoadLocation("America/Mexico_City")
	allDataIO := map[string]interface{}{
		"ruta":          c.FullPath(),
		"request_uri":   c.Request.RequestURI,
		"raw_request":   reqBody,
		"status_code":   c.Writer.Status(),
		"body_response": blw.body.String(),
		"errors":        c.Errors.Errors(),
		"created_at":    time.Now().In(loc).Format("2006-01-02T15:04:05"),
	}
	fmt.Sprintf("%v", allDataIO)

	//array to define which routes will be monitored in all status code
	allLogs := []string{
		"/payment-with-recurrence",
		"/buy-console",
		"/other-route",
	}

	if existAll, _ := tools.InArray(c.FullPath(), allLogs); existAll {

		if c.Writer.Status() == 500 {
			// go SendMailLog(allDataIO)
		}
		// go SaveLogs(allDataIO)
	}
}
