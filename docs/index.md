# Kubestack

Kubestack is a infrastructure orchestrator built on top of Kubernetes. The goal is to provide APIs to build private clouds using Kubernetes as a control plane and API.

## Overview

- [**Management**](./management/index.md)  
  As Kubestack is built on top of **bare-metal infrastructure**, it requires foundational management APIs to connect to existing infrastructure, such as network appliances.

- [**Networking**](./networking.md)  
  Kubestack provides low-level APIs to manage network infrastructure, such as `Interfaces`.

## Architectural principles

- **Agentless**  
  Kubestack aims to be agentless, meaning that it does not require any agent to be installed on the managed infrastructure. This may mean an increased amount of management traffic but also avoids the need to install and maintain agents, possibly allowing for easier development of new appliance drivers.

## License

Kubestack is and will always be licensed under the terms of the MIT license.
