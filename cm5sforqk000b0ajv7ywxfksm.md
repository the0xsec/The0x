---
title: "Building a Enterprise Security Gateway: A Technical Overview"
seoTitle: "Enterprise Security Gateway: Technical Guide"
seoDescription: "A technical overview of building an enterprise security gateway using microservices, focusing on security, consistency, and observability strategies"
datePublished: Sat Jan 11 2025 17:02:04 GMT+0000 (Coordinated Universal Time)
cuid: cm5sforqk000b0ajv7ywxfksm
slug: building-a-enterprise-security-gateway-a-technical-overview
tags: distributed-system, side-project, homelab

---

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1736611484942/9dc6d78f-1f4f-4599-82e0-0787c689b22e.png align="center")

Let me break down the thinking behind this distributed system design and the technical choices I've made. After working with microservices architectures for a few years not, but I am definetely not a expert at the software development side. I have found that the complexities isn’t in the individual services, but how data and requests flow between the services.

The request flow starts at our edge layer through Cloudflare, which ahdnles initial DDoS protection and TLS termination. I chose to implement progressive request enhancement, which where each service adds context and validation metadata to the request at each stage of the request lifecycle. This approach *should* prove more maintainable than having centralized validation.

Now, using service communication introduces some initial complexities, but I think it will significantly reduce longer-term technical debt. the services communicate through a pub/sub system which lets us add new functionality within modification of exiting services.

In the implementation of health checking, I will be taking significant step beyond the conventional basic HTTP health checks. Each service will need to provide comprehensive and detailed health metrics that offer valuable insights into the overall operational status of the system. These metrics include crucial information such as the status of dependencies, which indicates whether external services that our system relies on are functioning properly. Additionally, I will have to incorporate response times, which measure how quickly each service responds to requests, and error rates, which track the frequency of errors encountered during operations.

Now for data consistency… eh. I think I am going to try to go with the “Saga” pattern, but time will tell. The Saga pattern is a design pattern used to manage data consistency across distributed systems, particularly in microservices architectures. It addresses the challenges of maintaining consistency without using traditional distributed transactions, which can be complex and difficult to scale. With that, I think for this project having a eventual consistency pattern is a reasonable justification.

Security will be implemented as any good security program is, with multiple layers. The authentication service handles JWT generation and validation, while OAtuh2 managed external service integrations. Rate limiting will run at both the edge and on individual service layers. Each service validates it’s inputs and will have to run with minimal required permissions. (Whatever that means in this case…)

I am considering creating a observability stack using openTelemetry and prometheus for metrics and structured logging. This combination is a tried and true pattern for debugging production issues and traces will have to be correlated through unique request IDs.

Now, I chose Pulumi for the Infrastructure as Code because I want to keep the context between application and infrastructure the same, and I find Pulumi to be a great tool for doing just that. I am going to have to figure out how Pulumi and work with ngrok, if at all. That might be it’s own side project later in the year if needed.

Overall, I am quite excited to start working, and I am sure the people who have decided to follow along will learn as I learn. I don’t have a dedicated timeline to complete the project, and I can easily see this project growing and I expand my understanding, but that is kinda of the goal when it comes to side-projects/homelabs and the like. Continuous Learning.