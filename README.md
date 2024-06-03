# Deploy A Golang App to Minikube with HPA

## Setup Steps

1. Point your terminal to use docker daemon inside minikube run `eval $(minikube -p minikube docker-env)`
2. To verify your terminal is using minikube’s docker-env you can check the value of the environment variable `MINIKUBE_ACTIVE_DOCKERD` to reflect the cluster name. 
    `echo $MINIKUBE_ACTIVE_DOCKERD`
3. Build the docker image using `docker build -t training-app . `
4. Confirm that your newly built image is available in minikube `$ minikube image ls --format table`
5. Create kubernetes namespace `kubectl create ns dev`
6. Add bitnami `helm repo add https://charts.bitnami.com/bitnami`
5. Install Postgres using Helm ([link](https://artifacthub.io/packages/helm/bitnami/postgresql)). `helm -n dev install db bitnami/postgresql --version 15.5.1`
6. Create `db-credentials` secret in the `dev` namespace with your values for `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `JWT_PRIVATE_KEY`.
    `kubectl -n dev create secret generic db-credentials --from-literal=DB_HOST=... --from-literal=DB_USER=... --from-literal=DB_PASSWORD=... --from-literal=DB_NAME=... --from-literal=JWT_PRIVATE_KEY=anystringofyourchoice`
8. Deploy the app using the Kubernetes manifests `kubectl -n dev apply -f kubernetes/`
9. Enable needed minikube addons `minikube addons enable metrics-server`, `minikube addons enable ingress`. 
10. In another terminal Run `minikube tunnel` for ingress traffic.
11. Test the app with a POST request to create a user.
    `curl -X POST -H "Content-Type: application/json" -d '{"username": "demo", "password": "someverylongPASSWORD"}' http://localhost/auth/register`
12. Install hey for mac use `brew install hey` or `sudo apt install hey` depending on your platform
13. Load test the app using hey `hey -n 100000 -c 10 http://localhost/api/health` this sends 100000 requests with a concurrency of 10 to your API endpoint
13. Monitor autoscaling with `kubectl -n dev get hpa`.
14. Monitor pod creation with `kubectl -n dev get pods`.
