# Deploy A Golang App to Minikube with HPA

## Setup Steps

1. Dockerize the app, build and push to Dockerhub.
2. Create Kubernetes deployment files.
3. Install Postgres using Helm ([link](https://artifacthub.io/packages/helm/bitnami/postgresql)).
4. Create kubernetes namespace `kubectl create ns dev`
5. Create `db-credentials` secret in the `dev` namespace with your values for `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `JWT_PRIVATE_KEY`.
    `kubectl -n dev create secret generic db-credentials --from-literal=DB_HOST=... --from-literal=DB_USER=... --from-literal=DB_PASSWORD=... --from-literal=DB_NAME=... --from-literal=JWT_PRIVATE_KEY=anystringofyourchoice"`
6. Update the deployment manifest image url in the Kubernetes manifests located in `kubernetes` folder based on your computer architecture `profemzy/training-api:vArm` or `profemzy/training-api:v2` for arm and amd64 respectively.
5. Deploy the app using the Kubernetes manifests `kubectl -n dev apply -f kubernetes/`
6. Enable needed minikube addons `minikube addons enable metrics-server`, `minikube addons enable ingress`. 
6. Run `minikube tunnel` for ingress traffic.
7. Test the app with a POST request to create a user.
    `curl -X POST -H "Content-Type: application/json" -d '{"username": "demo", "password": "someverylongPASSWORD"}' http://localhost/auth/register`
8. Test HPA by uncommenting and applying `job.yaml` ensure you use the appropriate image for your architecture. `profemzy/training-load-test:vArm` or `profemzy/training-load-test:latest` for arm and amd64 respectively.
    `kubectl -n dev apply -f kubernetes/job.yaml`
9. Monitor autoscaling with `kubectl -n dev get hpa`.
10. Monitor pod creation with `kubectl -n dev get pods`.
