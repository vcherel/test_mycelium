import json
import asyncio
import nats

# Define a function to publish a message to a NATS subject
async def publish_message(subject, message):
    # Connect to the NATS server
    nc = await nats.connect("nats://nats-cluster.default.svc.cluster.local:4222")

    # Publish the message to the specified subject
    await nc.publish(subject, json.dumps(message).encode())

    # Ensure the message is sent
    await nc.flush()

    # Close the NATS connection
    await nc.close()

# Handle the incoming request
def handle(req):
    # Create a data payload to be sent as a message
    data = {"message": "Hello from function1!"}

    # Publish the data to a NATS subject named "function1-subject"
    
    # Create an event loop to run the asynchronous function
    loop = asyncio.get_event_loop()
    
    # Run the publish_message function asynchronously
    loop.run_until_complete(publish_message("function1-subject", data))

    # Return the data as a JSON response
    return json.dumps(data)
