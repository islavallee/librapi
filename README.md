# Librapi:  the API Rest key/value store

## How to start

Requirements:
 - docker
 - docker-compose (only for local development envivironment and CI commands)
 - microk8s with dns, ingress and storage addon enabled
 - kubectl (configured to use microk8s cluster)
 - helm

Run this command to start the app in a microk8s cluster
```bash
make librapi-in-mk8s
```
This command will do the following:
 - Build localy your image
 - Push the image to you local microk8s
 - Create a storage class in you cluster
 - Create a namespace
 - Deploy the helm chart in this namespace
 - Add the localhost domain librapi.local to your /etc/hosts

The app should be avaiable at librapi.local/

Follow the swagger documentation to use the API

## How it works

The idea here was to create a key-value storage that was not lost on container update in Kubernetes.

The api follows the restfull api guidelines, even if it's not complex it could be extended easily. The routing naming could also be improved. 
The application layer does nothing fancy here. In our API the user provides the storage key, so there is little to no logic in this layer. A better connection management could be considered.
Storage technology is BoltDB it's simple and fast for our use case. All logic about this key-value database is located in one file. It's really easy to switch to another storage technology. 

The development environment uses reflex who checked for file changes and reload code.
I didn't had time to do tests, I didn't chosed to do TDD here as it was a realy simple project with little to no logic in the code. I didn't had time to set up functional tests.

The code pakages in a docker container that runs as noroot user.

To deploy in kubernetes, I used a storage class as it does the job in microk8s without having to bind a path to my host (I didn't want to have stuff lasting on my computer). In a dedicated kubernetes environement, I could have used more resilient storage using a mounted volume on the node.

## Next Steps (or what could have been done with more time)

 - concurent writing in 'DB'
 - manage application scaling using replicas and autoscaling (needs the previous point)
 - add tests !!!
 - better api error feedback on missing parameters
 - better application error management
 - inject configuration using environement variable

## Source

Hacked implementation key/value store from: https://github.com/rapidloop/skv

