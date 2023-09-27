package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type CounterImpl struct {
	value int
	mu    sync.Mutex
}

func (c *CounterImpl) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *CounterImpl) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

func (c *CounterImpl) GetCounterValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

type Counter interface {
	Increment()
	Decrement()
	GetCounterValue() int
}

var counter = &CounterImpl{}
var counterTemplate = template.Must(template.New("counter").Parse("<div id=\"counter\">{{.CounterValue}}</div>"))

func GetCounter(c *gin.Context) {
	tmpl, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		if err != nil {
			handleError(c, err)
		}
	}

	data := map[string]int{
		"CounterValue": counter.GetCounterValue(),
	}

	err = tmpl.ExecuteTemplate(c.Writer, "index.html", data)
	if err != nil {
		handleError(c, err)
	}
}

func IncrementCounter(c *gin.Context) {
	counter.Increment()
	renderCounterTemplate(c)
}

func DecrementCounter(c *gin.Context) {
	counter.Decrement()
	renderCounterTemplate(c)

}

func renderCounterTemplate(c *gin.Context) {
	data := map[string]int{
		"CounterValue": counter.GetCounterValue(),
	}

	err := counterTemplate.ExecuteTemplate(c.Writer, "counter", data)
	if err != nil {
		handleError(c, err)
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func handleError(c *gin.Context, err error) {
	c.Writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(c.Writer).Encode(err)
	fmt.Println("Error: ", err)
}
