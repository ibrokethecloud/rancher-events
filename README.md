## Using rancher-events

The rancher-events is based on the go-rancher bindings to extract docker events from rancher.

One of the possible ways to get rancher events is querying the docker daemon directly however this module leverages the rancher events stream to extract the same information.

The benefit of this approach for me is in the fact that the code can be run in a container in the rancher environment and can schedule on any host on the environment.

The keys are injected into the container using the labels described as follows:


https://docs.rancher.com/rancher/v1.0/en/rancher-services/service-accounts/
