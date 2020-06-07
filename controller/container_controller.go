package controller

import (
	"container-manager/configs"
	"container-manager/models"
	"context"
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
	configs.JSON(w,namespaces)
}
