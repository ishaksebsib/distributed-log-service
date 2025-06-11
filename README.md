# TODO

**Part I – Foundation**

- [x] 1. Define a Data Model
- [x] 2. Structure Data with Protocol Buffers  
     ↳ Defined `Record` schema using Protobuf for efficient, versionable binary serialization
- [x] 3. Build a Log  
     ↳ Implemented segment, index, and store for local persistent logging

---

**Part II – Network**

- [x] 4. Serve Requests with gRPC  
     ↳ Expose the log over the network using gRPC and define service interfaces
- [x] 5. Secure Your Services  
     ↳ Add TLS encryption and client authentication to secure communication
- [ ] 6. Observe Your Systems  
     ↳ Integrate metrics, logging, and tracing to make the system observable

---

**Part III – Distribute**

- [ ] 7. Server-to-Server Service Discovery  
     ↳ Enable nodes to discover each other dynamically in a distributed setting
- [ ] 8. Coordinate Your Services with Consensus  
     ↳ Implement Raft consensus for leader election and replicated logs
- [ ] 9. Discover Servers and Load Balance from the Client  
     ↳ Add client-side load balancing and dynamic server discovery

---

**Part IV – Deploy**

- [ ] 10. Deploy Applications with Kubernetes Locally  
      ↳ Set up and run the system in a local Kubernetes cluster (e.g., Minikube)
- [ ] 11. Deploy Applications with Kubernetes to the Cloud  
      ↳ Deploy and scale the system on a cloud provider using Kubernetes
