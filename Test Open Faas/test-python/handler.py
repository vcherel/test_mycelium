import requests

def handle(req):
    # Replace 'function2-service' with the actual hostname of your second function
    function2_url = "http://test-python-2:8080"  # Assuming default OpenFaaS port

    # Send an HTTP GET request to function2
    response = requests.get(function2_url)

    if response.status_code == 200:
        return f"Function 1 called Function 2, and Function 2 responded: {response.text}"
    else:
        return "Error: Unable to communicate with Function 2"
