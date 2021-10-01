When I am developing some Golang application that running on Kubernetes, because my Kubernetes is running on some cheap VPS,  so the network is not that stable, and my application running inside will always panic, and I want to send the panic message out automatically. 

So I make this application, it will monitor all pods, or you can specify a namespace to listen on, when a pod crashed, it will get the latest output and call the webhook.

About the webhook, please see my another project, it can send message to multiple services like telegram, slack, matrix,  or email, and you are easy to extend the webhook handler and service handers.