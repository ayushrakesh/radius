---
type: docs
title: "Application networking"
linkTitle: "Networking"
description: "Learn how to add networking to your Radius application"
weight: 200
categories: "Concept"
tags: ["routes","gateways"]
---

Radius networking resources allow you to model:

- Communication between a user and a service
- Communication between services

## HTTP Routes

An `HttpRoute` resources defines HTTP communication between two [services]({{< ref container >}}). They can be used to define both one-way communication, as well as cycles of communication between services.

<img src="networking-cycles.png" style="width:400px" alt="Diagram of Radius service-to-service networking with cycles" /><br />

Refer to the [HTTP Route schema]({{< ref httproute >}}) for more information on how to model HTTP routes.

A gateway can optionally be added for external users to access the Route.

## Gateways

`Gateway` defines how requests are routed to different resources, and also provides the ability to expose traffic to the internet. Conceptually, gateways allow you to have a single point of entry for  traffic in your application, whether it be internal or external traffic.

`Gateway` in Radius are split into two main pieces; the `Gateway` resource itself, which defines which port and protocol to listen on, and Route(s) which define the rules for routing traffic to different resources.

<img src="networking-gateways.png" style="width:400px" alt="Diagram of Radius gateways" /><br />

Refer to the [Gateway schema]({{< ref gateway >}}) for more information on how to model gateways.

### SSL Passthrough

A gateway can be configured to passthrough encrypted SSL traffic to an HTTP route and container. This is useful for applications that already have SSL termination configured, and do not want to terminate SSL at the gateway.

To set up SSL passthrough, set `tls.sslPassthrough` to `true` on the gateway, and set a single route with no `path` defined (just `destination`).

## Example

### Path-based HTTP routing

{{< tabs Bicep >}}

{{< codetab >}}
{{< rad file="snippets/networking.bicep" embed=true >}}
{{< /codetab >}}

{{< /tabs >}}

### SSL Passthrough

{{< tabs Bicep >}}

{{< codetab >}}
{{< rad file="snippets/networking-sslpassthrough.bicep" embed=true marker="//GATEWAY" >}}
{{< /codetab >}}

{{< /tabs >}}