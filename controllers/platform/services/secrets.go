package services

import (
	"context"
	operatorapi "github.com/apache/incubator-kie-kogito-serverless-operator/api/v1alpha08"
	"github.com/apache/incubator-kie-kogito-serverless-operator/container-builder/client"
	"github.com/apache/incubator-kie-kogito-serverless-operator/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func GetSecretKeyValueString(ctx context.Context, client client.Client, secretName string, secretKey string, platform *operatorapi.SonataFlowPlatform) (string, error) {
	secret, err := client.CoreV1().Secrets(platform.Namespace).Get(ctx,
		secretName, metav1.GetOptions{})
	if err != nil {
		klog.V(log.E).InfoS("Error extracting secret: ", "namespace", platform.Namespace, "error", err)
		return "", err
	}

	return string(secret.Data[secretKey]), nil
}
