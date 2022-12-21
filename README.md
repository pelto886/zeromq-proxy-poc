# Zeromq router-dealer demo

This demo is trying to show the tendencies that exist in request-reply sockets, and how to solve those.

Note: This demo is not about resource effectiveness, but resource availability.

the problem:

start a solo-service
start a client-service

the solo-service will randomly reply between 0 and 100ms onReceive

observe that roughly 10% of the messages are dropped.

now add a client-service, and see this percentage rice quickly, as you add more and more services consuming it.

solution:

now start a router-dealer
and start 3-5 compute-services

the compute services are exacly like the solo, but optimised socket management to be combined with router-dealer

now start the two client-services again. observe the percentage down to 10% again.

conclusion: we can dynamically increase availability of a service though router-dealer middleman, without creating unnecessary networking options.