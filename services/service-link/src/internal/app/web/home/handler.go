package home

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"slices"
	"strings"

	"github.com/gorilla/mux"
)

func (s *service) CreateLink(w http.ResponseWriter, r *http.Request) {
	var linkIn Link
	if err := json.NewDecoder(r.Body).Decode(&linkIn); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}
	linkIn.LinkAddrIn = func() string {
		if linkIn.LinkAddrIn != "" {
			return linkIn.LinkAddrIn
		}
		var linkLen = 10
		var linkChars = fmt.Sprintf("%s%s%s",
			"QWERTYUIOPASDFGHJKLZXCVBNM",
			"qwertyuiopasdfghjklzxcvbnm",
			"1234567890",
		)
		result := make([]byte, linkLen)
		for i := range result {
			result[i] = linkChars[rand.Intn(len(linkChars))]
		}

		return string(result)
	}()
	for _, char := range linkIn.LinkAddrIn {
		if !slices.Contains(strings.Split("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm0123456789", ""), string(char)) {
			http.Error(w, "Only the following characters can be present in a username: «A-Z», «a-z», «0-9».", http.StatusInternalServerError)
			return
		}
	}
	if len(linkIn.LinkAddrOut) <= 0 {
		http.Error(w, "Empty link out (perelink)", http.StatusBadRequest)
		return
	}
	createLink, err := s.strg.LinkCreate(linkIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createLink)
}

func (s *service) LinkSearch(w http.ResponseWriter, r *http.Request) {
	link_addr_in := mux.Vars(r)["link_addr_in"]
	if len(link_addr_in) <= 0 {
		http.Error(w, "Invalid link in", http.StatusBadRequest)
		return
	}
	user, err := s.strg.LinkSearch(link_addr_in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
