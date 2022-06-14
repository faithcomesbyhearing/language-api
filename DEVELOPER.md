# sls-go-example

This repository contains a VSCode Remote Container Setup for developing GoLang based AWS Lambda functions using the Serverless Framework.

## Visual Studio Code - Remote Container Usage

> [Official Documentation](https://code.visualstudio.com/docs/remote/containers)

> _Note: Once you meet the prerequisites, you can run **ANY** Visual Studio Code Remote Container, which provides a Docker based development environment ensuring a consistent and reliable set of tooling needed to interact and execute a repository codebase_

### Prerequisites:

1. macOS, Windows, Linux -- [System Requirements](https://code.visualstudio.com/docs/remote/containers#_system-requirements)
2. Docker - [Documentation](https://code.visualstudio.com/docs/remote/containers#_installation)
3. Visual Studio Code - [Official Site](https://code.visualstudio.com/)
4. Remote - Containers _Visual Studio Code extension_ - [Marketplace](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

#### Environment Variables

The remote container honors the following environment variables to be set in the .devcontainer/.env
> _Note: You can copy the .devcontainer/.env.template file to .devcontainer/.env and supply the following variables_

* AWS_ACCESS_KEY_ID
* AWS_SECRET_ACCESS_KEY
* AWS_REGION _(optional) - defaults to `us-east-2`_
* DB_ROOT_PASSWORD _for local mysql. (default: password)_
* DB_DATABASE   _(default: LANGUAGE)_
* GITHUB_PERSONAL_ACCESS_TOKEN  _required to pull private GitHub datamodel repo_

####  Developer Configuration

To prepare the environment for _offline_ capabilities, once the repository is opened in the Remote Container, open a Terminal and type:

`yarn`

This will install the required `serverless-offline` dependency.

### Usage & Things you can do
#### Makefile Targets

There are some predefined makefile targets that can be used to simplify common tasks.

* `make deploy` - deploys the defined Lambda functions to AWS
* `make undeploy` - removes the defined Lambda functions from AWS
* `make offline` - runs the Lambda functions defined in serverless.yml locally.  It also provides a mock AWS API Gateway.

> _Note: Due to limitations in serverless-offline with GoLang, this does NOT allow for debugging_

#### Understanding the development environment
##### Host (eg your development machine)
- Docker Desktop (or comparable)
- Visual Studio Code with Remote Container plugin
- source code
    - .devcontainer/devcontainer.json: describes how to configure the Remote Container. Typically contains a Dockerfile or a docker-composes.yml, and any associated files. When the developer invokes Vscode "Reopen in container", devcontainer.json is consulted
##### VsCode Remote Container
When the vscode project is opened in Remote Container, the developer is actually executing within the Docker environment described by .devcontainer/devcontainer.json. In the bottom left of Vscode, it reads "Dev Container: Go and Mysql". Port 3306 is exposed based on configuration in docker-compose.yml

From this configuration, the developer can rapidly edit, build and test following this procedure:
1. from a terminal, type `make offline`. The images will be built if necessary, and then the containers started, following the configuration in docker-compose.yml. Note that sls offline exposes two additional ports: 3000 (for API Gateway requests) and 3002.
2. using Postman, submit HTTP requests similar to https://localhost:3000/language
    - the request will be handled by sls offline plugin listening at port 3000. The plugin will translate the HTTP request to an API Gateway Proxy request and pass it to the lambda container
    - the sls offline plugin will translate the API Gateway Proxy response to an HTTP response and return it to Postman

Note that you cannot run the debugger in this configuration.

#### Debugging

This Remote Container is preconfigured so that any Go Lambda handler can be debugged individually.   There is **NOT**, however, currently a way to do this via an API (over http(s)).  Since a Lambda function is essentially just an RPC call, we instead will simply start the Function with the Debugger attached, and then make an "RPC" call sending the desired payload to the function.

In this case the input type for these function(s) is an [APIGatewayProxyRequest](https://github.com/aws/aws-lambda-go/blob/main/events/apigw.go#L6).  Therefore,  input JSON must be in that form.  As a convention, there is an `events` folder at the root of the workspace to house these event payloads.  You can create as many of these files as necessary for testing and debugging your functions.  You can look at the pre-existing files in the event directory as examples, but you can also find more examples [here](https://github.com/aws/aws-lambda-go/tree/main/events/testdata)

To debug GoLang based Lambda functions, use the following procedure:

1. Open the .go file containing the Lambda Handler - i.e. hello/main.go
    * Set breakpoints as you wish in the left gutter of the editor
2. Open the _Run and Debug_ View in VS Code _(Shft+Cmd+D in OSX or Shft+Ctrl+D in Windows)_
3. Select _Debug Lambda Function_ and press _Run_
4. This will start the Lambda Function Handler, with the Debugger attached
5. To send an event to the running handler, either:
    1. From the Command Palette _(Cmd+Shft+P or Ctrl+Shft+P)_ select _Run Task_, then choose `send-lambda-event`, then choose the events/*.json file you wish to send as a Payload...or
    2. From a Terminal Window, execute `awslambdarpc -e events/<choosefile>.json`
6. At this point, it will execute your Handler, and any breakpoints you have set will be honored.
7. Debug as normal


#### Recommended Developer Workflow
Development with debugger:
1. open project in Dev Container
2. select go file containing Lambda Handler
3. start debugger
4. `awslambdarpc -e events/foo.json` (or whatever)

Black box testing locally: 
1. open project in Dev Container
2. `make offline`
3. submit requests with Postman to localhost:3000

Testing in AWS:
1. open project in Dev Container
2. `make deploy`
3. submit requests with Postman to API Gateway endpoint url
#### Pitfalls and Troubleshooting

1. Executing sls invoke local does not work.  this command uses the sls offline plugin, which configures the lambda container different than docker-compose.   More investigation is needed to explore this. Workaround: use Postman to call localhost:3000. 

2. awslambdarpc returns "Unexpected EOF"
Explanation: only use awslambdarpc when running in debug mode. The debugger exposes port 3000 for connections. In this configuration, no API Gateway is involved. Events submitted by awslambdarpc must contain the structure provided by the API Gateway
