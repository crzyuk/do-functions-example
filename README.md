# Digital Ocean Functions example

This is a simple example of using Digital Ocean functions using Go. 

## Start

1) Install the doctl app - https://docs.digitalocean.com/reference/doctl/how-to/install/

2) Run the deploy command

```
$ make deploy
doctl serverless deploy .
Deploying '/Users/crzy/code/do/web'
  to namespace 'fn-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx'
  on host 'https://faas-lon1-xxxxxxxx.doserverless.co'
Submitted action 'crzy/hello' for remote building and deployment in runtime go:1.20 (id: xxxx)
Submitted action 'crzy/test' for remote building and deployment in runtime go:1.20 (id: xxxx)

Deployed functions ('doctl sls fn get <funcName> --url' for URL):
  - crzy/hello
  - crzy/test
```

3) Invoke the function

```
$ make invoke name=test
doctl serverless fn invoke crzy/test
{
  "body": "This is a test"
}
```


## Where do I go from here? 

Once you have your functions deployed the world is your oyster! 
Have a look inside the packages folder where you can start to create new functions or edit the existing ones. 

You can find the Digital Ocean docs here - https://docs.digitalocean.com/products/functions/