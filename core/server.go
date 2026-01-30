package core

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net"
	"net/http"
)

type FelarofServer struct {
	ip       string
	port     int
	token    string
	files    []string
	listener net.Listener
	mux      *http.ServeMux
	template *template.Template
}

func NewServer(files []string, tmpl *template.Template, fileSystem fs.FS) (*FelarofServer, error) {
	ip := GetLocalIP()

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}

	token, err := GenerateToken()
	if err != nil {
		return nil, err
	}

	server := &FelarofServer{
		ip:       ip,
		port:     l.Addr().(*net.TCPAddr).Port,
		token:    token,
		files:    files,
		listener: l,
		mux:      http.NewServeMux(),
		template: tmpl,
	}

	server.mux.Handle("/static/", http.FileServer(http.FS(fileSystem)))

	basePath := "/" + server.token + "/"
	server.mux.HandleFunc(basePath, server.handleMainPage)

	go http.Serve(l, server.mux)

	return server, nil
}

func (s *FelarofServer) handleMainPage(w http.ResponseWriter, r *http.Request) {
	err := s.template.Execute(w, map[string]any{
		"HasFiles": len(s.files) > 0,
		"Success":  r.URL.Query().Get("status") == "ok",
	})
	if err != nil {
		log.Printf("Template execution error: %v", err)
	}
}

func (s *FelarofServer) Close() error {
	return s.listener.Close()
}

func (s *FelarofServer) GetURL() string {
	return fmt.Sprintf("http://%s:%d/%s/", s.ip, s.port, s.token)
}
