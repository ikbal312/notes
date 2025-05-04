Open Telemetry(OTel)

Terms:

Aggregation - সমষ্টি

    Is A process
    Combines multiple measurement[that took place during interval of time,during program execution] into exact or estimated
    The measurement . Used by Data Source,Metric

API / Application Programming Interface

    used to define how telemetry data is generated per Data Source in OTel Project

Application 

    One or more services designed for end users or other applications.

APM / Application Performance Monitoring

    APM is about monitoring software application 
    App performance(speed,reliability,availability, and so on) 
    detect issues, alerting and tooling for finding the root cause.

Attribute / OTel Terms for Metadata

    Adds key-value information to the entity producing telemetry.
    Used across Signals and Resources

Automatic Instrumentation

    Refers to telemetry collection methods, that do not required end-user to modify application source code.

Baggage / Metadata Propagation Mechanism,
    
    helps establish a casual relationship between events and services

Client-Side App

    example: browser-app, mobile-app, apps running on IoT devices


Collector / OTel Collector

    A single Binary that can be displayed as an agent or gateway.
    Receive->Process->Export OTel Data.

Contrib

    several instrumentation libraries and Collector
    non-core capabilities including vendor Exporter


Context Propagation

    Allow all data source to share an underlying context mechanism for storing state and accessing data across the lifespan of a Transaction.

DAG

    Directed Acyclic Graph

Data Source

Signals

    Traces - 
    Metrics
    Logs
    Baggage
    <!-- proposal -->
    Events a specific type of log
    Profiles

Span

    Represent a single operation within a Trace


Span-Link

    A span link is a link between casually-related span

Specification

    describes the cross-language requirements and expectations for all implementations


Status

    The result of the operation. Typically used to indicate whether an error occurred

Trace

    Responsible For creating spans

Distributed tracing

    Tracks the progression of a single Request, called a Trace, as it is handled by Services that make up an Application. A Distributed trace transverses process, network and security boundaries.

Semantic Convention

    Defines standard names and values

SDK

    Software Development Kits.
    Refers to Otel SDK that denotes a Library that implements the Otel API

Sampling 

    A mechanism to control the amount of data exported.
    Mostly used with the Tracing Data Source.

Resource

    Captures information about the entity producing telemetry as Attributes.
    example, a process producing telemetry that is running in a container on Kubernetes has a process name, a pod name, a namespace, and possibly a deployment name. 
    All these attributes can be included in the Resource



Instrumented library

    Denotes the Library for which the telemetry signals (Traces, Metrics, Logs) are gathered. See more

Instrumentation library

    Denotes the Library that provides the instrumentation for a given Instrumented library. 
    Instrumented library and Instrumentation library can be the same Library if it has built-in OpenTelemetry instrumentation


HTTP

    Short for Hypertext Transfer Protocol.

Rest

    Representation State Transfer

RPC

    Remote Procedure Call
gRPC

    A high-performance, open source universal RPC framework. More on gRPC here

Exporter

    Provides functionality to emit telemetry to consumers. Exporters can be push- or pull-based

Event

    An Event is a Log Record with an event name and a well-known structure. 
    For example, browser events in OpenTelemetry follow a particular naming convention and carry particular data in a common structure.


Dimension

    A term used specifically by Metrics. See Attribute.

Log

    Sometimes used to refer to a collection of Log records. Can be ambiguous since people also sometimes use Log to refer to a single Log record. Where ambiguity is possible, use additional qualifiers, for example, Log record. See more
    Log record

    A recording of data with a timestamp and a severity. May also have a Trace ID and Span ID when correlated with a trace. See more.


Instrumentation

     Instrumentation is the process of enabling your application code to generate telemetry data.




