# Felloe - k8s as code

Felloe is an alternative tool to Helm.

Instead of rendering manifests using a template engine, Felloe uses a JavaScript VM.

It exposes some native modules to help you interact with your k8s cluster and easily create resources.

## Anatomy of a script

Every script needs to export a default function and return a list of k8s resources to deploy.

```js
export default function() {
    return [];
}
```

You can import external or natives modules using an `import` statement. All native modules are prefixed with `felloe/`.

```js
// calc.js
export default class Calc {
    constructor() {}

    sum(a, b) {
        return a + b;
    }
}
```

```js
import { Client, Namespace } from 'felloe/k8s'; // Native module
import Calc from './calc'; // Your module
export default function() {
    let calc = new Calc();
    console.info(calc.sum(1,1));
    
    let resourcesToDeploy = [];    


    let namespaces = Client.listNamespaces();
    if(namespaces.length == 2) {
        let n = new Namespace("my-new-namespace");
        resourcesToDeploy.push(n);
    }
    
    return resourcesToDeploy;
}
``` 

## Example

```js
import { Client, Namespace } from 'felloe/k8s';
import { DeploymentFactory } from 'felloe/k8s/deployment'

export default function() {

    let resourcesToDeploy = [];

    let podList = Client.listPods("kube-system");
    for(let i = 0; i < podList.length; i++) {
        console.log(podList[i].name);
    }

    let deployment = DeploymentFactory.name("deployment")
        .namespace("my-namespace")
        .annotations({
            "environment": "production"
        })
        .selector({
            matchLabels: {
                "app.kubernetes.io/instance": "my-app"
            }
        })
        .podLabels({
            "app.kubernetes.io/instance": "my-app"
        })
        .containers([
            {
                name: "app",
                image: "nginx",
            }
        ]).build();

    resourcesToDeploy.push(deployment);

    let namespace = new Namespace("my-namespace");
    resourcesToDeploy.push(namespace);
   
    return resourcesToDeploy;
}
```

## Usage

The only available command at the moment is `template`. It will run a script and return the generated manifests.

```sh
$ felloe template ./myscript.js 
```

## Project Status

This is currently a POC and a project I've started to expand my knowledge in golang. **Don't use for production workloads**

## Contributing

If you find the concept interesting feel free to submit a PR or message me if you would like to be an active contributor. 
