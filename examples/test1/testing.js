import Calc from './module';
import { Client, Namespace } from 'felloe/k8s';
import { DeploymentFactory } from 'felloe/k8s/deployment'

export default function() {

    var a = new Calc();
    console.error(a.sum(1,1));

    let podList = Client.listPods("");
    for(let i = 0; i < podList.length; i++) {
        console.log(podList[i].name);
    }

    console.log("----");
    podList = Client.listPods("kube-system");
    for(let i = 0; i < podList.length; i++) {
        console.log(podList[i].name);
    }

    let deployment = DeploymentFactory.name("deployment")
        .namespace("spinnaker")
        .annotations({
            "environment": "production"
        })
        .addAnnotation("another", "annotation")
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
                env: [{
                    name: "POD_NAME",
                    valueFrom: {
                        fieldRef: {
                            apiVersion: "v1",
                            fieldPath: "metadata.name"
                        }
                    }
                }]
            },
            {
                name: "app2",
                image: "nginx:2"
            },
        ])
        .build();

    let namespace = new Namespace("spinnaker");

    return [deployment, namespace];
}