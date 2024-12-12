## Simple gRPC server on golang

```console
grpcurl -plaintext -d '{"text": "some note"}' localhost:8080 NoteService/AddNote
grpcurl -plaintext -d '{"id": 1}' localhost:8080 NoteService/GetNote
```
