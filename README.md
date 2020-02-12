# message-mediator
Design and stub implementation of a message mediator service

## Overall architecutre: 
![UML](https://img.techpowerup.org/200212/untitled-diagram.jpg)

## Details of diagram above:

### POST /email/batch
Both API handlers `/email` and `/email/batch` are similar enough that they are grouped together in the diagram above.  `/email` will simply put the incoming email onto the queue as a message.  `/email/batch` will do the same, but will break the body of the incomming request into 1 queue message per email.  

Both queue messages will also contain the `messageID` GUID that is returned by the API

#### High Availability and Performance Concerns
One main limitation here is how well the API is able to horizontaly scale, the concern of which is outside the scope of this project
The other concern that is the size of the email body and attachment.  The API instance handling the request will need sufficent memory to handle expected body sizes.  In general a queueing service (ex. RabbitMQ) is capable of message sizes exceeding reasonable emails.

### Queue
Specific queue implementation is not particularly important.  The nature of a email sending provided is that the client traffic tends to be bursty.  The queue acts as a buffer to allow down stream resources time to process the request and scale.

#### High Availability and Performance Concerns
It is reasonable to expect a a clustered queuing service (Ex. RabbitMQ) would be able to scale as needed to accomidate the load.  
One feature that would make the system much more recoverable in the case of failure is if the messages are persisted to disk.

### Message Mediator
