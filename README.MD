# Kubernetes studies using golang as server

This project intends to save my kubernetes studies using a very simple server in Golang to simulate a real application

## Preparing the environment

    - We need [Docker](https://docs.docker.com/engine/install/) installed.
    - We're using [Kind](https://kind.sigs.k8s.io/) to run Kubernetes Clusters locally. 
    So you need to install it.
    - In addition we also need to install [Kubectl](https://kubernetes.io/docs/tasks/tools/)
    as our Kubernetes command-line tool

## Setting up our cluster
```sh
    kind create cluster --config=k8s/kind.yaml --name=fullcycle
```

## Running an isolated pod
```sh
    kubectl apply -f k8s/pod.yaml
```

Verify if the pod is up
```sh
    kubectl get pods
```

Right now we don't have any motor running to forward the pod's inner port, so we need to do this manually
```sh
    kubectl port-forward pod/goserver 80:80
```

Now we can access: localhost:80

## Running our replicaset
```sh
    kubectl apply -f k8s/replicaset.yaml
```

Verify if the replicaset is up
```sh
    kubectl get replicasets
```

## Running our deployment
```sh
    kubectl apply -f k8s/deployment.yaml
```

Verify if the deployment is up
```sh
    kubectl get deployments
```

## Running the service
```sh
    kubectl apply -f k8s/service.yaml
```

Verify if the service is up
```sh
    kubectl get services
```

The service will act as a Loadbalancer and distribute the request across the pods.
But like in pods, we need to forward the port manually
```sh
    kubectl port-forward svc/goserver-service 8000:80
```

Now we can access: localhost:8000