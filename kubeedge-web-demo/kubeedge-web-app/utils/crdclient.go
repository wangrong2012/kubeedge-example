package utils

import (
	"log"

	"github.com/kubeedge/kubeedge/cloud/pkg/devicecontroller/apis/devices/v1alpha2"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
)

// NewCRDClient is used to create a restClient for crd
func NewCRDClient(cfg *rest.Config) (*rest.RESTClient, error) {
	scheme := runtime.NewScheme()
	schemeBuilder := runtime.NewSchemeBuilder(addDeviceCrds)

	err := schemeBuilder.AddToScheme(scheme)
	if err != nil {
		return nil, err
	}

	config := *cfg
	config.APIPath = "/apis"
	config.GroupVersion = &v1alpha2.SchemeGroupVersion
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		log.Fatalf("Failed to create REST Client due to error %v", err)
		return nil, err
	}

	return client, nil
}

func addDeviceCrds(scheme *runtime.Scheme) error {
	// Add Device
	scheme.AddKnownTypes(v1alpha2.SchemeGroupVersion, &v1alpha2.Device{}, &v1alpha2.DeviceList{})
	v1.AddToGroupVersion(scheme, v1alpha2.SchemeGroupVersion)
	// Add DeviceModel
	scheme.AddKnownTypes(v1alpha2.SchemeGroupVersion, &v1alpha2.DeviceModel{}, &v1alpha2.DeviceModelList{})
	v1.AddToGroupVersion(scheme, v1alpha2.SchemeGroupVersion)

	return nil
}
