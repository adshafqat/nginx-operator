operator-sdk init --domain com --license apache2 --owner "Adeel Shafqat" --repo=github.com/adshafqat/nginx-operator
operator-sdk create api --group zaynsolutions --version v1 --kind ZaynNginx --resource=true --controller=true
// ZaynNginxSpec defines the desired state of ZaynNginx
type ZaynNginxSpec struct {
	Replicas int32 `json:"replicas"`
}

// ZaynNginxStatus defines the observed state of ZaynNginx
type ZaynNginxStatus struct {vi config/samples/zaynsolutions_v1_zaynnginx.yaml
	PodNames []string `json:"podNames"`
}

docker login docker.io
ashafqat
password
make docker-build docker-push IMG=docker.io/ashafqat/nginx-operator:v1.0
docker images
make intall
Create a new app using oc new app and image name or using console

vi config/samples/zaynsolutions_v1_zaynnginx.yaml
apiVersion: zaynsolutions.com/v1
kind: ZaynNginx
metadata:
  name: zaynnginx-sample
spec:
  replicas: 2
  # Add fields here
  foo: bar


kubectl apply -f config/samples/zaynsolutions_v1_zaynnginx.yaml 
