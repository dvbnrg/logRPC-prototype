# This service captures log events over a log stream or transactions and stores the events in a database of choice

    - The services are defined over gRPC / Protobufs
    - The logging endpoint then inserts the log entry to the data store of choice
