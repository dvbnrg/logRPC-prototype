# logRPC

    - This service captures log events over a log stream or transactions
    - The services are defined over gRPC / Protobufs
    - The logging endpoint will then inserts the log entry to the data store of choice (TBD)
        - For now it just saves the output to a json file
        - TODO: add in a log rolling mechanism

<a href="https://imgflip.com/i/48xkh5"><img src="https://i.imgflip.com/48xkh5.jpg" title="made at imgflip.com"/></a><div><a href="https://imgflip.com/memegenerator">from Imgflip Meme Generator</a></div>
