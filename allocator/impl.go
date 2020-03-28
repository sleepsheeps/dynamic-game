package allocator

import (
	"agones.dev/agones/pkg/client/clientset/versioned"
	"k8s.io/client-go/rest"
	"net/http"
)

func getOnly(h handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)
			return
		}
		http.Error(w, "Get Only", http.StatusMethodNotAllowed)
	}
}

func deleteOnly(h handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			h(w, r)
			return
		}
		http.Error(w, "Delete Only", http.StatusMethodNotAllowed)
	}
}

func getAgonesClient() (*versioned.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return versioned.NewForConfig(config)
}
