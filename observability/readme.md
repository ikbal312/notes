Observability

    Definition:
    Observability is the ability to measure a system's current state based on the data it generates, recorded as logs, metrics, and traces. 

    Tracing:
    Tracing is a method to monitor and track requests as they traverse through various services and systems. 






    In fact, logs, metrics, and traces are known as the “three pillars of observability.”

    There are two major steps involved in setting up observability for your application:

    1. Collecting relevant data that indicates the application health
    2. Storing, managing, and visualizing the collected data to take quick actions
Tools:

Opentelemetry

    OpenTelemetry use to solve the first problem of Observability

    1. OpenTelemetry provides an instrumentation layer for your application code, while Jaeger is a backend analysis tool used for storage and visualization of trace data.
    2. Using OpenTelemetry libraries, you can generate logs, metrics, and traces. Jaeger does not support logs and metrics.
    3. OpenTelemetry can only be used to generate and collect data. It does not provide a storage layer. 
    4. OpenTelemetry does not provide any web UI components. Jaeger comes with a web UI component that is used for visualizing trace data.

Jaeger: 

    Jaeger use to solve the second problem of Observability

    Jaeger is a popular open-source distributed tracing tool that was originally built by teams at Uber and then open-sourced. It is used to monitor and troubleshoot applications based on microservices architecture.

    A distributed tracing tool tracks user requests across services and gives a central overview of how different components of a microservices architecture interact to process user requests. Jaeger is used to store, analyze and visualize tracing data.

    Jaeger does not support logs and metrics.

    It provides instrumentation libraries that were built on OpenTracing standards. The libraries cover popular programming languages like Go, Java, Node, Python, C++, and C#. For storing trace data, it supports two storage backends:

    1. Cassandra
    2. Elasticsearch
    \
    Jaeger provides a minimal UI to analyze the trace data captured.


    1. Jaeger provides Cassandra and Elasticsearch as two options for storing data.
    2. Jaeger comes with a web UI component that is used for visualizing trace data.
   
    