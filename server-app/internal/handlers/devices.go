package handlers

import (
	"net/http"
)

func DevicePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/device/add" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
}

func DeviceGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/device/get" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
}

func DevicePut(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/device/update" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
}

func DeviceDelete(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/device/delete" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
}

func DeviceGetAll(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/device/get/all" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
}
