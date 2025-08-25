A server is actually a `transport.Server`, it does all the server logic by relying on other services 

A server holds a [[Service Registery]] and a [[Peer Manager]] which it uses to handle queries and respond to service requests by calling the appropirate service. A server actually bases on a `GRPC server`.