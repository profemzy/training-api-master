# Deploy A Golang App to Minikube with HPA

## Setup Steps

1. Dockerize the app, build and push to Dockerhub.
2. Create Kubernetes deployment files.
3. Install Postgres using Helm ([link](https://artifacthub.io/packages/helm/bitnami/postgresql)).
4. Create `db-credentials` secret in the `dev` namespace or namespace of your choice with your values for `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `JWT_PRIVATE_KEY`.
    `kubectl create secret generic db-credentials --from-literal=DB_HOST=... --from-literal=DB_USER=... --from-literal=DB_PASSWORD=... --from-literal=DB_NAME=... --from-literal=JWT_PRIVATE_KEY=...`
5. Update the deployment image url in the Kubernetes manifests based on your computer architecture `profemzy/training-api:vArm` or `profemzy/training-api:v2` for arm and amd64 respectively.
5. Deploy the app using the Kubernetes manifests.
6. Enable needed minikube addons `minikube addons enable metrics-server`, `minikube addons enable ingress`. 
6. Run `minikube tunnel` for ingress traffic.
7. Test the app with a POST request to create a user.
    `curl -X POST -H "Content-Type: application/json" -d '{"username": "demo", "password": "someverylongPASSWORD"}' http://localhost/auth/register`
8. Test HPA by uncommenting and applying `job.yaml` ensure you use the appropriate image for your architecture. `profemzy/training-load-test:vArm` or `profemzy/training-load-test:latest` for arm and amd64 respectively.
    `kubectl apply -f job.yaml`
9. Monitor autoscaling with `kubectl get hpa`.
10. Monitor pod creation with `kubectl get pods`.