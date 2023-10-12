import json
import asyncio
import nats

# Define a function to subscribe to messages on a NATS subject
async def subscribe_message(subject):
    # Create a NATS client instance
    nc = await nats.connect("nats://nats-cluster.default.svc.cluster.local:4222")

    # Define a message handler function for incoming messages
    async def message_handler(msg):
        # Parse the incoming message data as JSON
        data = json.loads(msg.data.decode())
        
        # Print the received message and subject
        print(f"Received message on {msg.subject}: {data}")

    # Subscribe to the specified NATS subject and specify the message handler
    await nc.subscribe(subject, cb=message_handler)

# Handle the incoming request
def handle(req):
    # Create an event loop to run the asynchronous function
    loop = asyncio.get_event_loop()

    # Subscribe to messages on the "function1-subject" NATS subject
    loop.run_until_complete(subscribe_message("function1-subject"))

    # Return a response indicating that Function2 is listening for messages
    return "Function2 is listening for messages..."
