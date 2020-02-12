# message-mediator
Design and stub implementation of a message mediator service

## Overall architecture:
![UML](https://img.techpowerup.org/200212/untitled-diagram.jpg)

## Details of diagram above:

### POST /email/batch
Both API handlers `/email` and `/email/batch` are similar enough that they are grouped together in the diagram above.  `/email` will simply put the incoming email onto the queue as a message.  `/email/batch` will do the same, but will break the body of the incoming request into 1 queue message per email.

Both queue messages will also contain the `messageID` GUID that is returned by the API

#### High Availability and Performance Concerns
One main limitation here is how well the API is able to horizontally scale, the concern of which is outside the scope of this project
The other concern that is the size of the email body and attachment.  The API instance handling the request will need sufficient memory to handle expected body sizes.  In general a queueing service (ex. RabbitMQ) is capable of message sizes exceeding reasonable emails.

### Queue
Specific queue implementation is not particularly important.  The nature of a email sending provided is that the client traffic tends to be bursty.  The queue acts as a buffer to allow down stream resources time to process the request and scale.

#### High Availability and Performance Concerns
It is reasonable to expect a a clustered queuing service (Ex. RabbitMQ) would be able to scale as needed to accommodate the load.
One feature that would make the system much more recoverable in the case of failure is if the messages are persisted to disk.

### Message Mediator
This is inspired by the Mediator behaviour pattern.  Without going into too much detail, the idea is this module coordinates calling business logic on each message and ultimately adds it to a sending pipeline.

#### High Availability and Performance Concerns
This implementation can be deployed as many times as needed to handle load, via running on multiple MVs or in containers.  This module knows what order the business logic must be executed, as well as what must be run in serial and what can be run in parallel.

### Business Logic

#### Message Qualifying
Represents one or many pieces of business logic that has some criteria for what considers a message to be qualified.  It is the responsibility of the Message Mediator to interpret the results, to either log and reject or continue forward.

#### Other Business Logic
The represents any other business logic that needs to be executed on the email.  Any Results it returns will also be the responsibility of the Message Mediator to interpret and act on.  For example the Message Mediator would know what results need to be logged or how to chain together discreet business logic pieces if the are dependant on each other.

#### Pipeline Logic
Another subset of general business logic, this module or service will analyze the email and provide sending pipeline info.  The Message Mediator knows how to map the response of this service, to a specific pipeline.

#### High Availability and Performance Concerns
Depending on the nature of the business logic, it could be contained in a separate code class, package, module, or API.  The case of evaluating emails a DNM (do not market) list would be a likely candidate for a separate API as it would potentially be a third parties data that needs to be checked.  Scaling these items would of course be depending on thier implementation.  If it is just a code module, then expanding the resources available to the mediator would of course address some issues.  In the case of an API, the inherent network latency would make running logic in parallel where possible very important.

### Sending Pipelines
This represents the services that is actually responsible for sending an email.  The responsibility of the Message Mediator is to map the incoming message to a pipeline and also know how to map that message to the correct interface for the pipeline.  That mapping would probably take the form of a simple mapping class.  In this way, we contain all interfacing logic inside the mediator class, and can easily plug in new sending pipelines, regardless of their interface, and add the mapping here.

Depending on the nature and complexity of the interface mapping to the different pipelines, there is a design alternative.  One would be to implement the Facade design pattern as a wrapper around the sending pipelines.  In this, way we could standardize the communication leaving the Message Mediator and contain the pipeline specific mapping logic within each pipeline specific Facade.