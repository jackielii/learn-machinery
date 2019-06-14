### example output:

```
~/gomod/learn-machinery   master ●  go run main.go
INFO: 2019/06/14 16:16:54 worker.go:46 Launching a worker with the following settings:
INFO: 2019/06/14 16:16:54 worker.go:47 - Broker: redis://localhost:6379
INFO: 2019/06/14 16:16:54 worker.go:49 - DefaultQueue:
INFO: 2019/06/14 16:16:54 worker.go:53 - ResultBackend: redis://localhost:6379
INFO: 2019/06/14 16:16:54 redis.go:86 [*] Waiting for messages. To exit press CTRL+C
DEBUG: 2019/06/14 16:16:54 redis.go:284 Received new message: {"UUID":"task_bceaf623-fbdf-4a66-95d9-08704308cfbc","Name":"task","RoutingKey":"","ETA":null,"GroupUUID":"","GroupTaskCount":0,"Args":[{"Name":"","Type":"string","Value":"hello please"}],"Headers":{"foo":"bar","map":{"cool":"nested header"},"struct":{"Foo":"foo struct"}},"Immutable":false,"RetryCount":0,"RetryTimeout":0,"OnSuccess":null,"OnError":null,"ChordCallback":null,"BrokerMessageGroupId":""}
===================================================
(*context.valueCtx)(0xc00038c2d0)(context.Background.WithValue(tasks.signatureCtxType{}, &tasks.Signature{UUID:"task_bceaf623-fbdf-4a66-95d9-08704308cfbc", Name:"task", RoutingKey:"", ETA:(*time.Time)(nil), GroupUUID:"", GroupTaskCount:0, Args:[]tasks.Arg{tasks.Arg{Name:"", Type:"string", Value:"hello please"}}, Headers:tasks.Headers{"foo":"bar", "map":map[string]interface {}{"cool":"nested header"}, "struct":map[string]interface {}{"Foo":"foo struct"}}, Immutable:false, RetryCount:0, RetryTimeout:0, OnSuccess:[]*tasks.Signature(nil), OnError:[]*tasks.Signature(nil), ChordCallback:(*tasks.Signature)(nil), BrokerMessageGroupId:""}).WithValue(opentracing.contextKey{}, opentracing.noopSpan{}))
received  hello please
===================================================
DEBUG: 2019/06/14 16:16:54 worker.go:248 Processed task task_bceaf623-fbdf-4a66-95d9-08704308cfbc. Results = returning from task
result 0: returning from taskresults:  returning from task
WARNING: 2019/06/14 16:16:54 broker.go:100 Stop channel
```
