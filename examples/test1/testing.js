import { Client } from 'felloe/k8s';
import { DeploymentFactory } from 'felloe/k8s/deployment'

export default function() {


    var podList = Client.listPods("");
    for(var i = 0; i < podList.length; i++) {
        console.log(podList[i].name);
    }

    console.log("----");
    podList = Client.listPods("kube-system");
    for(var i = 0; i < podList.length; i++) {
        console.log(podList[i].name);
    }

    var deployment = DeploymentFactory.name("deployment")
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

    return [deployment];
}