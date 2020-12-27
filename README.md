# Serverless Go in Azure Functions with custom handlers

This repository contains kind of a transcript for the YouTube video [Serverless Go in Azure Functions with custom handlers](https://youtu.be/RPCEH247twU) by Anthony Chu. The video contains a lot of helpful information about custom handlers and beyond.
This repository contains the code for the single steps Anthony is executing in his video. Each step is represented in an own branch. The corresponding starting point of the video is referenced accordingly in the sections describing the branch.

## Prerequisites

The complete tutorial is build in WSL2 on Windows10 to mimic the non-Windows environment of the original video. I used the Ubuntu 20.04 image.
Beside having WSL2 running you also need some further things to be installed in WSL2:

* Go via apt i.e.`sudo apt-get install golang`.
* Node e.g. installed via NVM (Node version manager) as described [here](https://github.com/nvm-sh/nvm#installing-and-updating)
* [Azurite](https://github.com/Azure/Azurite#npm) as storage emulator installed via npm. I also created a dedicated location for storing the data in `~/azurite_store`.
* VS Code [Azure Function Extension](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-azurefunctions)
* VS Code [Azure Storage Extension](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-azurestorage)

You should also install the [Azure Storage Explorer](https://azure.microsoft.com/en-us/features/storage-explorer/) on your Windows host. This can directly access Azurite via port-forwarding from WSL2.

## Remark

As the `local.settings.json` file is usually not committed I created a placeholder file called `local.settings sample.json` that can serve as a template.

## Walk Through

### Setting up the Go Server

_YouTube Reference_: [Link](https://youtu.be/RPCEH247twU?t=124)

_Branch_: main

_Hints_:

* Do not forget to build the go file.
* Ports are forwarded from WSL2, so you can directly access the endpoints from your Windows host.

### Transforming the Server to an Azure Function

_YouTube Reference_: [Link](https://youtu.be/RPCEH247twU?t=204)

_Branch_: function_custom_handler

_Hints_:

* Set the `defaultExecutablePath` to `server` in the `host.json` for the `custom handler`- no `server.exe` as we are on Linux.
* Add the parameter `enableForwardingHttpRequest` and set it to `true`.
* Adopt the `server.go` to fetch the port from the environment variable `FUNCTIONS_CUSTOMHANDLER_PORT`.
* Do not forget to build the go file.

### Adding a Queue Trigger

_YouTube Reference_: [Link](https://youtu.be/RPCEH247twU?t=437)

_Branch_: queue_trigger

_Hints_:

* Create a new function with a queue trigger.
* Adopt the `server.go` file to handle a new route. See code for details.
* Do not forget to build the go file.
* Do not forget to set the `AzureWebJobsStorage` parameter to `UseDevelopmentStorage=true;` in the `local.settings.json` file.
* Do not forget to launch Azurite before starting the function via `azurite --silent --location ~/azurite_store`.

### Deploying to Azure

_YouTube Reference_: [Link](https://youtu.be/RPCEH247twU?t=757)

_Branch_: cross_compile

_Hints_:

* As deployment is done to a Windows app on Azure we need to cross-compile the app. The relevant statement is available in the file `crosscompile.sh`.
* Do not forget to execute `chmod +x` on the file to make it executable.
* Do not forget to set the `defaultExecutablePath` to `server.exe` in the `host.json`.
* In order to be able to proceed with the local scenario, we override the setting in `host.json` the setting in the `local.settings.json` via setting teh parameter `AzureFunctionsJobHost__customHandler__description__defaultExecutablePath` to the value `server` (Remark: This can be done with any parameter in `host.json` using the parameter path separated by `__`).

### Routing (aka _catch all route_)

__YouTube Reference__: [Link](https://youtu.be/RPCEH247twU?t=1068)

__Branch__: routing

__Hints__:

* Adopt the go file and do not forget to register a new generic route.
* Do not forget to build the go file.
* Rename the function to `all`.
* Add the route property to the `function.json` file for the all function and set it to a wildcard.

### Debug

_YouTube Reference_: [Link](https://youtu.be/RPCEH247twU?t=1246)

_Branch_: debug

_Hints_:

* Add the property `logLevel` to `host.json` and set it to `{ "default": "Trace" }` in order to get an extensive log output in your console.
