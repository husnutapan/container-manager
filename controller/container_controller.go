package controller

import (
	"container-manager/configs"
	"container-manager/models"
	"context"
	"github.com/gorilla/mux"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type BaseHandler struct {
	kubeConf *configs.ClientConfig
}

func NewBaseHandler(conf *configs.ClientConfig) *BaseHandler {
	return &BaseHandler{
		kubeConf: conf,
	}
}

func (bh *BaseHandler) GetNamespaces(w http.ResponseWriter, r *http.Request) {
	var namespaces []models.Namespaces
	list, _ := bh.kubeConf.GetClientSet().CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	for _, item := range list.Items {
		namespace := models.Namespaces{Name: item.ObjectMeta.Name}
		namespaces = append(namespaces, namespace)
	}
	configs.JSON(w, namespaces)
}

func (bh *BaseHandler) GetPods(w http.ResponseWriter, r *http.Request) {
	var pods []models.Pod
	vars := mux.Vars(r)
	namespace := vars["namespace"]
	list, _ := bh.kubeConf.GetClientSet().CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	for _, item := range list.Items {
		pod := models.Pod{Name: item.ObjectMeta.Name}
		pods = append(pods, pod)
	}
	configs.JSON(w, pods)
}

func (bh *BaseHandler) DeletePod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	podName := vars["podsName"]
	namespace := vars["namespace"]
	bh.kubeConf.GetClientSet().CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	status := "Success"
	configs.JSON(w, status)
}
