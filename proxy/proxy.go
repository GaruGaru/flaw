package proxy

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Proxy struct {
	config Config
	client http.Client
}

func New(config Config) Proxy {
	return Proxy{
		config: config,
		client: http.Client{},
	}
}

func (pr Proxy) ProxyHandler(writer http.ResponseWriter, request *http.Request) {

	proxyRequest := request.URL
	proxyRequest.Scheme = pr.config.Scheme
	proxyRequest.Host = pr.config.Host

	proxyReq := &http.Request{
		Method: request.Method,
		URL:    proxyRequest,
		Header: request.Header,
	}

	response, err := pr.client.Do(proxyReq)

	if err != nil {
		writer.WriteHeader(500)
		_, err := writer.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		writer.WriteHeader(500)
		_, err := writer.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
		return
	}

	for k, v := range response.Header {
		writer.Header().Add(k, v[0])
	}

	writer.WriteHeader(response.StatusCode)

	_, err = writer.Write(data)

	if err != nil {
		panic(err)
	}

}

func (pr Proxy) Run() error {
	router := mux.NewRouter()

	for _, flaw := range pr.config.EnabledFlaws {
		router.Use(flaw.Middleware)
	}

	router.PathPrefix("/").HandlerFunc(pr.ProxyHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    pr.config.Bind,
	}

	logrus.Infof("flaw started on port %s", pr.config.Bind)
	logrus.Infof("proxying http://%s => %s://%s", pr.config.Bind, pr.config.Scheme, pr.config.Host)

	return srv.ListenAndServe()
}
