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

* AWS_ACCESS_KEY_ID
* AWS_SECRET_ACCESS_KEY
* AWS_REGION _(optional) - defaults to `us-east-2`_

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