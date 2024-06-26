
## duckspark.io - go bigger ![](https://github.githubassets.com/images/icons/emoji/unicode/1f680.png?v2)

<https://duckspark.io/>  

duckspark is library for enabling distributed processing of large data sets. It allows you to efficiently process data in parallel across multiple machines or nodes, enabling faster and scalable data processing.

[![](https://skillicons.dev/icons?i=go)](https://skillicons.dev) 


##  Features   

- Distributed processing: The library allows you to distribute the processing of large data sets across multiple nodes, leveraging the power of parallel processing.  

- Scalability: It supports scaling your data processing by adding more machines or nodes to the cluster, allowing you to handle larger datasets or increasing the processing speed.
Fault tolerance: The library incorporates fault tolerance mechanisms to handle failures or crashes in the cluster. It provides automatic recovery and resilience to ensure uninterrupted processing.


- Load balancing: It implements intelligent load balancing algorithms to distribute the workload evenly across nodes, optimizing resource utilization.


- Data partitioning: The library offers efficient data partitioning techniques, enabling parallel processing on smaller subsets of the data across different nodes.
Simplified API: It provides a simple and intuitive API to facilitate the development of distributed data processing applications.

----

##  Building duckspark 

``` apt install -y protobuf-compiler ```

``` go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 ```

```protoc --go_out=. protos/*.proto```


## Contributing

for information on how to get started contributing to the project.